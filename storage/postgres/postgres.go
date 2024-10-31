package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kramerProject/deputies-chamber/application"
	_ "github.com/lib/pq" // Import do driver do PostgreSQL
)

type DeputyDB struct {
	db *sql.DB
}

// Função para conectar ao banco de dados e retornar um `DeputyDB`
func NewDeputyDB(connStr string) (*DeputyDB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DeputyDB{db: db}, nil
}

// Método para fechar a conexão com o banco de dados
func (d *DeputyDB) Close() error {
	return d.db.Close()
}

// Método para salvar um único deputado no banco de dados
func (d *DeputyDB) SaveDeputy(deputy application.Deputy) error {
	fmt.Println("saving", deputy)
	_, err := d.db.Exec(`
        INSERT INTO deputies (deputy_id, uri, name, party_label, party_url, state, legislature_id, picture_url, email)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT (deputy_id) DO NOTHING`,
		deputy.ID, deputy.URI, deputy.Name, deputy.PartyLabel, deputy.PartyURL, deputy.State, deputy.LegislatureID, deputy.PictureURL, deputy.Email)

	if err != nil {
		return fmt.Errorf("erro ao inserir o deputy %d: %v", deputy.ID, err)
	}
	return nil
}

// Método para salvar uma lista de deputados no banco de dados
func (d *DeputyDB) SaveDeputies(deputies application.Deputies) error {
	for _, deputy := range deputies.DeputiesList {
		if err := d.SaveDeputy(deputy); err != nil {
			log.Printf("Erro ao salvar deputado %d: %v", deputy.ID, err)
		}
	}
	d.Close()
	return nil
}
