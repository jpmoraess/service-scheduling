package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
)

type ProfessionalHandler struct {
	createProfessional *usecase.CreateProfessional
	getProfessional    *usecase.GetProfessional
	findProfessional   *usecase.FindProfessional
}

func NewProfessionalHandler(
	createProfessional *usecase.CreateProfessional,
	getProfessional *usecase.GetProfessional,
	findProfessional *usecase.FindProfessional) *ProfessionalHandler {
	return &ProfessionalHandler{
		createProfessional: createProfessional,
		getProfessional:    getProfessional,
		findProfessional:   findProfessional,
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

func (h *ProfessionalHandler) HandleGetProfessional(c *fiber.Ctx) error {
	output, err := h.getProfessional.Execute(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}
	return c.JSON(output)
}

func (h *ProfessionalHandler) HandleFindProfessional(c *fiber.Ctx) error {
	page := int64(c.QueryInt("page"))
	size := int64(c.QueryInt("size"))
	output, err := h.findProfessional.Execute(c.Context(), page, size)
	if err != nil {
		return err
	}
	return c.JSON(output)
}
