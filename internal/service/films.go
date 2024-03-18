package service

import (
	"github.com/ell1jah/film_library/internal/models"
)

type FilmsRepo interface {
	InsertFilm(film *models.Film) (int, error)
	UpdateFilm(film *map[string]any) error
	DeleteFilm(filmID int) error
	GetFilms(attrs *map[string]any) (*[]models.Film, error)
	InsertFilmActors(filmID int, actors []models.Actor) error
	DeleteFilmActors(filmID int) error
}

type FilmsService struct {
	repo FilmsRepo
}

func NewFilmsService(repo FilmsRepo) *FilmsService {
	return &FilmsService{
		repo: repo,
	}
}

func (s *FilmsService) InsertFilm(film *models.Film) (int, error) {
	return s.repo.InsertFilm(film)
}

func (s *FilmsService) UpdateFilm(film *map[string]any) error {
	return s.repo.UpdateFilm(film)
}

func (s *FilmsService) DeleteFilm(filmID int) error {
	return s.repo.DeleteFilm(filmID)
}

func (s *FilmsService) GetFilms(attrs *map[string]any) (*[]models.Film, error) {
	return s.repo.GetFilms(attrs)
}
