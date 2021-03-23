package db

import (
	"context"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/* ReadTweetsFollowers */
func ReadTweetsFollowers(ID string, page int) ([]models.ReturnTweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("app-like-twitter")
	col := db.Collection("relation")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userId": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from": "tweet",
			"localField": "userRelationShipId",
			"foreignField": "userId",
			"as": "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind":"$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.dateTweet": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)
	var result []models.ReturnTweetsFollowers

	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}