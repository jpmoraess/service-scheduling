package entity

import (
	"time"

	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
)

type Professional struct {
	ID              string       `bson:"_id,omitempty" json:"id,omitempty"`
	AccountID       string       `bson:"accountID" json:"accountID"`
	EstablishmentID string       `bson:"establishmentID" json:"establishmentID"`
	Name            string       `bson:"name" json:"name"`
	WorkPlan        *vo.WorkPlan `bson:"workPlan" json:"workPlan"`
	Active          bool         `bson:"active" json:"active"`
	CreatedAt       time.Time
}

func NewProfessional(accountID, establishmentID, name string) (*Professional, error) {
	workPlan, err := vo.DefaultWorkPlan()
	if err != nil {
		return nil, err
	}
	return &Professional{
		AccountID:       accountID,
		EstablishmentID: establishmentID,
		Name:            name,
		WorkPlan:        workPlan,
		Active:          true,
		CreatedAt:       time.Now(),
	}, nil
}
