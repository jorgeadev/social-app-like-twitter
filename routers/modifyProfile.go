package routers

import (
	"encoding/json"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"net/http"
)

/**/
func ModifyProfile(writer http.ResponseWriter, request *http.Request) {
	var t models.User

	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(writer, "Incorrect data "+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.ModifyRegister(t, IDUser)
	if err != nil {
		http.Error(writer, "An error has occurred to intent modify the register "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(writer, "Failed to modify user registry ", 400)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
