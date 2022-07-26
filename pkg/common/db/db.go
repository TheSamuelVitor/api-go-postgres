package db

import (
	"log"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db,err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Membro{})
	db.AutoMigrate(&models.Equipe{})
	db.AutoMigrate(&models.Tarefa{})
	db.AutoMigrate(&models.Projeto{})

	return db
}