package models

import "gorm.io/gorm"

type Kitap struct {
	gorm.Model
	Adi        string   `json:"adi"`
	Yazar      string   `json:"yazari"`
	KatagoriID int      `json:"katagori_id"`
	Katagori   Katagori `json:"katagori"`
}
