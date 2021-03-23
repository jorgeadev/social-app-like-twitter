package routers

import (
	"encoding/json"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"net/http"
	"strconv"
)

/* ListUsers rear a list of users */
func ListUsers(writer http.ResponseWriter, request *http.Request) {
	typeUser := request.URL.Query().Get("type")
	page := request.URL.Query().Get("page")
	search := request.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(writer, "You must send page parameter like int and more great than 0 !!!", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)

	result, status := db.ReadAllUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(writer, "Error reading the users !!!", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(result)
}