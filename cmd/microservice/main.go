package main

import (
	"log"

	"github.com/bquenin/microservice/api"
	"github.com/bquenin/microservice/cmd/microservice/config"
	"github.com/bquenin/microservice/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	postgres, err := database.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password)
	if err != nil {
		log.Fatal(err.Error())
	}

	authorService := api.NewAuthorService(postgres)

	router := gin.Default()
	router.POST("/authors", authorService.Create)
	router.GET("/authors", authorService.List)

	router.Run()
}
