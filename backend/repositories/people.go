package repositories

import (
	"errors"
	"go-fiber/starter/backend/entities"
	"go-fiber/starter/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PeopleRepo struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewPeopleRepo(db *gorm.DB, log *logrus.Logger) *PeopleRepo {
	return &PeopleRepo{db, log}
}

func (r *PeopleRepo) GetPeople(page, limit int) ([]entities.People, int, error) {

	var people []entities.People

	var id int64

	if errGetAllPeople := r.db.Preload("people", func(db *gorm.DB) *gorm.DB {
		return db.Count(&id)
	}).Scopes(utils.NewPagination(nil, &page, &limit, nil).PaginationQuery).
		Where("deleted_at IS NULL").
		Find(&people).Error; errGetAllPeople != nil {
		r.log.Errorf("Error on Repo: %v", errGetAllPeople)
		return nil, 0, errors.New("technical issue happen, please try again after a while")
	}

	return people, int(id), nil
}

func (r *PeopleRepo) GetPerson(id int) (entities.People, error) {

	var person entities.People

	if errGetPerson := r.db.Where("id = ?", id).Find(&person).Error; errGetPerson != nil {
		r.log.Errorf("Error on Repo: %v", errGetPerson)
		return entities.People{}, errors.New("technical issue happen, please try again after a while")
	}

	return person, nil
}

func (r *PeopleRepo) InsertPerson(data entities.People) (entities.People, error) {

	if errInsertPerson := r.db.Save(&data).Error; errInsertPerson != nil {
		r.log.Errorf("Error on Repo: %v", errInsertPerson)
		return entities.People{}, errors.New("technical issue happen, please try again after a while")
	}

	return data, nil
}

func (r *PeopleRepo) InsertPeople(data []entities.People) ([]entities.People, error) {

	if errInsertPeople := r.db.CreateInBatches(&data, len(data)).Error; errInsertPeople != nil {
		r.log.Errorf("Error on Repo: %v", errInsertPeople)
		return nil, errors.New("technical issue happen, please try again after a while")
	}

	return data, nil
}

func (r *PeopleRepo) UpdatePerson(data entities.People) (entities.People, error) {

	if errUpdatePerson := r.db.Where("id = ?", data.ID).Updates(&data).Error; errUpdatePerson != nil {
		r.log.Errorf("Error on Repo: %v", errUpdatePerson)
		return entities.People{}, errors.New("technical issue happen, please try again after a while")
	}

	return data, nil
}

func (r *PeopleRepo) DeletePerson(id int) error {

	if errDeletePerson := r.db.Where("id = ?", id).Delete(&entities.People{}).Error; errDeletePerson != nil {
		r.log.Errorf("Error on Repo: %v", errDeletePerson)
		return errors.New("technical issue happen, please try again after a while")
	}
	return nil
}
