package dto

type CreateSchedulingInput struct {
	ServiceID       string `json:"serviceID"`
	CustomerID      string `json:"customerID"`
	ProfessionalID  string `json:"professionalID"`
	EstablishmentID string `json:"establishmentID"`
	Date            string `json:"date"`
	Time            string `json:"time"`
}
