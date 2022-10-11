package equipe

import (
	"github.com/TheSamuelVitor/api-go-postgres/config/database"
	modelApresentacao "github.com/TheSamuelVitor/api-go-postgres/domain/equipe/model"
	"github.com/TheSamuelVitor/api-go-postgres/infra/equipe"

)

func NovaEquipe(req *modelApresentacao.ReqEquipe) {
	
	db := database.Conectar()
	defer db.Close()

	equipeRepo := equipe.NovoRepo(db)

	str := *req.Nome

	req.Nome = &str

	equipeRepo.NovaEquipe(req)
}


func MostraEquipes()  {
	
	db := database.Conectar()
	defer db.Close()

	equipeRepo := equipe.NovoRepo(db)

	res, err  = equipeRepo.ListarEquipes()

	for i := range res {

	}

}