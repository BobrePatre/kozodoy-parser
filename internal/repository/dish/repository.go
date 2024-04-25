package dish

import (
	"bytes"
	"github.com/BobrePatre/kozodoy-parser/internal/config/datasources"
	"github.com/BobrePatre/kozodoy-parser/internal/models"
	"github.com/BobrePatre/kozodoy-parser/internal/providers/service_access"
	"github.com/BobrePatre/kozodoy-parser/internal/service/parser"
	"github.com/goccy/go-json"
	"log/slog"
	"net/http"
)

var _ parser.DishRepository = (*Repository)(nil)

func NewRepository(provider service_access.Provider, cfg *datasources.NetworkConfig) *Repository {
	return &Repository{
		provider:                provider,
		networkDatacourceConfig: cfg,
	}
}

type Repository struct {
	provider                service_access.Provider
	networkDatacourceConfig *datasources.NetworkConfig
}

func (r Repository) CreateDish(dish models.Dish, categoryId string) error {
	reqData := struct {
		CategoryId string  `json:"categoryId"`
		Title      string  `json:"title"`
		Weight     string  `json:"weight"`
		Price      float64 `json:"price"`
	}{
		CategoryId: categoryId,
		Title:      dish.Title,
		Weight:     dish.Weight,
		Price:      dish.Price,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", r.networkDatacourceConfig.CoreBackendHost+"/dishes", bytes.NewBufferString(string(jsonData)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.provider.DoRequest(req)
	if err != nil {
		return err
	}

	slog.Debug("CreateDish dish status", "status", resp.StatusCode)
	return nil
}
