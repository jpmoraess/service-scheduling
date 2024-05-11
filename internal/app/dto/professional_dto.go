package dto

type CreateProfessionalInput struct {
	EstablishmentID string `json:"establishmentID"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phoneNumber"`
	Password        string `json:"password"`
}
