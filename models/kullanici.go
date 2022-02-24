package models

import "gorm.io/gorm"

type Kullanici struct {
	gorm.Model
	Isim      string  `json:"isim"`
	Soyisim   string  `json:"soyisim"`
	Eposta    string  `json:"eposta"`
	Parola    string  `json:"parola"`
	Kitaplari []Kitap `gorm:"many2many:kullanici_kitap;"`
}
