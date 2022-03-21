package models

import "gorm.io/gorm"

type Kitap struct {
	gorm.Model
	Adi        string   `json:"adi"`
	Yazar      string   `json:"yazar"`
	Ozet       string   `json:"ozet"`
	KatagoriID int      `json:"katagori_id"`
	Katagori   Katagori `json:"katagori"`
}
