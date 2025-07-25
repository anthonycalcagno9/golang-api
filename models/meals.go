package models

type Meal struct {
	Breakfast    string `json:"breakfast"`
	Lunch        string `json:"lunch"`
	Dinner       string `json:"dinner"`
	Rating       int    `json:"rating"`
	DayOfTheWeek string `json:"dayoftheweek"`
}
