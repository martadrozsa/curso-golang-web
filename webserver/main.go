package main

import (
	"encoding/json"
	"net/http"
)

type Tasks struct {
	Name string
	Done bool
}

func Hello(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello World"))
}

func TaskHandler(w http.ResponseWriter, r *http.Request)  {

	task := Tasks{
		Name: "Aprender Golang",
		Done: true,
	}

	j, _ := json.Marshal(task)
	w.Write(j)
}

func main()  {

	http.HandleFunc("/", Hello)
	http.HandleFunc("/tasks", TaskHandler)
	http.ListenAndServe(":8080", nil)
}
