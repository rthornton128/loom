package config

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rthornton128/loom/pkg/routing"
)

// TODO: temporary, testing handlers
func handleUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"users": "index"})
}

func handleUserShow(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"users": ":id"})
}

func Draw(r *routing.Router) {
	fmt.Println("setting up routes")
	r.Get("/users/:id").With(handleUserShow)
	r.Get("/users").With(handleUsers)
}
