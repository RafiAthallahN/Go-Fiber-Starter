package routes

import (
	"go-fiber/starter/backend/handlers"

	"github.com/gofiber/fiber/v2"
)

type PeopleRoutes struct {
	App    *fiber.App
	People *handlers.PeopleHandler
}

func (r *PeopleRoutes) SetupPeopleRoutes() {
	people := r.App.Group("/api/v1/people")

	people.Get("/", r.People.GetPeople)
	people.Get("/:id", r.People.GetPerson)

	people.Post("/", r.People.InsertPeople)
	people.Post("/individual", r.People.InsertPerson)

	people.Put("/", r.People.UpdatePerson)

	people.Delete("/:id", r.People.DeletePerson)
}
