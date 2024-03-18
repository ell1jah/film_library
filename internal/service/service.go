package service

import (
	"github.com/ell1jah/film_library/internal/models"
	"github.com/ell1jah/film_library/internal/repository"
)

type Actors interface {
	DeleteActor(actorID int) error
	InsertActor(actor *models.Actor) (int, error)
	UpdateActor(actor *map[string]any) error
	GetActors() (*[]models.Actor, error)
}

type Films interface {
	InsertFilm(film *models.Film) (int, error)
	UpdateFilm(film *map[string]any) error
	DeleteFilm(filmID int) error
	GetFilms(attrs *map[string]any) (*[]models.Film, error)
}

type Authorization interface {
	CreateAdmin(user *models.Admin) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
	Films
	Actors
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Films:         NewFilmsService(repo),
		Actors:        NewActorsService(repo),
	}
}
