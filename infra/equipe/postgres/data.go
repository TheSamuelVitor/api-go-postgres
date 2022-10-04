package postgres

import (
	"database/sql"
	"fmt"

	modelData "github.com/TheSamuelVitor/api-go-postgres/infra/equipe/model"
)

type DBEquipe struct {
	DB *sql.DB
}

func (postgres *DBEquipe) NovaEquipe(req *modelData.Equipe) {

	_, err := postgres.DB.Exec(`INSERT INTO equipes(nome) VALUES($1)`, req.Nome)
	if err != nil {
		fmt.Println(err)
	}
}
