package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/application/usecase"
)

type PasswordResetHandler struct {
	resetPasswordUseCase        *usecase.ResetPasswordUseCase
	requestPasswordResetUseCase *usecase.RequestPasswordResetUseCase
}

func NewPasswordResetHandler(
	resetPasswordUseCase *usecase.ResetPasswordUseCase,
	requestPasswordResetUseCase *usecase.RequestPasswordResetUseCase,
) *PasswordResetHandler {
	return &PasswordResetHandler{
		resetPasswordUseCase:        resetPasswordUseCase,
		requestPasswordResetUseCase: requestPasswordResetUseCase,
	}
}

func (h *PasswordResetHandler) HandleRequestPasswordReset(c *fiber.Ctx) error {
	var input usecase.RequestPasswordResetInputDTO
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.requestPasswordResetUseCase.Execute(c.Context(), input); err != nil {
		return err
	}
	return c.JSON("request password reset successfully")
}

func (h *PasswordResetHandler) HandleResetPassword(c *fiber.Ctx) error {
	var input usecase.ResetPasswordInputDTO
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.resetPasswordUseCase.Execute(c.Context(), c.Query("token"), input); err != nil {
		return err
	}
	return c.JSON("password reset successfully")
}
