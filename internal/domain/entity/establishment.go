package entity

import (
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Establishment struct {
	id        string
	accountID string
	name      string
	slug      *vo.Slug
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
	}, nil
}

func RestoreEstablishment(id, accountID, name, slug string) (*Establishment, error) {
	slugValue, err := vo.NewSlug(slug)
	if err != nil {
		return nil, err
	}
	return &Establishment{
		id:        id,
		accountID: accountID,
		name:      name,
		slug:      slugValue,
	}, nil
}

func (a *Establishment) SetID(id string) {
	a.id = id
}

func (a *Establishment) ID() string {
	return a.id
}

func (a *Establishment) AccountID() string {
	return a.accountID
}

func (a *Establishment) Name() string {
	return a.name
}

func (a *Establishment) Slug() string {
	return a.slug.Value()
}
