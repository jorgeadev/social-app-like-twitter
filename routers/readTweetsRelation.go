package routers

import (
	"encoding/json"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"net/http"
	"strconv"
)

/**/
func ReadTweetsFollowers(writer http.ResponseWriter, request *http.Request) {
	if len(request.URL.Query().Get("page")) < 1 {
		http.Error(writer, "You must send the page parameter!!! ", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		http.Error(writer, "You must send the page parameter with int more great than 0 !!! ", http.StatusBadRequest)
		return
	}

	response, correct := db.ReadTweetsFollowers(IDUser, page)
	if correct == false {
		http.Error(writer, "Error reading tweets!!! ", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(response)
}
