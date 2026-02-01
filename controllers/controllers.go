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

// AllPersonalities returns all personalities from the database json
func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p []models.Personality	
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

//get personality by id
func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

//post new personality
func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Create(&personality)
	json.NewEncoder(w).Encode(personality)
}

//delete personality by id
func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(personality)
}