package vo

import (
	"errors"
)

type AccountType int

const (
	OwnerType AccountType = iota + 1
	ProfessionalType
)

func (at AccountType) String() string {
	switch at {
	case OwnerType:
		return "Owner"
	case ProfessionalType:
		return "Professional"
	default:
		return "Unknown"
	}
}

func ParseAccountTypeFromString(s string) (AccountType, error) {
	switch s {
	case "OwnerType":
		return OwnerType, nil
	case "ProfessionalType":
		return ProfessionalType, nil
	default:
		return -1, errors.New("invalid AccountType")
	}
}

func ParseAccountTypeFromInt(s int) (AccountType, error) {
	switch s {
	case 1:
		return OwnerType, nil
	case 2:
		return ProfessionalType, nil
	default:
		return -1, errors.New("invalid AccountType")
	}
}
