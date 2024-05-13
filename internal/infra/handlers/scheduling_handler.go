package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
)

type SchedulingHandler struct {
	createScheduling *usecase.CreateScheduling
}

func NewSchedulingHandler(createScheduling *usecase.CreateScheduling) *SchedulingHandler {
	return &SchedulingHandler{
		createScheduling: createScheduling,
	}
}

func (h *SchedulingHandler) HandleCreateScheduling(c *fiber.Ctx) error {
	var input dto.CreateSchedulingInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.createScheduling.Execute(c.Context(), input); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON("scheduling created successfully")
}
