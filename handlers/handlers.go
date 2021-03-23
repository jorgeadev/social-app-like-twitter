package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/middlewares"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

/* Handlers set my port a server listen */
func Handlers() {
	router := mux.NewRouter()

	/* Routers */
	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modify-profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.RecordTweet))).Methods("POST")
	router.HandleFunc("/read-tweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/delete-tweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	/* Images Work */
	router.HandleFunc("/uploadAvatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlewares.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlewares.CheckDB(routers.GetBanner)).Methods("GET")

	/* Relations */
	router.HandleFunc("/upRelationShip", middlewares.CheckDB(middlewares.ValidateJWT(routers.UpRelationShip))).Methods("POST")
	router.HandleFunc("/downRelationShip", middlewares.CheckDB(middlewares.ValidateJWT(routers.DownRelationShip))).Methods("DELETE")
	router.HandleFunc("/consultRelationShip", middlewares.CheckDB(middlewares.ValidateJWT(routers.ConsultRelationShip))).Methods("GET")

	/**/
	router.HandleFunc("/listUsers", middlewares.CheckDB(middlewares.ValidateJWT(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/readTweetsFollowers", middlewares.CheckDB(middlewares.ValidateJWT(routers.ReadTweetsFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8085"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}