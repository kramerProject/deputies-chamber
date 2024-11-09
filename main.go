package main

import (
	"fmt"
	"time"

	"github.com/kramerProject/deputies-chamber/application"
	deputies_client "github.com/kramerProject/deputies-chamber/client/deputies_client"
	httpclient "github.com/kramerProject/deputies-chamber/pkg/http_client"
	"github.com/kramerProject/deputies-chamber/server"
	"github.com/kramerProject/deputies-chamber/storage/postgres"
)

func main() {
	fmt.Println("start")
	deputiesURL := "https://dadosabertos.camara.leg.br/api/v2/deputados?ordem=ASC&ordenarPor=nome"
	connString := "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable"

	httpClient := httpclient.NewHTTPClient(120 * time.Second)

	storage, _ := postgres.NewDeputyDB(connString)

	deputiesClient := deputies_client.NewClient(httpClient, deputiesURL)
	service := application.NewService(deputiesClient, storage)

	serv := server.MakeNewWebserver()
	serv.Service = *service
	serv.Serve()
	fmt.Println("running")

}
