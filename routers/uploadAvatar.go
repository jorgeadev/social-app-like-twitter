package routers

import (
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"io"
	"net/http"
	"os"
	"strings"
)

/* UploadAvatar upload the avatar image at server */
func UploadAvatar(writer http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var namefile string = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(namefile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(writer, "Error uploading file avatar!!! " + err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(writer, "Error copying file avatar!!! " + err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension
	status, err = db.ModifyRegister(user, IDUser)
	if err != nil || status == false {
		http.Error(writer, "Error saving avatar file in database!!! ", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}