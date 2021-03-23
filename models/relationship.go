package models

/* RelationShip is the model to save the relation of the user with others */
type RelationShip struct {
	UserID				string		`bson:"userId" json:"userId"`
	UserRelationShipID	string		`bson:"userRelationShipId" json:"userRelationShipId"`
}
