package parser

import (
	"github.com/BobrePatre/kozodoy-parser/internal/models"
	"github.com/xuri/excelize/v2"
	"io"
	"log/slog"
	"strconv"
	"strings"
)

type MenuRepository interface {
	CreateMenu(title string, menuType string, dateTo string) (string, error)
	GetMenuByType(menuType string) (string, error)
	UpdateMenu(menuId string, dateTo string) error
	ClearMenu(menuId string) error
}

type CategoryRepository interface {
	CreateCategory(category models.Category, menuId string) (string, error)
}

type DishRepository interface {
	CreateDish(dish models.Dish, categoryId string) error
}

type Service struct {
	menuRepository     MenuRepository
	categoryRepository CategoryRepository
	dishRepository     DishRepository
}

func NewService(
	menuRepository MenuRepository,
	categoryRepository CategoryRepository,
	dishRepository DishRepository,
) *Service {
	return &Service{
		menuRepository:     menuRepository,
		categoryRepository: categoryRepository,
		dishRepository:     dishRepository,
	}
}

const (
	menuTabKey       = "меню для кассы"
	menuTypeToday    = "today"
	menuTypeTomorrow = "tomorrow"
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

	title := map[string]string{
		menuTypeToday:    "На сегодня",
		menuTypeTomorrow: "На завтра",
	}

	menu := models.Menu{
		Title: title[menuType],
	}

	rows, err := excelReader.GetRows(menuTabKey)
	if err != nil {
		return err
	}

	rowsLen := len(rows)
	currentIndex := -1
	for i, row := range rows {
		if checkRowIsEmpty(row) {
			continue
		}

		if (i == rowsLen-1) && (i == rowsLen-2) && (i == rowsLen-3) {
			continue
		}

		if strings.Contains(row[1], "Меню на") {
			menu.DateTo = strings.Replace(row[1], "Меню на", "", 1)
			continue
		}

		if row[0] == "" && row[1] != "" {
			menu.Categories = append(menu.Categories, models.Category{
				Title: row[1],
			})
			currentIndex++
		}
		if row[0] != "" {
			if len(menu.Categories) <= 0 {
				return models.MenuNotValidErr
			}
			if strings.Contains(row[2], "/") {
				slog.Debug("row", "row", row)
				prices := strings.Split(row[2], "/")
				weights := strings.Split(row[0], "/")
				slog.Debug("Данные дробные", "weights", weights, "prices", prices)
				price1, _ := strconv.ParseFloat(prices[0], 64)
				price2, _ := strconv.ParseFloat(prices[1], 64)
				menu.Categories[currentIndex].Dishes = append(menu.Categories[currentIndex].Dishes, models.Dish{
					Title:  strings.TrimSpace(row[1]),
					Weight: weights[0],
					Price:  price1,
				})
				menu.Categories[currentIndex].Dishes = append(menu.Categories[currentIndex].Dishes, models.Dish{
					Title:  strings.TrimSpace(row[1]),
					Weight: weights[1],
					Price:  price2,
				})
				continue
			}
			price, _ := strconv.ParseFloat(row[2], 64)
			menu.Categories[currentIndex].Dishes = append(menu.Categories[currentIndex].Dishes, models.Dish{
				Weight: strings.TrimSpace(row[0]),
				Title:  strings.TrimSpace(row[1]),
				Price:  price,
			})
		}
	}
	//slog.Debug("parsed data", "categories", menu.Categories)

	remoteMenuId, err := s.menuRepository.GetMenuByType(menuType)
	if err != nil {
		slog.Info("can not get remote menu id", "err", err)
		remoteMenuId, err = s.menuRepository.CreateMenu(menu.Title, menuType, "s")
		if err != nil {
			slog.Error("can not create menu by type", "err", err)
			return err
		}
		slog.Info("created menu by type", "type", menuType)
	}
	menu.Id = remoteMenuId
	slog.Debug("remote Data", "menuId", remoteMenuId)
	err = s.menuRepository.UpdateMenu(menu.Id, menu.DateTo)
	if err != nil {
		slog.Error("can not update remote menu", "err", err)
		return err
	}
	for _, category := range menu.Categories {
		createdCategoryId, err := s.categoryRepository.CreateCategory(category, menu.Id)
		if err != nil {
			slog.Error("can not create category", "err", err)
			continue
		}
		for _, dish := range category.Dishes {
			err := s.dishRepository.CreateDish(dish, createdCategoryId)
			if err != nil {
				slog.Error("error when create dish", "err", err)
				continue
			}
		}
	}

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
