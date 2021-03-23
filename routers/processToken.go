package routers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"strings"
)

/* Email is a value used in all EndPoints */
var Email string
/* IDUser is the ID returned of the model, that will be used in all EndPoints */
var IDUser string

/* ProcessToken process the token to get values */
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("CourseToDevelopAndAppLikeTwitterUdemy")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, founded, _ := db.CheckUserExist(claims.Email)
		if founded == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, founded, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}


/*func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("CourseToDevelopAndAppLikeTwitterUdemy")
	claims := &models.Claim{}
	
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, founded, _ := db.CheckUserExist(claims.Email)
		if founded == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, founded, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}*/