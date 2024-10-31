package application

import (
	"errors"
	"fmt"
)

type DeputiesService struct {
	DeputiesClient  DeputiesClientInterface
	DeputiesStorage DeputiesStorageInterface
}

func NewService(deputiesClient DeputiesClientInterface, deputiesStorage DeputiesStorageInterface) *DeputiesService {
	return &DeputiesService{
		DeputiesClient:  deputiesClient,
		DeputiesStorage: deputiesStorage,
	}
}

func (ds *DeputiesService) GetAll() (Deputies, error) {
	fmt.Println("get alllll")
	deputies, err := ds.DeputiesClient.GetAll()
	if err != nil {
		fmt.Println("error getting deputies")
		return Deputies{}, errors.New("error returning deputies")
	}

	err = ds.DeputiesStorage.SaveDeputies(deputies)
	if err != nil {
		fmt.Println("erro", err)
		return Deputies{}, errors.New("error inserting deputies")
	}

	return deputies, nil
}
