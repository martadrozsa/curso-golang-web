package main

import (
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() //utilizando o contexto da request
	log.Println("Iniciou minha request")
	defer log.Println("Finalizou minha request")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Request cancelada")
		w.Write([]byte("Requisição processada com sucesso"))
		case <- ctx.Done():

			http.Error(w, "Request cancelada", http.StatusRequestTimeout)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
