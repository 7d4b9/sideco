package scoring

//go:generate protoc api.proto --go_out=plugins=grpc:.

import (
	context "context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	*mongo.Database
}

type User struct {
	SiderID string
	Tags    []string
}

type ScoredUser struct {
	*User
	Score        int32
	MatchingTags []string
}

func (s *Handler) GetScoreTags(ctx context.Context, taskID string, users []User) (scoredUsers []ScoredUser, tags []string, err error) {
	tagsColl := s.Collection("tasksTags")
	res := tagsColl.FindOne(ctx, bson.D{{"_id", taskID}})
	if err = res.Err(); err != nil {
		log.WithError(err).WithField("task_id", taskID).Error("db error")
		return
	}
	var tasks struct {
		Tags []string `bson:"tags"`
	}
	if err = res.Decode(&tasks); err != nil {
		log.WithError(err).Error("decode user")
		return
	}
	taskTags := tasks.Tags
	for _, tag := range taskTags {
		for i := range users {
			user := &users[i]
			var matchingTags []string
			for _, userTag := range user.Tags {
				if tag == userTag {
					matchingTags = append(matchingTags, tag)
				}
			}
			matching, total := len(matchingTags), len(taskTags)
			if matching != 0 && total != 0 {
				scoredUsers = append(scoredUsers, ScoredUser{
					User:         user,
					MatchingTags: matchingTags,
					Score:        int32((100 * matching) / total),
				})
			}
		}
	}
	tags = taskTags
	return
}
