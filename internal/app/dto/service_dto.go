package dto

type CreateServiceInput struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	DurationInMinutes int64   `json:"durationInMinutes"`
}

type ServiceOutput struct {
	ID                string  `json:"ID"`
	EstablishmentID   string  `json:"establishmentID"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	DurationInMinutes int64   `json:"durationInMinutes"`
	Available         bool    `json:"available"`
}
