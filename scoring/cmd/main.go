package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/bbrod/sideco/mongo"
	"gitlab.com/bbrod/sideco/scoring"
	grpc "gitlab.com/bbrod/sideco/scoring/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.NewClient(ctx)
	if err != nil {
		log.WithError(err).Fatal("no database")
	}
	tagDB := client.Database("tag")
	handler := scoring.Handler{Database: tagDB}
	server := grpc.Server{Handler: &handler}
	<-server.Run(ctx)
}
