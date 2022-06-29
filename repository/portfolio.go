package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"princebillygk.portfolio.io/model"
)

type Portfolio struct {
	collection *mongo.Collection
}

func NewPortfolio(c *mongo.Client) *Portfolio {
	return &Portfolio{collection: c.Database("portfolio").Collection("resume")}
}

func (r Portfolio) GetById(ctx context.Context, id string) (*model.Portfolio, error) {
	var result model.Portfolio
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.collection.FindOne(ctx, bson.D{{"_id", oid}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &result, nil
}
