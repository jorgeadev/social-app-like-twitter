package routers

import (
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"net/http"
)

/**/
func DownRelationShip(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "The id parameter is mandatory!!! ", http.StatusBadRequest)
		return
	}

	var t models.RelationShip
	t.UserID = IDUser
	t.UserRelationShipID = ID

	status, err := db.DeleteRelationShip(t)
	if err != nil {
		http.Error(writer, "An error occurred when trying to delete the relation!!! " + err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(writer, "The relationship could not be deleted!!! ", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}
