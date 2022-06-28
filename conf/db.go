package conf

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"princebillygk.portfolio.io/pkg/util"
)

func NewMongoClient(ctx context.Context) (client *mongo.Client, closeDB func(), err error) {
	uri := util.MustHaveEnv("MONGODB_URI")
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return
	}
	closeDB = func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
	return
}
