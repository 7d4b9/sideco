package scoring

//go:generate protoc api.proto --go_out=plugins=grpc:.

import (
	context "context"

	"gitlab.com/bbrod/sideco/scoring"
)

type Server struct {
	*scoring.Handler
}

func (s *Server) ListTaskTags(ctx context.Context, req *ListTaskTagsRequest) (*ListTaskTagsResponse, error) {
	tags, err := s.Handler.ListTaskTags(ctx, req.TaskId)
	if err != nil {
		return nil, err
	}
	return &ListTaskTagsResponse{Tag: tags}, nil
}

func (s *Server) Run(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)

	}()
	return done
}
