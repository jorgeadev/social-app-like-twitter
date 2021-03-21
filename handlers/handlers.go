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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}