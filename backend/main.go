package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/bbrod/sideco/mongo"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.NewClient(ctx)
	if err != nil {
		log.WithError(err).Fatal("no database")
	}
	// tagDB := client.Database("main").
	client.Database("main")
	<-make(chan struct{})
}
