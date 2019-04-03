package scoring

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	*mongo.Database
}

func (s *Handler) ListTaskTags(ctx context.Context, taskID string) ([]string, error) {
	// use s.Database to retrieve the tags
	return []string{}, nil
}
