package models

import "time"

/* RecordTweet is the format or struct for own tweet in database */
type RecordTweet struct {
	UserID		string		`bson:"userId" json:"userId,omitempty"`
	Message		string		`bson:"message" json:"message,omitempty"`
	DateTweet	time.Time	`bson:"dateTweet" json:"dateTweet,omitempty"`
}
