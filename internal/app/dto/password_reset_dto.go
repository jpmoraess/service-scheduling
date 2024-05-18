package dto

type RequestPasswordResetInput struct {
	Email string `json:"email"`
}

type ResetPasswordInput struct {
	NewPassword string `json:"newPassword"`
}
