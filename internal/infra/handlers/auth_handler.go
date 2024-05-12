package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/internal/app/dto"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
)

type AuthHandler struct {
	signup *usecase.Signup
	signin *usecase.Signin
}

func NewAuthHandler(signup *usecase.Signup, signin *usecase.Signin) *AuthHandler {
	return &AuthHandler{
		signup: signup,
		signin: signin,
	}
}

func (h *AuthHandler) HandleSignup(c *fiber.Ctx) error {
	var input dto.SignupInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	if err := h.signup.Execute(c.Context(), input); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON("signup successfully")
}

func (h *AuthHandler) HandleSignin(c *fiber.Ctx) error {
	var input dto.SigninInput
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	output, err := h.signin.Execute(c.Context(), input)
	if err != nil {
		return err
	}
	return c.JSON(output)
}
