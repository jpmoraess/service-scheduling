package entity

type Customer struct {
	id              string
	establishmentID string
	name            string
	phoneNumber     string
	email           string
}

func NewCustomer(establishmentID, name, phoneNumber, email string) (*Customer, error) {
	return &Customer{
		establishmentID: establishmentID,
		name:            name,
		phoneNumber:     phoneNumber,
		email:           email,
	}, nil
}

func RestoreCustomer(id string, establishmentID, name, phoneNumber, email string) (*Customer, error) {
	return &Customer{
		id:              id,
		establishmentID: establishmentID,
		name:            name,
		phoneNumber:     phoneNumber,
		email:           email,
	}, nil
}

func (a *Customer) SetID(id string) {
	a.id = id
}

func (a *Customer) ID() string {
	return a.id
}

func (a *Customer) EstablishmentID() string {
	return a.establishmentID
}

func (a *Customer) Name() string {
	return a.name
}

func (a *Customer) PhoneNumber() string {
	return a.phoneNumber
}

func (a *Customer) Email() string {
	return a.email
}
