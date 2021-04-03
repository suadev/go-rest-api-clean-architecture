package main

import (
	"github.com/suadev/go-rest-api-clean-architecture/config"
	"github.com/suadev/go-rest-api-clean-architecture/entity"
	product "github.com/suadev/go-rest-api-clean-architecture/internal"
	"github.com/suadev/go-rest-api-clean-architecture/server"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config := config.LoadConfig(".")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDBURL(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to the DB.")
	}

	db.AutoMigrate(&entity.Product{})

	repo := product.NewRepository(db)
	service := product.NewService(repo)
	handler := product.NewHandler(service)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}
}
