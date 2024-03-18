package main

import (
	"github.com/ell1jah/film_library/internal/handler"
	"github.com/ell1jah/film_library/internal/repository"
	"github.com/ell1jah/film_library/internal/service"
	"github.com/sirupsen/logrus"
)

// @title REST API for VK
// @version 1.0
// @description API Server for VK Intern test task

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	db, err := repository.NewPostgresDB(repository.ReadConfig())
	if err != nil {
		logrus.Fatalf("error connecting database: %s", err)
	}
	repo := repository.NewRepository(db)
	s := service.NewService(repo)

	h := handler.NewHandler("8080", s)
	h.InitRoutes()
}
