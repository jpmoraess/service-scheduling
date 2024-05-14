package entity

import "time"

type Customer struct {
	id              string    //`bson:"_id,omitempty" json:"id,omitempty"`
	establishmentID string    //`bson:"establishmentID" json:"establishmentID"`
	name            string    //`bson:"name" json:"name"`
	phoneNumber     string    //`bson:"phoneNumber" json:"phoneNumber"`
	email           string    //`bson:"email" json:"email"`
	createdAt       time.Time //`bson:"createdAt" json:"createdAt"`
}

func NewCustomer(establishmentID, name, phoneNumber, email string) (*Customer, error) {
	return &Customer{
		establishmentID: establishmentID,
		name:            name,
		phoneNumber:     phoneNumber,
		email:           email,
		createdAt:       time.Now(),
	}, nil
}

func (a *Customer) SetID(id string) {
	a.id = id
}

func (a *Customer) GetID() string {
	return a.id
}
