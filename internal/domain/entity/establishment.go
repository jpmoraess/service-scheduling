package entity

import (
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Establishment struct {
	id        string
	accountID string
	name      string
	slug      *vo.Slug
	createdAt time.Time
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

func RestoreEstablishment(id, accountID, name, slug string, createdAt time.Time) (*Establishment, error) {
	slugValue, err := vo.NewSlug(slug)
	if err != nil {
		return nil, err
	}
	return &Establishment{
		id:        id,
		accountID: accountID,
		name:      name,
		slug:      slugValue,
		createdAt: createdAt,
	}, nil
}

func (e *Establishment) SetID(id string) {
	e.id = id
}

func (e *Establishment) ID() string {
	return e.id
}

func (e *Establishment) AccountID() string {
	return e.accountID
}

func (e *Establishment) Name() string {
	return e.name
}

func (e *Establishment) Slug() string {
	return e.slug.Value()
}

func (e *Establishment) CreatedAt() time.Time {
	return e.createdAt
}
