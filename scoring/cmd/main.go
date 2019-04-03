package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/bbrod/sideco/mongo"
	"gitlab.com/bbrod/sideco/scoring"
	grpc "gitlab.com/bbrod/sideco/scoring/grpc"
)

func main() {
	ctx := context.TODO() // catch SIGINT
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	client, err := mongo.NewClient(ctx)
	if err != nil {
		log.WithError(err).Fatal("no database")
	}
	tagDB := client.Database("tags")
	handler := scoring.Handler{Database: tagDB}
	server := grpc.Server{Handler: &handler}
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	<-server.Run(ctx)
}
