package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/SeongHyunHan/Tugether/models"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:Pa55w0rD!@localhost:5432/Tugether"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "public",
			SingularTable: false,
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Accounts{})

	return db
}
