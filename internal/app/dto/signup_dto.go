package dto

type SignupInput struct {
	Name              string `json:"name"`
	EstablishmentName string `json:"establishmentName"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	Password          string `json:"password"`
}
