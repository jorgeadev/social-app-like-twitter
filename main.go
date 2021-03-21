package main

import (
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/db"
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/handlers"
	"log"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No connection to database")
		return
	}
	handlers.Handlers()
}
