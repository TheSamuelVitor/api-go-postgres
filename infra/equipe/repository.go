package equipe

import (
	"database/sql"

	modelApresentacao "github.com/TheSamuelVitor/api-go-postgres/domain/equipe/model"
	modelData "github.com/TheSamuelVitor/api-go-postgres/infra/equipe/model"
	"github.com/TheSamuelVitor/api-go-postgres/infra/equipe/postgres"
)

type repositorio struct {
	Data *postgres.DBEquipe
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBEquipe{DB: novoDB},
	}
}

func (r *repositorio) NovaEquipe(req *modelApresentacao.ReqEquipe) {
	r.Data.NovaEquipe(&modelData.Equipe{
		Nome: req.Nome,
	})
}

func (r * repositorio) ListarEquipes() {

}