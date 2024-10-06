package application

type DeputyServiceInterface interface {
	GetAll() (Deputies, error)
}

type DeputiesClientInterface interface {
	GetAll() (Deputies, error)
}

type Deputies struct {
	DeputiesList []Deputy `json:"dados"`
}

type Deputy struct {
	ID            int    `json:"id"`
	URI           string `json:"uri"`
	Name          string `json:"nome"`
	PartyLabel    string `json:"siglaPartido"`
	PartyURL      string `json:"uriPartido"`
	State         string `json:"siglaUf"`
	LegislatureID int    `json:"idLegislatura"`
	PictureURL    string `json:"urlFoto"`
	Email         string `json:"email"`
}
