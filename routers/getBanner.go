package routers

import (
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"io"
	"net/http"
	"os"
)

/* GetBanner send the banner to http */
func GetBanner(writer http.ResponseWriter, request *http.Request) {
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

	OpenFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(writer, "Banner not found!!! ", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(writer, OpenFile)
	if err != nil {
		http.Error(writer, "Error copying the Banner!!! ", http.StatusBadRequest)
	}
}
