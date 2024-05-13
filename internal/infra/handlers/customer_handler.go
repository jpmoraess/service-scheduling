package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
)

type CustomerHandler struct {
	createCustomer *usecase.CreateCustomer
}

func NewCustomerHandler(createCustomer *usecase.CreateCustomer) *CustomerHandler {
	return &CustomerHandler{
		createCustomer: createCustomer,
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
