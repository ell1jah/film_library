package repository

import (
	"github.com/ell1jah/film_library/internal/models"
	"github.com/jmoiron/sqlx"
)

type Actors interface {
	InsertActor(actor *models.Actor) (int, error)
	UpdateActor(actor *map[string]any) error
	DeleteActor(id int) error
	DeleteActorFilms(actorID int) error
	GetActors() (*[]models.Actor, error)
}

type Films interface {
	InsertFilm(film *models.Film) (int, error)
	UpdateFilm(film *map[string]any) error
	DeleteFilm(filmID int) error
	GetFilms(attrs *map[string]any) (*[]models.Film, error)
	InsertFilmActors(filmID int, actors []models.Actor) error
	DeleteFilmActors(filmID int) error
}

type Authorization interface {
	CreateAdmin(admin *models.Admin) (int, error)
	GetAdmin(adminname, password string) (*models.Admin, error)
}

type Repository struct {
	Actors
	Films
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Actors:        NewActorsPostgres(db),
		Films:         NewFilmsPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
