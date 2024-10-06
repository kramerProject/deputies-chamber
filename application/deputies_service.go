package application

type DeputiesService struct {
	DeputiesClient DeputiesClientInterface
}

func NewService(deputiesClient DeputiesClientInterface) *DeputiesService {
	return &DeputiesService{DeputiesClient: deputiesClient}
}

func (ds *DeputiesService) GetAll() (Deputies, error) {
	deputies, err := ds.DeputiesClient.GetAll()
	if err != nil {
		return Deputies{}, err
	}
	return deputies, nil
}
