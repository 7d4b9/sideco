package backend

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/bbrod/sideco/scoring"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type setting string

const (
	timeout setting = "timeout"
)

var v = viper.New()

func init() {
	v.AutomaticEnv()
	v.SetEnvPrefix("backend")
	v.SetDefault(string(timeout), "30s")
}

type EvalRequest struct {
	TaskID string `json:"taskId"`
}

type Applicant struct {
	SiderID   string `json:"siderId" bson:"_id"`
	Score     int32  `json:"score"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}

type EvalResponse struct {
	TaskID      string      `json:"taskId" bson:"_id,omitempty"`
	Applicants  []Applicant `json:"applicants" bson:"_"`
	Description string      `json:"description" bson:"description,omitempty"`
	Country     string      `json:"country" bson:"country,omitempty"`
	Tags        []string    `json:"tags" bson:"tags,omitempty"`
}

type Handler struct {
	*mongo.Database
	scoring.ScoringClient
}

func (s *Handler) Eval(req *EvalRequest) (resp *EvalResponse, err error) {
	taskID := req.TaskID
	resp = &EvalResponse{TaskID: taskID}
	ctx, cancel := context.WithTimeout(context.Background(), v.GetDuration(string(timeout)))
	defer cancel()
	usersColl := s.Collection("users")
	cur, err := usersColl.Find(ctx, bson.D{{"taskApplications", taskID}})
	if err != nil {
		log.WithField("task_id", taskID).Error("db error")
		return nil, err
	}
	defer cur.Close(ctx)
	scoringRequest := &scoring.GetScoreRequest{TaskId: taskID}
	for i := 0; cur.Next(ctx); i++ {
		var user struct {
			SiderID string   `bson:"_id"`
			Tags    []string `bson:"tags"`
		}
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		scoringRequest.Applicants = append(scoringRequest.Applicants, &scoring.GetScoreRequest_Applicant{SiderID: user.SiderID, Tags: user.Tags})
	}
	scoringResponse, err := s.GetScore(ctx, scoringRequest)
	if err != nil {
		return nil, err
	}
	if len(scoringResponse.Scores) == 0 {
		return nil, err
	}
	var applicants []Applicant
	for i := range scoringResponse.Scores {
		resp := scoringResponse.Scores[i]
		res := usersColl.FindOne(ctx, bson.D{{"_id", resp.SiderID}})
		if err := res.Err(); err != nil {
			log.Error("query user details")
			return nil, err
		}
		applicant := Applicant{SiderID: resp.SiderID, Score: resp.Score}
		if err := res.Decode(&applicant); err != nil {
			log.Error("cannot decode applicant")
			return nil, err
		}
		applicants = append(applicants, applicant)
	}
	taskColl := s.Collection("tasks")
	task := taskColl.FindOne(ctx, bson.D{{"_id", taskID}})
	if err != nil {
		log.Error("cannot query task collection")
		return nil, err
	}
	out := EvalResponse{TaskID: taskID, Tags: scoringResponse.Tags, Applicants: applicants}
	if err := task.Decode(&out); err != nil {
		log.Error("cannot add task information")
		return nil, err
	}
	return &out, nil
}
