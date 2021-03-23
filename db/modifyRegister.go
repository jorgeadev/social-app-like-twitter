package db

import (
	"context"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/* ModifyRegister allows modify the profile of user */
func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("app-like-twitter")
	col := db.Collection("users")

	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		register["lastName"] = u.LastName
	}

	register["birthDate"] = u.BirthDate

	if len(u.Avatar) > 0 {
		register["banner"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}

	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}

	if len(u.Location) > 0 {
		register["location"] = u.Location
	}

	if len(u.WebSite) > 0 {
		register["webSite"] = u.WebSite
	}

	updateString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}