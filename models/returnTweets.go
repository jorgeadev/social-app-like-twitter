package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/* ReturnTweets is the struct with we return los tweets */
type ReturnTweets struct {
	ID			primitive.ObjectID		`bson:"_id" json:"_id,omitempty"`
	UserID		string					`bson:"userID" json:"userID,omitempty"`
	Message		string					`bson:"message" json:"message,omitempty"`
	DateTweet	time.Time				`bson:"dateTweet" json:"dateTweet,omitempty"`
}