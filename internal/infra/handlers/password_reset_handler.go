package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
)

type PasswordResetHandler struct {
	requestPasswordReset *usecase.RequestPasswordReset
}

func NewPasswordResetHandler(requestPasswordReset *usecase.RequestPasswordReset) *PasswordResetHandler {
	return &PasswordResetHandler{
		requestPasswordReset: requestPasswordReset,
	}
}

func (h *PasswordResetHandler) HandleRequestPasswordReset(c *fiber.Ctx) error {
	var input dto.RequestPasswordResetInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.requestPasswordReset.Execute(c.Context(), input); err != nil {
		return err
	}
	return c.JSON("request password reset successfully")
}
