package models

type (
	Menu struct {
		Id         string
		Title      string
		Categories []Category
		DateTo     string
	}
)
