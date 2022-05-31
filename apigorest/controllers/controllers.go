package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/martadrozsa/curso-golang-web/apigorest/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>Home Page</h1>")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request)  {
	err := json.NewEncoder(w).Encode(models.Personalidades)
	if err != nil {
		return
	}
}