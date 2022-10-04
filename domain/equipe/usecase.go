package equipe

import (
	"github.com/TheSamuelVitor/api-go-postgres/config/database"
	modelApresentacao "github.com/TheSamuelVitor/api-go-postgres/domain/equipe/model"
	"github.com/TheSamuelVitor/api-go-postgres/infra/equipe"

	"github.com/gin-gonic/gin"
)

func NovaEquipe(c *gin.Context, req *modelApresentacao.ReqEquipe) {
	
	// abre a conexao com o banco de dados
	db := database.Conectar()

	// fecha a conexao ao fim da funcao
	defer db.Close()

	equipeRepo := equipe.NovoRepo(db)

	str := *req.Nome
	str = "Equipe: " + str

	req.Nome = &str

	equipeRepo.NovaEquipe(req)
}