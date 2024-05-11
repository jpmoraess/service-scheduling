package dto

// ACCOUNT SIGNIN

type AccountSigninInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountSigninOutput struct {
	AccessToken string `json:"accessToken"`
}
