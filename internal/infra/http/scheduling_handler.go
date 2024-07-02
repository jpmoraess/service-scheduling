package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/application/usecase"
)

type SchedulingHandler struct {
	createSchedulingUseCase *usecase.CreateSchedulingUseCase
}

func NewSchedulingHandler(createSchedulingUseCase *usecase.CreateSchedulingUseCase) *SchedulingHandler {
	return &SchedulingHandler{
		createSchedulingUseCase: createSchedulingUseCase,
	}
}

func (h *SchedulingHandler) HandleCreateScheduling(c *fiber.Ctx) error {
	var input usecase.CreateSchedulingInputDTO
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	output, err := h.createSchedulingUseCase.Execute(c.Context(), input)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(output)
}
