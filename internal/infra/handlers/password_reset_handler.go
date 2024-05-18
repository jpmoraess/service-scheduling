package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
)

type PasswordResetHandler struct {
	resetPassword        *usecase.ResetPassword
	requestPasswordReset *usecase.RequestPasswordReset
}

func NewPasswordResetHandler(
	resetPassword *usecase.ResetPassword,
	requestPasswordReset *usecase.RequestPasswordReset,
) *PasswordResetHandler {
	return &PasswordResetHandler{
		resetPassword:        resetPassword,
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

func (h *PasswordResetHandler) HandleResetPassword(c *fiber.Ctx) error {
	var input dto.ResetPasswordInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.resetPassword.Execute(c.Context(), c.Query("token"), input); err != nil {
		return err
	}
	return c.JSON("password reset successfully")
}
