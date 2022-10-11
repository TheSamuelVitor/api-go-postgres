package postgres

import (
	"database/sql"

	modelData "github.com/TheSamuelVitor/api-go-postgres/infra/equipe/model"
	modelApresentacao "github.com/TheSamuelVitor/api-go-postgres/domain/equipe/model"
)

type DBEquipe struct {
	DB *sql.DB
}

func (postgres *DBEquipe) NovaEquipe(req *modelData.Equipe) (*modelApresentacao.ReqEquipe, error) {

	sqlStatement := `insert into equipes(nome) values($1::TEXT) returning *`

	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome)
	err := row.Scan(&equipe.Nome)
	if err != nil {
		return nil, err
	}
	return equipe, nil
}

func (postgres *DBEquipe) MostraEquipes() ([]modelApresentacao.ReqEquipe, error) {

	sqlStatement := "select * from equipes order by nome"
	var res = []modelApresentacao.ReqEquipe{}
	var equipe = modelApresentacao.ReqEquipe{}

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&equipe.Nome)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
	
		res = append(res, equipe)
	}
	return res, nil

}

func (postgres *DBEquipe) BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error) {
	sqlStatement := `select * from equipes where id_equipe = $1`
	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, id)
	err := row.Scan(&equipe.Nome)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	return equipe, nil
}