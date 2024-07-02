package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/application/usecase"
)

type ProfessionalHandler struct {
	getProfessionalUseCase    *usecase.GetProfessionalUseCase
	findProfessionalUseCase   *usecase.FindProfessionalUseCase
	createProfessionalUseCase *usecase.CreateProfessionalUseCase
}

func NewProfessionalHandler(
	getProfessionalUseCase *usecase.GetProfessionalUseCase,
	findProfessionalUseCase *usecase.FindProfessionalUseCase,
	createProfessionalUseCase *usecase.CreateProfessionalUseCase,
) *ProfessionalHandler {
	return &ProfessionalHandler{
		getProfessionalUseCase:    getProfessionalUseCase,
		findProfessionalUseCase:   findProfessionalUseCase,
		createProfessionalUseCase: createProfessionalUseCase,
	}
}

func (h *ProfessionalHandler) HandleCreateProfessional(c *fiber.Ctx) error {
	var input usecase.CreateProfessionalInputDTO
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	output, err := h.createProfessionalUseCase.Execute(c.Context(), input)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(output)
}

func (h *ProfessionalHandler) HandleGetProfessional(c *fiber.Ctx) error {
	output, err := h.getProfessionalUseCase.Execute(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}
	return c.JSON(output)
}

func (h *ProfessionalHandler) HandleFindProfessional(c *fiber.Ctx) error {
	page := int64(c.QueryInt("page"))
	size := int64(c.QueryInt("size"))
	output, err := h.findProfessionalUseCase.Execute(c.Context(), page, size)
	if err != nil {
		return err
	}
	return c.JSON(output)
}
