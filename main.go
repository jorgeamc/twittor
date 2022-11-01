package main

import (
	"github.com/jorgeamc/twittor/bd"
	"github.com/jorgeamc/twittor/handlers"
	"log"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Connetion failed")
		return
	}
	handlers.Handlers()
}
