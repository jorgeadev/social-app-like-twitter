package db

import (
	"context"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"time"
)

/* DeleteRelationShip delete the relation in database */
func DeleteRelationShip(t models.RelationShip) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("app-like-twitter")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
