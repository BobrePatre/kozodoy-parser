package parser

import (
	"github.com/BobrePatre/kozodoy-parser/internal/models"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) InsertIntoMenu(menuType string, categories []models.Category) error {
	return nil
}
