package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/**/
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("app-like-twitter")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id": objID,
		"userId": UserID,
	}
	_, err := col.DeleteOne(ctx, condition)
	return err
}