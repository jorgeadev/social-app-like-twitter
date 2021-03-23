package routers

import (
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"net/http"
)

/* DeleteTweet allow delete an specific tweet */
func DeleteTweet(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "You must send the id parameter!!! ", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(writer, "An error occurred while trying to delete the tweet!!! " + err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}
