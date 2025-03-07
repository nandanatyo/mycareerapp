package util

import (
	"errors"
	"mycareerapp/internal/domain/entity"
)

func CheckGender(gender string) (entity.Gender, error) {
	switch gender {
	case "male":
		return entity.Male, nil
	case "female":
		return entity.Female, nil
	default:
		return "", errors.New("invalid gender")
	}
}
