package grpc

import (
	context "context"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/bbrod/sideco/scoring"
	grpc "google.golang.org/grpc"
)

type setting string

const (
	port setting = "port"
	host setting = "host"
)

var v = viper.New()

func init() {
	v.AutomaticEnv()
	v.SetEnvPrefix("scoring")
	v.SetDefault(string(port), "8080")
	v.SetDefault(string(host), "scoring")
}

type Server struct {
	*scoring.Handler
}

func (s *Server) GetScore(ctx context.Context, req *scoring.GetScoreRequest) (*scoring.GetScoreResponse, error) {
	var users []scoring.User
	for i := range req.Applicants {
		user := req.Applicants[i]
		users = append(users, scoring.User{SiderID: user.SiderID, Tags: user.Tags})
	}
	scoreds, tags, err := s.GetScoreTags(ctx, req.TaskId, users)
	if err != nil {
		return nil, err
	}
	var out scoring.GetScoreResponse
	for i := range scoreds {
		scored := &scoreds[i]
		out.Scores = append(out.Scores, &scoring.GetScoreResponse_Applicant{Score: scored.Score, SiderID: scored.SiderID})
	}
	out.Tags = tags
	return &out, nil
}

func (s *Server) Run(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		lis, err := net.Listen("tcp", v.GetString(string(host))+":"+v.GetString(string(port)))
		if err != nil {
			log.WithError(err).Error("listen failed")
			return
		}
		grpcServer := grpc.NewServer()
		scoring.RegisterScoringServer(grpcServer, s)
		if err := grpcServer.Serve(lis); err != nil {
			log.WithError(err).Error("grpc server down")
		}
	}()
	return done
}
