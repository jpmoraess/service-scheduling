package dto

type SigninInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninOutput struct {
	AccessToken string `json:"accessToken"`
}
