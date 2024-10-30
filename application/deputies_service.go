package application

import "errors"

type DeputiesService struct {
	DeputiesClient DeputiesClientInterface
}

func NewService(deputiesClient DeputiesClientInterface) *DeputiesService {
	return &DeputiesService{DeputiesClient: deputiesClient}
}

func (ds *DeputiesService) GetAll() (Deputies, error) {
	deputies, err := ds.DeputiesClient.GetAll()
	if err != nil {
		return Deputies{}, errors.New("error returning deputies")
	}
	return deputies, nil
}
