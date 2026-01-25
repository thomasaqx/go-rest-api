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
	r.HandleFunc("/personalities", controllers.AllPersonalities) // quando for feito uma requisição para a rota "/personalities", chama a função AllPersonalities do controllers
	r.HandleFunc("/personalities/{id}", controllers.GetPersonalityById)
	log.Fatal(http.ListenAndServe(":8080", r))                   //subindo o servidor na porta 8080
}
