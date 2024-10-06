package main

import (
	"fmt"
	"time"

	"github.com/kramerProject/deputies-chamber/application"
	deputies_client "github.com/kramerProject/deputies-chamber/client/deputies_client"
	httpclient "github.com/kramerProject/deputies-chamber/pkg/http_client"
)

func main() {
	deputiesURL := "https://dadosabertos.camara.leg.br/api/v2/deputados?ordem=ASC&ordenarPor=nome"

	httpClient := httpclient.NewHTTPClient(120 * time.Second)

	deputiesClient := deputies_client.NewClient(httpClient, deputiesURL)
	service := application.NewService(deputiesClient)

	result, _ := service.GetAll()
	fmt.Println("result", result)

}
