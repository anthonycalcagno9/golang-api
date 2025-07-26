package handlers

import (
	"encoding/json"
	"golangapi/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerSendMeals(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlerSendMeals)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var meals []models.Meal
	err = json.Unmarshal(rr.Body.Bytes(), &meals)
	if err != nil {
		t.Errorf("error unmarshaling response: %v", err)
	}

	if len(meals) != 3 {
		t.Errorf("expected 3 meals, got %d", len(meals))
	}

	expectedMeals := []models.Meal{meal1, meal2, meal3}
	for i, meal := range meals {
		if meal != expectedMeals[i] {
			t.Errorf("expected meal %d to be %v, got %v", i+1, expectedMeals[i], meal)
		}
	}
}
