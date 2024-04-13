package parser

import (
	"fmt"
	"github.com/BobrePatre/kozodoy-parser/internal/models"
	"github.com/xuri/excelize/v2"
	"io"
	"log/slog"
	"strconv"
	"strings"
)

type Repository interface {
	InsertIntoMenu(menuType string, categories []models.Category) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

const (
	menuTabKey = "меню для кассы"
)

func (s *Service) Parse(fileReader io.Reader, menuType string) error {
	excelReader, err := excelize.OpenReader(fileReader)
	defer func(excelReader *excelize.File) {
		err := excelReader.Close()
		if err != nil {
			slog.Error("can not close excelize reader", "err", err)
		}
	}(excelReader)
	if err != nil {
		return err
	}

	var categories []models.Category

	rows, err := excelReader.GetRows(menuTabKey)
	if err != nil {
		return err
	}
	currentIndex := -1
	for _, row := range rows {
		if checkRowIsEmpty(row) {
			continue
		}
		if strings.Contains(row[1], "Меню на") {
			continue
		}
		slog.Debug(fmt.Sprint(row))

		if row[0] == "" && row[1] != "" {
			categories = append(categories, models.Category{
				Title: row[1],
			})
			currentIndex++
		}
		if row[0] != "" {
			if len(categories) <= 0 {
				return models.MenuNotValidErr
			}
			price, _ := strconv.ParseFloat(row[2], 64)
			categories[currentIndex].Dishes = append(categories[currentIndex].Dishes, models.Dish{
				Weight: strings.TrimSpace(row[0]),
				Title:  strings.TrimSpace(row[1]),
				Price:  price,
			})
		}
	}
	slog.Debug("parsed data", "categories", categories)
	return nil
}

func checkRowIsEmpty(row []string) bool {
	if row == nil || len(row) == 0 {
		return true
	}
	if row[0] == "" && row[1] == "" && row[2] == "" {
		return true
	}
	return false
}
