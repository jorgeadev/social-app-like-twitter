package models

/* Tweet get from Body the message */
type Tweet struct {
	Message		string		`bson:"message" json:"message"`
}
