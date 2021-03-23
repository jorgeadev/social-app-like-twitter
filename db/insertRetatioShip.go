package db

import (
	"context"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"time"
)

/* InsertRelationShip save the relation in database */
func InsertRelationShip(t models.RelationShip) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("app-like-twitter")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
