package db

import (
	"context"
	"fmt"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/* ConsultRelations consult the relation between two users */
func ConsultRelations(t models.RelationShip) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("app-like-twitter")
	col := db.Collection("relation")

	condition := bson.M{
		"userId": t.UserID,
		"userRelationShipId": t.UserRelationShipID,
	}

	var result models.RelationShip
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
