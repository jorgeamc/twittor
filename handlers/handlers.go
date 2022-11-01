package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jorgeamc/twittor/middlew"
	"github.com/jorgeamc/twittor/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/create", middlew.CheckConnectionBd(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
