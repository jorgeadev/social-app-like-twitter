package routers

import (
	"encoding/json"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"net/http"
)

/* ConsultRelationShip check if there is a relation between two users */
func ConsultRelationShip(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "The id parameter is mandatory!!! ", http.StatusBadRequest)
		return
	}

	var t models.RelationShip
	t.UserID = IDUser
	t.UserRelationShipID = ID

	var resp models.ResponseConsultRelation
	status, err := db.ConsultRelations(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(resp)
}
