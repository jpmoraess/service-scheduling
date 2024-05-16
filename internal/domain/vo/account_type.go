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

func (at AccountType) Int() int {
	switch at {
	case OwnerType:
		return 1
	case ProfessionalType:
		return 2
	default:
		return -1
	}
}

func AccountTypeFromString(s string) (AccountType, error) {
	switch s {
	case "OwnerType":
		return OwnerType, nil
	case "ProfessionalType":
		return ProfessionalType, nil
	default:
		return -1, errors.New("invalid AccountType")
	}
}

func AccountTypeFromInt(i int) (AccountType, error) {
	switch i {
	case 1:
		return OwnerType, nil
	case 2:
		return ProfessionalType, nil
	default:
		return -1, errors.New("invalid AccountType")
	}
}
