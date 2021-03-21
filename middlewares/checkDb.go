package middlewares

import (
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"net/http"
)

/* CheckDB is the middleware to know the status of database */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(writer, "Lost connection with database!!!", 500)
			return
		}
		next .ServeHTTP(writer, request)
	}
}