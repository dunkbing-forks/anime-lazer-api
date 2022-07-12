package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Paginate(page, limit int64, findOptions *options.FindOptions) (int64, int64) {
	if page == 1 {
		findOptions.SetSkip(0)
		findOptions.SetLimit(limit)
		return page, limit
	}
	findOptions.SetSkip((page - 1) * limit)
	findOptions.SetLimit(limit)
	return page, limit
}

func QuickCrud(collection *mongo.Collection, query interface{}, update interface{}, options *options.UpdateOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.UpdateOne(ctx, query, update, options)
	if err != nil {
		return err
	}
	return nil
}
