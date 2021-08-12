package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ranon-rat/proyectoClaseProgramacion/src/data"
)

func MakePurchases(rw http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["buy"]
	if !ok || len(keys) < 1 {
		rw.Write([]byte("you doesnt buy something"))
	}
	var items []data.Food
	for _, idString := range keys {
		idInt, _ := strconv.Atoi(idString)
		items = append(items, data.GetBuyFood(idInt))

	}
	json.NewEncoder(rw).Encode(items)
}
