package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/* ReturnTweetsFollowers */
type ReturnTweetsFollowers struct {
	ID					primitive.ObjectID		`bson:"_id" json:"_id,omitempty"`
	UserID				string					`bson:"userId" json:"userId,omitempty"`
	UserRelationShipID	string					`bson:"userRelationShipId" json:"userRelationShipId,omitempty"`
	Tweet struct {
		Message		string		`bson:"message" json:"message,omitempty"`
		DateTweet	time.Time	`bson:"dateTweet" json:"dateTweet,omitempty"`
		ID			string		`bson:"_id" json:"_id,omitempty"`
	}
}