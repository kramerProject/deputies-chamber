package deputies_client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/kramerProject/deputies-chamber/application"
	httpclient "github.com/kramerProject/deputies-chamber/pkg/http_client"
)

type deputiesClient struct {
	httpClient  httpclient.HTTPClient
	deputiesURL string
}

func NewClient(client httpclient.HTTPClient, deputiesURL string) *deputiesClient {
	return &deputiesClient{
		httpClient:  client,
		deputiesURL: deputiesURL,
	}
}

func (dc *deputiesClient) GetAll() (application.Deputies, error) {
	var deputies application.Deputies
	url := dc.deputiesURL

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return application.Deputies{}, err
	}

	res, err := dc.httpClient.Do(req)
	if err != nil {
		return application.Deputies{}, err
	}

	defer res.Body.Close()

	// Lendo o corpo da resposta
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Erro ao ler o corpo da resposta: %v", err)
	}

	// Fazendo o parsing do JSON para a struct Deputies
	err = json.Unmarshal(body, &deputies)
	if err != nil {
		log.Fatalf("Erro ao fazer parsing do JSON: %v", err)
		return application.Deputies{}, err
	}

	return deputies, nil
}
