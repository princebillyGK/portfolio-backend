package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Portfolio struct {
	collection *mongo.Collection
}

func NewPortfolio(c *mongo.Client) *Portfolio {
	return &Portfolio{collection: c.Database("portfolio").Collection("resume")}
}

func (r Portfolio) GetById(ctx context.Context, id string) error {
	var result bson.M

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = r.collection.FindOne(ctx, bson.D{{"_id", oid}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return err
	}
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", jsonData)
	return nil
}
