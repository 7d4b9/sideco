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
	TaskID      string    `json:"taskId" bson:"_id,omitempty"`
	Applicants  Applicant `json:"applicants" bson:"_"`
	Description string    `json:"description" bson:"description,omitempty"`
	Country     string    `json:"country" bson:"country,omitempty"`
	Tags        []string  `json:"tags" bson:"tags,omitempty"`
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
		log.WithError(err).WithField("task_id", taskID).Error("db error")
		return nil, err
	}
	defer cur.Close(ctx)
	best := &EvalResponse{TaskID: taskID}
	for cur.Next(ctx) {
		var user struct {
			Tags []string `bson:"tags"`
		}
		if err := cur.Decode(&user); err != nil {
			log.WithError(err).Error()
			return nil, err
		}
		scoringRequest := &scoring.GetScoreRequest{TaskId: taskID, Tag: user.Tags}
		scoringResponse, err := s.GetScore(ctx, scoringRequest)
		if err != nil {
			log.WithError(err).Error()
			return nil, err
		}
		replaceBest := func() error {
			if err := cur.Decode(&best.Applicants); err != nil {
				log.Error("cannot replace best response")
				return err
			}
			best.Applicants.Score = scoringResponse.Score
			best.Tags = scoringResponse.MatchingTag
			return nil
		}
		if scoringResponse.Score > best.Applicants.Score {
			if err := replaceBest(); err != nil {
				return nil, err
			}
			continue
		}
		if scoringResponse.Score == best.Applicants.Score && len(scoringResponse.MatchingTag) > len(best.Tags) {
			if err := replaceBest(); err != nil {
				return nil, err
			}
			continue
		}
	}
	taskColl := s.Collection("tasks")
	task := taskColl.FindOne(ctx, bson.D{{"_id", taskID}})
	if err != nil {
		log.Error("cannot query task collection")
		return nil, err
	}
	if err := task.Decode(&best); err != nil {
		log.WithError(err).Error("cannot add task information")
		return nil, err
	}
	return best, nil
}
