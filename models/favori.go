package models

import "gorm.io/gorm"

type Favori struct {
	gorm.Model
	KullaniciID int       `json:"kullanici_id"`
	Kullanici   Kullanici `json:"kullanici"`

	KitapID int   `json:"kitap_id"`
	Kitap   Kitap `json:"kitap"`
}
