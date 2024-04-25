package keycloakToken

import (
	securityConfig "github.com/BobrePatre/kozodoy-parser/internal/config/security"
	"github.com/BobrePatre/kozodoy-parser/internal/providers/service_access"
	"github.com/goccy/go-json"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

var _ service_access.Provider = (*Provider)(nil)

type Access struct {
	Token string `json:"access_token"`
}

type Provider struct {
	config securityConfig.WebAuthConfig
}

func NewProvider(config securityConfig.WebAuthConfig) *Provider {
	return &Provider{config: config}
}

func (p *Provider) DoRequest(request *http.Request) (*http.Response, error) {
	access, err := p.GetAccess()
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", "Bearer "+access.Token)
	return http.DefaultClient.Do(request)
}

func (p *Provider) GetAccess() (*Access, error) {

	data := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {p.config.ServiceAccessClient},
		"client_secret": {p.config.ServiceAccessSecret},
	}

	// Создаем запрос
	req, err := http.NewRequest("POST", p.config.ServiceAccessTokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("error closing body")
		}
	}(resp.Body)

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var access Access
	err = json.Unmarshal(body, &access)
	if err != nil {
		return nil, err
	}

	// Предполагается, что токен возвращается в формате JSON, в примере не рассматривается его парсинг
	return &access, nil
}
