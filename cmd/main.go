package main

import (
	"log"
	"net/http"

	"github.com/IcaroSilvaFK/picpay/cmd/routes"
)

func main() {

	mx := http.NewServeMux()

	mx.HandleFunc("GET /heathy", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		w.Write([]byte(`{"ok": true}`))
	})

	routes.NewApiRouter(mx)

	log.Println("Server running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mx))
}
