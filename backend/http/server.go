package http

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/bbrod/sideco/backend"
)

var v = viper.New()

type setting string

const (
	shutdownTimeout setting = "shutdown_timeout"
)

func init() {
	v.AutomaticEnv()
	v.SetEnvPrefix("backend_http")
	v.SetDefault(string(shutdownTimeout), "10s")
}

type Server struct {
	*http.Server
}

func NewServer(backendHandler *backend.Handler) *Server {
	return &Server{
		&http.Server{
			Addr:           ":8080",
			Handler:        NewHandler(backendHandler),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Run(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		if err := s.ListenAndServe(); err != nil {
			log.WithError(err).Error("on exit")
		}
	}()
	go func() {
		select {
		case <-ctx.Done():
			timeout := v.GetDuration(string(shutdownTimeout))
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			s.Shutdown(ctx)
		case <-done:
		}
	}()
	return done
}
