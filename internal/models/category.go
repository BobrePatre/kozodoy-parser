package models

type (
	Category struct {
		Id     string
		Title  string
		Dishes []Dish
	}
)
