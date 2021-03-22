package middlewares

import (
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/routers"
	"net/http"
)

/**/
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _, _, err := routers.ProcessToken(request.Header.Get("Authorization"))
		if err != nil {
			http.Error(writer, "Token error!!! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
