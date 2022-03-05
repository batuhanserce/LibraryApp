package models

import "gorm.io/gorm"

type Kitap struct {
	gorm.Model
	Adi        string   `json:"adi"`
	Yazar      string   `json:"yazar"`
	KatagoriID int      `json:"katagori_id"`
	Katagori   Katagori `json:"katagori"`
}
