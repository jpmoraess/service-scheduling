package entity

import "time"

type Customer struct {
	id              string
	establishmentID string
	name            string
	phoneNumber     string
	email           string
	createdAt       time.Time
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

func RestoreCustomer(id string, establishmentID, name, phoneNumber, email string, createdAt time.Time) (*Customer, error) {
	return &Customer{
		id:              id,
		establishmentID: establishmentID,
		name:            name,
		phoneNumber:     phoneNumber,
		email:           email,
		createdAt:       createdAt,
	}, nil
}

func (c *Customer) SetID(id string) {
	c.id = id
}

func (c *Customer) ID() string {
	return c.id
}

func (c *Customer) EstablishmentID() string {
	return c.establishmentID
}

func (c *Customer) Name() string {
	return c.name
}

func (c *Customer) PhoneNumber() string {
	return c.phoneNumber
}

func (c *Customer) Email() string {
	return c.email
}

func (c *Customer) CreatedAt() time.Time {
	return c.createdAt
}
