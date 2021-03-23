package db

import (
	"context"
	"fmt"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/* ReadAllUsers Read the users registered in the system, if received "R" get only that are related with me */
func ReadAllUsers(ID string, page int64, search string, _type string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("app-like-twitter")
	col := db.Collection("users")

	var result []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	var founded, include bool
	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return result, false
		}
		var r models.RelationShip
		r.UserID = ID
		r.UserRelationShipID = s.ID.Hex()

		include = false
		founded, err = ConsultRelations(r)
		if _type == "new" && founded == false {
			include = true
		}
		if _type == "follow" && founded == true {
			include = true
		}
		if r.UserRelationShipID == ID {
			include = false
		}
		if include == true {
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			result = append(result, &s)
		}
	}
	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}
	cursor.Close(ctx)
	return result, true
}
