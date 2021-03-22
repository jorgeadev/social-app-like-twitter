package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Claim is the structure used by process the JWT  */
type Claim struct {
	Email	string					`json:"email"`
	ID		primitive.ObjectID		`bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
