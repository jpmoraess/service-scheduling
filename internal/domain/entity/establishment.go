package entity

import (
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Establishment struct {
	id        string    //`bson:"_id,omitempty" json:"id,omitempty"`
	accountID string    //`bson:"accountID" json:"accountID"`
	name      string    //`bson:"name" json:"name"`
	slug      *vo.Slug  //`bson:"slug" json:"slug"`
	createdAt time.Time //`bson:"createdAt" json:"createdAt"`
}

func NewEstablishment(accountID, name, slug string) (*Establishment, error) {
	slugValue, err := vo.NewSlug(slug)
	if err != nil {
		return nil, err
	}
	return &Establishment{
		accountID: accountID,
		name:      name,
		slug:      slugValue,
		createdAt: time.Now(),
	}, nil
}

func (a *Establishment) SetID(id string) {
	a.id = id
}

func (a *Establishment) ID() string {
	return a.id
}
