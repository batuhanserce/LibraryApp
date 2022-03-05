package models

import "gorm.io/gorm"

type Katagori struct {
	gorm.Model
	Cinsi string `json:"cinsi"`
}
