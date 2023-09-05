package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/samoei/coffeeapp/services"
	"github.com/samoei/coffeeapp/services/helpers"
)

var coffee services.Coffee

// Get/coffees
func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	all, err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

// GET/coffees/coffee{id}
func GetCoffeeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	coffee, err := coffee.GetCoffeeById(id)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, coffee)
}

// UPDATE/coffees/coffee{id}
func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffee services.Coffee
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&coffee)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedCoffee, err := coffee.UpdateCoffe(id, coffee)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, updatedCoffee)

}

// DELETE/coffees/coffee{id}
func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := coffee.DeleteCoffe(id)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, nil)
}

// POST/coffees/coffee
func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	coffeeCreated, err := coffee.CreateCoffee(coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, coffeeCreated)
}
