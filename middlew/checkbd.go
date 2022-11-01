package middlew

import (
	"github.com/jorgeamc/twittor/bd"
	"net/http"
)

func CheckConnectionBd(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Connections failed with MongoDB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
