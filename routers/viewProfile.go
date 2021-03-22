package routers

import (
	"encoding/json"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"net/http"
)

/* ViewProfile allows you to extract the values from the profile  */
func ViewProfile(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "You must send the id parameter ", http.StatusBadRequest)
		return
	}
	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(writer, "An error occurred while trying to find the record "+err.Error(), 400)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(profile)
}
