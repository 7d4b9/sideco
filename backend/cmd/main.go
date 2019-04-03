package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/bbrod/sideco/backend"
	"gitlab.com/bbrod/sideco/backend/http"
	"gitlab.com/bbrod/sideco/mongo"
	"gitlab.com/bbrod/sideco/scoring"
	"gitlab.com/bbrod/sideco/scoring/grpc"
)

func main() {
	ctx := context.TODO() // catch SIGINT
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	client, err := mongo.NewClient(ctx)
	if err != nil {
		log.WithError(err).Fatal("no database")
	}
	mainDB := client.Database("main")
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	scoringConn, err := grpc.Connect(ctx)
	if err != nil {
		log.WithError(err).Fatal("no scoring connection")
	}
	defer scoringConn.Close()
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	scoringClient := scoring.NewScoringClient(scoringConn)
	handler := &backend.Handler{Database: mainDB, ScoringClient: scoringClient}
	server := http.NewServer(handler)
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	<-server.Run(ctx)
}
