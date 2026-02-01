package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thomasaqx/go-rest-api/database"
	"github.com/thomasaqx/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//função home que responde na rota "/" e exibe uma mensagem de boas vindas
	_, error := fmt.Fprintf(w, "Welcome to the Home Page!")
	if error != nil {
		fmt.Println("Error:", error)
	}
}

// AllPersonalities retorna todas as personalidades em formato JSON
func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)

}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personality := range models.Personalities {
		if personality.Id == id {
			json.NewEncoder(w).Encode(personality)
			return
		}
	}
}
