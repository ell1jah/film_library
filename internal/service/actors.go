package service

import (
	"github.com/ell1jah/film_library/internal/models"
)

type ActorsRepo interface {
	InsertActor(actor *models.Actor) (int, error)
	UpdateActor(actor *map[string]any) error
	DeleteActor(id int) error
	DeleteActorFilms(actorID int) error
	GetActors() (*[]models.Actor, error)
}

type ActorsService struct {
	repo ActorsRepo
}

func NewActorsService(repo ActorsRepo) *ActorsService {
	return &ActorsService{
		repo: repo,
	}
}

func (s *ActorsService) GetActors() (*[]models.Actor, error) {
	return s.repo.GetActors()
}

func (s *ActorsService) DeleteActor(actorID int) error {
	return s.repo.DeleteActor(actorID)
}

func (s *ActorsService) InsertActor(actor *models.Actor) (int, error) {
	return s.repo.InsertActor(actor)
}

func (s *ActorsService) UpdateActor(actor *map[string]any) error {
	return s.repo.UpdateActor(actor)
}
