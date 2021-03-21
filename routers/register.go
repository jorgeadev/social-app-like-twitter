package routers

import (
	"encoding/json"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"net/http"
)

/* Register is the function to create the users register in database */
func Register(writer http.ResponseWriter, request *http.Request)  {
	var t models.User
	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(writer, "Error in received data " + err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(writer, "Email is required!!!", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(writer, "The password must be at least 6 characters!!!", 400)
		return
	}
	_, founded, _ := db.CheckUserExist(t.Email)
	if founded == true {
		http.Error(writer, "There is a user registered with that email in the database!!!", 400)
		return
	}
	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(writer, "An error occurred while trying to register the user!!!" + err.Error(), 400)
		return
	}
	if status == false {
		http.Error(writer, "The user record could not be inserted into the database!!!", 400)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
