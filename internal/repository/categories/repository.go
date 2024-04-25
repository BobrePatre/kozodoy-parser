package categories

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

var _ parser.CategoryRepository = (*Repository)(nil)

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

func (r *Repository) CreateCategory(category models.Category, menuId string) (string, error) {

	reqData := struct {
		MenuId string `json:"menuId"`
		Title  string `json:"title"`
	}{
		MenuId: menuId,
		Title:  category.Title,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return "", err
	}

	slog.Debug("backend address", "address", r.networkDatacourceConfig.CoreBackendHost)
	req, err := http.NewRequest("POST", r.networkDatacourceConfig.CoreBackendHost+"/categories", bytes.NewBufferString(string(jsonData)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.provider.DoRequest(req)
	if err != nil {
		return "", err
	}

	respData := struct {
		Id string `json:"id"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return "", err
	}
	slog.Debug(resp.Status)
	return respData.Id, nil
}
