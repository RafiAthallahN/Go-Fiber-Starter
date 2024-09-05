package apps

import (
	"go-fiber/starter/backend/handlers"
	"go-fiber/starter/backend/repositories"
	"go-fiber/starter/backend/services"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupPeople(db *gorm.DB, log *logrus.Logger) *handlers.PeopleHandler {
	peopleRepo := repositories.NewPeopleRepo(db, log)
	peopleService := services.NewPeopleService(peopleRepo, log)
	return handlers.NewPeopleHandler(peopleService, log)
}
