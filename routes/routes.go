package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thomasaqx/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()		   								 //criando uma nova rota com mux
	r.HandleFunc("/", controllers.Home)                          // quando for feito uma requisição para a rota "/", chama a função Home do controllers
	r.HandleFunc("/api/personalities", controllers.AllPersonalities) // quando for feito uma requisição para a rota "/api/personalities", chama a função AllPersonalities do controllers
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById) // rota para obter personalidade por ID
	log.Fatal(http.ListenAndServe(":8081", r))                   //subindo o servidor na porta 8080
}
