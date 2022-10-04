package equipe

import modelApresentacao "github.com/TheSamuelVitor/api-go-postgres/domain/equipe/model"

type IEquipe interface {
	NovaEquipe(req *modelApresentacao.ReqEquipe)
	ListarEquipes()
}