package routers

import (
	"encoding/json"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"net/http"
	"time"
)

/**/
func RecordTweet(writer http.ResponseWriter, request *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(request.Body).Decode(&message)

	register := models.RecordTweet{
		UserID: IDUser,
		Message: message.Message,
		DateTweet: time.Now(),
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(writer, "An error occurred while inserting the record, please try again!!!" + err.Error(), 400)
		return
	}

	if status == false {
		http.Error(writer, "The tweet could not be inserted !!!", 400)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}