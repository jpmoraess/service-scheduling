package dto

type CreateProfessionalInput struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type ProfessionalOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	//Email       string `json:"email"`
	//PhoneNumber string `json:"phoneNumber"`
	//Password    string `json:"password"`
}
