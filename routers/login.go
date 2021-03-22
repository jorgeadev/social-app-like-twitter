package routers

import (
	"encoding/json"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/jwt"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"log"
	"net/http"
	"time"
)

/* Login to login */
func Login(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(request.Body).Decode(&t)
	log.Println(json.NewDecoder(request.Body).Decode(&t))

	if err != nil {
		http.Error(writer, "User or password incorrect!!! "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(writer, "Email field is required!!!", 400)
		return
	}
	document, exist := db.IntentLogin(t.Email, t.Password)
	if exist == false {
		http.Error(writer, "User or password incorrect!!! ", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(writer, "An error was occurred at intent generate the Token !!! "+err.Error(), 400)
		return
	}
	resp := models.ResponseLogin {
		Token: jwtKey,
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(resp)


	/* Set cookie navigation */
	expitationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(writer, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expitationTime,
	})
}
