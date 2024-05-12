package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
)

type ProfessionalHandler struct {
	createProfessional *usecase.CreateProfessional
}

func NewProfessionalHandler(createProfessional *usecase.CreateProfessional) *ProfessionalHandler {
	return &ProfessionalHandler{
		createProfessional: createProfessional,
	}
}

func (h *ProfessionalHandler) HandleCreateProfessional(c *fiber.Ctx) error {
	var input dto.CreateProfessionalInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.createProfessional.Execute(c.Context(), input); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON("professional created successfully")
}
