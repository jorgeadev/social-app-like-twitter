package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"time"
)

/* GenerateJWT generate the encrypted con JWT */
func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("CourseToDevelopAndAppLikeTwitterUdemy")

	payload := jwt.MapClaims{
		"email": t.Email,
		"name": t.Name,
		"lastName": t.LastName,
		"birthDate": t.BirthDate,
		"biography": t.Biography,
		"location": t.Location,
		"website": t.WebSite,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}