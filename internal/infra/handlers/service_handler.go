package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
)

type ServiceHandler struct {
	createService usecase.CreateService
	listServices  usecase.ListServices
}

func NewServiceHandler(createService usecase.CreateService, listServices usecase.ListServices) *ServiceHandler {
	return &ServiceHandler{
		createService: createService,
		listServices:  listServices,
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
	return c.Status(fiber.StatusCreated).JSON("service created successfully")
}

func (h *ServiceHandler) HandleListServicesByEstablishment(c *fiber.Ctx) error {
	output, err := h.listServices.Execute(c.Context(), c.Query("establishmentID"))
	if err != nil {
		return err
	}
	return c.JSON(output)
}
