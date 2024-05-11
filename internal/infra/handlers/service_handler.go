package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
)

type ServiceHandler struct {
	createService usecase.CreateService
	findServices  usecase.FindServices
}

func NewServiceHandler(createService usecase.CreateService, findServices usecase.FindServices) *ServiceHandler {
	return &ServiceHandler{
		createService: createService,
		findServices:  findServices,
	}
}

func (h *ServiceHandler) HandleCreateService(c *fiber.Ctx) error {
	var input dto.CreateServiceInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.createService.Execute(c.Context(), input); err != nil {
		return err
	}
	return c.Status(201).JSON("service created successfully")
}

func (h *ServiceHandler) HandleFindServicesByEstablishment(c *fiber.Ctx) error {
	output, err := h.findServices.Execute(c.Context(), c.Query("establishmentID"))
	if err != nil {
		return err
	}
	return c.JSON(output)
}
