package handlers

import (
	"encoding/json"
	"fmt"
	"golangapi/models"
	"net/http"
)

var meal1 = models.Meal{
	Breakfast:    "eggs",
	Lunch:        "sandwich",
	Dinner:       "tacos",
	Rating:       7,
	DayOfTheWeek: "Monday",
}

var meal2 = models.Meal{
	Breakfast:    "omelet",
	Lunch:        "rice",
	Dinner:       "wings",
	Rating:       8,
	DayOfTheWeek: "Tuesday",
}

var meal3 = models.Meal{
	Breakfast:    "pancakes",
	Lunch:        "fruit",
	Dinner:       "steak",
	Rating:       9,
	DayOfTheWeek: "Wednesday",
}

var xmeals = []models.Meal{meal1, meal2, meal3}

func HandlerSendMeals(w http.ResponseWriter, r *http.Request) {

	for _, value := range xmeals {
		fmt.Printf("day of the week%v\n", value.DayOfTheWeek)
		fmt.Printf("breakfast = %v\n", value.Breakfast)
		fmt.Printf("lunch = %v\n", value.Lunch)
		fmt.Printf("dinner = %v\n", value.Dinner)
		fmt.Printf("rating = %v\n", value.Rating)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	response, err := json.Marshal(xmeals)
	if err != nil {
		http.Error(w, "error marshaling xmeals", http.StatusInternalServerError)
		return
	}

	resp, err := w.Write(response)
	if err != nil {
		http.Error(w, "error writing to responseWriter", http.StatusInternalServerError)
		fmt.Printf("error writing to responseWriter: %v\n", err)
		return
	}

	fmt.Printf("response writer bytes written = %v\n", resp)

}
