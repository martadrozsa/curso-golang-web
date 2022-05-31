package routes

import (
	"github.com/martadrozsa/curso-golang-web/apigorest/controllers"
	"log"
	"net/http"

)

func HandleRequest()  {

	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

