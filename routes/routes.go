package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thomasaqx/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()                                                                   //criando uma nova rota com mux
	r.HandleFunc("/", controllers.Home)                                                    // quando for feito uma requisição para a rota "/", chama a função Home do controllers
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")        // quando for feito uma requisição para a rota "/api/personalities", chama a função AllPersonalities do controllers
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("Get") // rota para obter personalidade por ID
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")    // rota para criar uma nova personalidade
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete") // rota para deletar uma personalidade por ID
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")      // rota para editar uma personalidade por ID
	log.Fatal(http.ListenAndServe(":8080", r))                                             //subindo o servidor na porta 8080
}
