package scoring

//go:generate protoc api.proto --go_out=plugins=grpc:.

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	*mongo.Database
}

func (s *Handler) GetScoreTags(ctx context.Context, taskID string, userTags []string) (score int32, tags []string, err error) {
	tagsColl := s.Collection("tasksTags")
	cur, err := tagsColl.Find(ctx, bson.D{{"_id", taskID}})
	if err != nil {
		log.WithError(err).WithField("task_id", taskID).Error("db error")
		return
	}
	defer cur.Close(ctx)
	var (
		matchCount int32
		tagsCount  = int32(len(userTags))
	)
	var matchingTags []string
	for cur.Next(ctx) {
		var tasks struct {
			ID   string   `bson:"_id"`
			Tags []string `bson:"tags"`
		}
		if err = cur.Decode(&tasks); err != nil {
			log.WithError(err).Error("decode user")
			return
		}
		if count := int32(len(tasks.Tags)); tagsCount < count {
			tagsCount = count
		}
		for _, tag := range tasks.Tags {
			for _, userTag := range userTags {
				if tag == userTag {
					matchingTags = append(matchingTags, tag)
					matchCount++
				}
			}
		}
		continue //process only the first matching
	}
	if tagsCount != 0 {
		score = ((100 * matchCount) / tagsCount)
		tags = matchingTags
	}
	return
}
