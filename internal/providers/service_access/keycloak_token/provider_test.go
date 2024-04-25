package keycloakToken

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config/security"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestServiceAccessProvider(t *testing.T) {

	t.Run("Getting Access Token", func(t *testing.T) {
		config := security.WebAuthConfig{
			ServiceAccessSecret:   "0bpXfFQf3444S2cUAAXv5NbmWsMmsjHc",
			ServiceAccessTokenUrl: "http://localhost:8180/realms/kozodoy/protocol/openid-connect/token",
			ServiceAccessClient:   "menu-client",
		}

		provider := NewProvider(config)

		access, err := provider.GetAccess()
		if err != nil {
			t.Fatal(err)
		}
		if access.Token == "" {
			t.Fatal("access is empty", access)
		}
		t.Log("Access Token:", access.Token)
	})

	t.Run("Executing request to core backend", func(t *testing.T) {
		config := security.WebAuthConfig{
			ServiceAccessSecret:   "0bpXfFQf3444S2cUAAXv5NbmWsMmsjHc",
			ServiceAccessTokenUrl: "http://localhost:8180/realms/kozodoy/protocol/openid-connect/token",
			ServiceAccessClient:   "menu-client",
		}
		provider := NewProvider(config)

		testRequest, err := http.NewRequest("GET", "http://localhost:2000/menu/type/today", nil)
		if err != nil {
			log.Fatal(err)
		}

		response, err := provider.DoRequest(testRequest)
		if err != nil {
			t.Fatal(err)
		}
		respBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		t.Log("status code: ", response.StatusCode, "response body: ", string(respBody))

	})

}
