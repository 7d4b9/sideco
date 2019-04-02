package mongo

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type setting string

const (
	host     setting = "host"
	port     setting = "port"
	username setting = "username"
	password setting = "password"
	database setting = "database"
)

var v = viper.New()

func init() {
	v.AutomaticEnv()
	v.SetEnvPrefix("mongo")
	v.SetDefault(string(username), "side")
	v.SetDefault(string(password), "test")
	v.SetDefault(string(host), "mongo")
	v.SetDefault(string(port), "27017")
	v.SetDefault(string(database), "test")
}

// Client abstract an underlying client layer.
type Client struct {
	*mongo.Client
}

// NewClient creates a new client based upon an underlying layer and check additional startup sanity conditions.
func NewClient(ctx context.Context) (*Client, error) {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s:%s/%s`,
		v.GetString(string(username)),
		v.GetString(string(password)),
		v.GetString(string(host)),
		v.GetString(string(port)),
		v.GetString(string(database)),
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}
