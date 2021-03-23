package routers

import (
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"io"
	"net/http"
	"os"
)

/* GetAvatar send the banner to http */
func GetAvatar(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "You must send the id parameter!!!", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(writer, "User not found!!! ", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(writer, "Avatar not found!!! ", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(writer, OpenFile)
	if err != nil {
		http.Error(writer, "Error copying the avatar!!! ", http.StatusBadRequest)
	}
}
