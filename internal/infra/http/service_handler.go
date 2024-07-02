package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/application/dto"
	"github.com/jpmoraess/service-scheduling/internal/application/usecase"
)

type ServiceHandler struct {
	findServiceUseCase   *usecase.FindServiceUseCase
	createServiceUseCase *usecase.CreateServiceUseCase
}

func NewServiceHandler(
	findServiceUseCase *usecase.FindServiceUseCase,
	createServiceUseCase *usecase.CreateServiceUseCase,
) *ServiceHandler {
	return &ServiceHandler{
		findServiceUseCase:   findServiceUseCase,
		createServiceUseCase: createServiceUseCase,
	}
}

func (h *ServiceHandler) HandleCreateService(c *fiber.Ctx) error {
	var input dto.CreateServiceInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	output, err := h.createServiceUseCase.Execute(c.Context(), input)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(output)
}

func (h *ServiceHandler) HandleFindServiceByEstablishment(c *fiber.Ctx) error {
	output, err := h.findServiceUseCase.Execute(c.Context(), c.Query("establishmentID"))
	if err != nil {
		return err
	}
	return c.JSON(output)
}
