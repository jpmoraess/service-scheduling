package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/application/usecase"
)

type AuthHandler struct {
	signupUseCase *usecase.SignupUseCase
	signinUseCase *usecase.SigninUseCase
}

func NewAuthHandler(signupUseCase *usecase.SignupUseCase, signinUseCase *usecase.SigninUseCase) *AuthHandler {
	return &AuthHandler{
		signupUseCase: signupUseCase,
		signinUseCase: signinUseCase,
	}
}

func (h *AuthHandler) HandleSignup(c *fiber.Ctx) error {
	var input usecase.SignupInputDTO
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.signupUseCase.Execute(c.Context(), input); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON("signup successfully")
}

func (h *AuthHandler) HandleSignin(c *fiber.Ctx) error {
	var input usecase.SigninInputDTO
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	output, err := h.signinUseCase.Execute(c.Context(), input)
	if err != nil {
		return err
	}
	return c.JSON(output)
}
