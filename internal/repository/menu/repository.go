package menu

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/BobrePatre/kozodoy-parser/internal/config/datasources"
	"github.com/BobrePatre/kozodoy-parser/internal/providers/service_access"
	"github.com/BobrePatre/kozodoy-parser/internal/service/parser"
	"github.com/goccy/go-json"
	"io"
	"log/slog"
	"net/http"
)

var _ parser.MenuRepository = (*Repository)(nil)

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

func (r *Repository) CreateMenu(title string, menuType string) (string, error) {
	reqData := struct {
		Title    string `json:"title"`
		MenuType string `json:"type"`
	}{
		Title:    title,
		MenuType: menuType,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return "", err
	}

	slog.Debug("backend address", "address", r.networkDatacourceConfig.CoreBackendHost)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", r.networkDatacourceConfig.CoreBackendHost, "menu"), bytes.NewBufferString(string(jsonData)))
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
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return "", err
	}

	slog.Debug("backend response create menu", "id", respData.Id)
	return respData.Id, nil
}

func (r *Repository) GetMenuByType(menuType string) (string, error) {

	slog.Debug("backend address", "address", r.networkDatacourceConfig.CoreBackendHost)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", r.networkDatacourceConfig.CoreBackendHost, "menu/type", menuType), nil)
	if err != nil {
		slog.Error("error when creating http req to get menu by type", "err", err, "address", fmt.Sprintf("%s/%s/%s", r.networkDatacourceConfig.CoreBackendHost, "menu/type", menuType))
		return "", err
	}
	slog.Debug("lalal", "address", fmt.Sprintf("%s/%s/%s", r.networkDatacourceConfig.CoreBackendHost, "menu/type", menuType))

	req.Header.Set("Content-Type", "application/json")

	resp, err := r.provider.DoRequest(req)
	if err != nil {
		return "", err
	}
	respBody, err := io.ReadAll(resp.Body)

	data := struct {
		Id string `json:"id"`
	}{}

	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return "", err
	}

	if data.Id == "" {
		return "", errors.New("menu not found")
	}

	slog.Debug("created menu resp status", "resp", data.Id)
	return data.Id, nil
}

func (r *Repository) ClearMenu(menuId string) error {
	//TODO implement me
	slog.Error("ClearMenu NOT IMPLEMENTED!!!")
	return nil
}
