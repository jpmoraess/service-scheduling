package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
)

type CustomerHandler struct {
	createCustomer *usecase.CreateCustomer
	findCustomer   *usecase.FindCustomer
}

func NewCustomerHandler(
	createCustomer *usecase.CreateCustomer, findCustomer *usecase.FindCustomer) *CustomerHandler {
	return &CustomerHandler{
		createCustomer: createCustomer,
		findCustomer:   findCustomer,
	}
}

func (h *CustomerHandler) HandleCreateCustomer(c *fiber.Ctx) error {
	var input dto.CreateCustomerInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.createCustomer.Execute(c.Context(), input); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON("customer created successfully")
}

func (h *CustomerHandler) HandleFindCustomer(c *fiber.Ctx) error {
	page := int64(c.QueryInt("page"))
	size := int64(c.QueryInt("size"))
	output, err := h.findCustomer.Execute(c.Context(), page, size)
	if err != nil {
		return err
	}
	return c.JSON(output)
}
