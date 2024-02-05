package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/CristianSch248/App_GoLang/configs"
)

var DB *gorm.DB

func OpenConn() (*gorm.DB, error) {
	conf := configs.GetDB()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Configurar o banco de dados para usar singular table name (ex: 'User' será 'user' na tabela)
	DB.NamingStrategy = schema.NamingStrategy{
		SingularTable: true,
	}

	return DB, nil
}

func Migrate(models ...interface{}) {
	DB.AutoMigrate(models...)
}
