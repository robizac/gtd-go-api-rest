package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Capitura struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
}

func ValidaDadosDeCapitura(capitura *Capitura) error {
	if err := validator.Validate(capitura); err != nil {
		return err
	}
	return nil
}
