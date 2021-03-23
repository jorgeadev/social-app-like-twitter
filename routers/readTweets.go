package routers

import (
	"encoding/json"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"net/http"
	"strconv"
)

/* ReadTweets Read the tweets */
func ReadTweets(writer http.ResponseWriter, request *http.Request)  {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "You must sent the parameter id!!! ", http.StatusBadRequest)
		return
	}

	if len(request.URL.Query().Get("page")) < 1 {
		http.Error(writer, "You must sent the parameter page!!! ", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		http.Error(writer, "You must sent the parameter page with value most big to 0 !!! ", http.StatusBadRequest)
		return
	}
	pag := int64(page)
	response, correct := db.ReadTweets(ID, pag)
	if correct == false {
		http.Error(writer, "Error at read tweets !!! ", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(response)
}
