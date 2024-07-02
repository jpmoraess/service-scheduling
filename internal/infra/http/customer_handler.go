package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/application/usecase"
)

type CustomerHandler struct {
	findCustomerUseCase   *usecase.FindCustomerUseCase
	createCustomerUseCase *usecase.CreateCustomerUseCase
}

func NewCustomerHandler(
	findCustomerUseCase *usecase.FindCustomerUseCase,
	createCustomerUseCase *usecase.CreateCustomerUseCase,
) *CustomerHandler {
	return &CustomerHandler{
		findCustomerUseCase:   findCustomerUseCase,
		createCustomerUseCase: createCustomerUseCase,
	}
}

func (h *CustomerHandler) HandleCreateCustomer(c *fiber.Ctx) error {
	var input usecase.CreateCustomerInputDTO
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	output, err := h.createCustomerUseCase.Execute(c.Context(), input)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(output)
}

func (h *CustomerHandler) HandleFindCustomer(c *fiber.Ctx) error {
	page := int64(c.QueryInt("page"))
	size := int64(c.QueryInt("size"))
	output, err := h.findCustomerUseCase.Execute(c.Context(), page, size)
	if err != nil {
		return err
	}
	return c.JSON(output)
}
