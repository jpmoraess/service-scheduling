package entity

type Customer struct {
	ID              string `bson:"_id,omitempty" json:"id,omitempty"`
	EstablishmentID string `bson:"establishmentID" json:"establishmentID"`
	Name            string `bson:"name" json:"name"`
	PhoneNumber     string `bson:"phoneNumber" json:"phoneNumber"`
	Email           string `bson:"email" json:"email"`
}

func NewCustomer(establishmentID, name, phoneNumber, email string) (*Customer, error) {
	return &Customer{
		EstablishmentID: establishmentID,
		Name:            name,
		PhoneNumber:     phoneNumber,
		Email:           email,
	}, nil
}
