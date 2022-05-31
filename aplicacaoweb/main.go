package main

import (
	"fmt"
	"net/http"
)

// handler é uma função que criaremos para executar determinada ação ao acessarmos o caminho http://localhost:8080
// O primeiro parâmetro será responsável pela resposta da requisição, enquanto o segundo será responsável pelo tratamento da mesma.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World</h1>\n")
}

func main() {
	// lidar com as requisições
	http.HandleFunc("/hello", handler)

	// especificar em qual porta rodará a nossa aplicação e como lidaremos com o servidor da aplicação
	http.ListenAndServe(":8080", nil)
}
