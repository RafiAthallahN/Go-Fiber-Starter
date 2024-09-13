package handlers

import (
	"go-fiber/starter/backend/entities"
	"go-fiber/starter/backend/services"
	"go-fiber/starter/backend/utils"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PeopleHandler struct {
	service *services.PeopleService
	log     *logrus.Logger
}

func NewPeopleHandler(service *services.PeopleService, log *logrus.Logger) *PeopleHandler {
	return &PeopleHandler{service, log}
}

func (h *PeopleHandler) GetPeople(c *fiber.Ctx) error {

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	people, peopleCount, errGetPeople := h.service.GetPeople(page, limit)
	if errGetPeople != nil {
		h.log.Errorln("Error ")
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed retrieve Person List",
			Error:      errGetPeople,
		})
	}

	totalPage := int(math.Ceil(float64(peopleCount) / float64(limit)))

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Successfully Get People List",
		Data:       people,
		Meta: &utils.Pagination{
			Total:     &peopleCount,
			Page:      &page,
			Limit:     &limit,
			TotalPage: &totalPage,
		},
	})
}

func (h *PeopleHandler) GetPerson(c *fiber.Ctx) error {

	id, errParam := c.ParamsInt("id")
	if errParam != nil {
		h.log.Errorf("Error parse parameter: %v", errParam)
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed retrieve Person Data",
			Error:      "Request not valid",
		})
	}

	person, errGetPerson := h.service.GetPerson(id)
	if errGetPerson != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed retrieve Person Data",
			Error:      errGetPerson,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Succeed retrieve person data",
		Data:       person,
	})
}

func (h *PeopleHandler) InsertPeople(c *fiber.Ctx) error {

	data, errParseAndValidate := utils.ParseAndValidate[[]entities.People](c, c.Request().Body())
	if errParseAndValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errParseAndValidate)
	}

	response, errInsertPeople := h.service.InsertPeople(data)
	if errInsertPeople != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed to Insert People",
			Error:      errInsertPeople,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utils.Response{
		StatusCode: fiber.StatusCreated,
		Message:    "Succeed insert people",
		Data:       response,
	})
}

func (h *PeopleHandler) InsertPerson(c *fiber.Ctx) error {

	data, errParseAndValidate := utils.ParseAndValidate[entities.People](c, c.Request().Body())
	if errParseAndValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errParseAndValidate)
	}

	response, errInsertPerson := h.service.InsertPerson(data)
	if errInsertPerson != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to insert person",
			Error:      errInsertPerson,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utils.Response{
		StatusCode: fiber.StatusCreated,
		Message:    "Succeed insert person",
		Data:       response,
	})
}

func (h *PeopleHandler) UpdatePerson(c *fiber.Ctx) error {

	data, errParseAndValidate := utils.ParseAndValidate[entities.People](c, c.Request().Body())
	if errParseAndValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errParseAndValidate)
	}

	response, errUpdatePerson := h.service.UpdatePerson(data)
	if errUpdatePerson != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to update person",
			Error:      errUpdatePerson,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utils.Response{
		StatusCode: fiber.StatusCreated,
		Message:    "Succeed update person",
		Data:       response,
	})
}

func (h *PeopleHandler) DeletePerson(c *fiber.Ctx) error {

	id, errParam := c.ParamsInt("id")
	if errParam != nil {
		h.log.Errorf("Error parse parameter: %v", errParam)
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed retrieve Person Data",
			Error:      "Request not valid",
		})
	}

	if errDeletePerson := h.service.DeletePerson(id); errDeletePerson != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed delete Person Data",
			Error:      errDeletePerson,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Succeed delete person data",
		Data:       nil,
	})
}
