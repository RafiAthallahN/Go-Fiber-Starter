package services

import (
	"go-fiber/starter/backend/entities"
	"go-fiber/starter/backend/repositories"

	"github.com/sirupsen/logrus"
)

type PeopleService struct {
	repo *repositories.PeopleRepo
	log  *logrus.Logger
}

func NewPeopleService(repo *repositories.PeopleRepo, log *logrus.Logger) *PeopleService {
	return &PeopleService{repo, log}
}

func (s *PeopleService) GetPeople(page, limit int) ([]entities.People, int, error) {
	return s.repo.GetPeople(page, limit)
}

func (s *PeopleService) GetPerson(id int) (entities.People, error) {
	return s.repo.GetPerson(id)
}

func (s *PeopleService) InsertPeople(data []entities.People) ([]entities.People, error) {
	return s.repo.InsertPeople(data)
}

func (s *PeopleService) InsertPerson(data entities.People) (entities.People, error) {
	return s.repo.InsertPerson(data)
}

func (s *PeopleService) UpdatePerson(data entities.People) (entities.People, error) {
	return s.repo.UpdatePerson(data)
}

func (s *PeopleService) DeletePerson(id int) error {
	return s.repo.DeletePerson(id)
}
