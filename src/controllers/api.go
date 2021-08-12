package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ranon-rat/proyectoClaseProgramacion/src/data"
)

func SendApi(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		var food data.Food
		json.NewDecoder(r.Body).Decode(&food)
		data.AddFood(food)

	case "GET":
		apiChan := make(chan data.Api)
		go data.MakeApi(apiChan)
		if err := json.NewEncoder(w).Encode(<-apiChan); err != nil {
			http.Error(w, "something is wrong", http.StatusInternalServerError)
		}

	}

}
