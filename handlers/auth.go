package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/LibraryApp/database"
	"github.com/omerfruk/LibraryApp/models"
	"github.com/omerfruk/LibraryApp/utils"
	"github.com/omerfruk/LibraryApp/viewmodels"
	"github.com/pkg/errors"
)

func Login(c *fiber.Ctx) error {
	var gelenModel viewmodels.LoginVM
	err := c.BodyParser(&gelenModel)
	if err != nil {
		return err
	}
	var dbModel models.Kullanici
	err = database.DB().Where("eposta = ?", gelenModel.Eposta).Find(&dbModel).Error
	if err != nil || dbModel.Eposta == "" {
		return errors.New("Şifre veya eposta hatali")
	}
	if dbModel.Parola != utils.Sha256String(gelenModel.Parola) {
		return errors.New("Şifre veya eposta hatali")
	}
	m := make(map[string]string)
	m["isim"] = dbModel.Isim
	m["soyisim"] = dbModel.Soyisim
	m["id"] = dbModel.Eposta

	return c.JSON(m)
}

func SingUp(c *fiber.Ctx) error {
	var gelenModel models.Kullanici
	err := c.BodyParser(&gelenModel)
	if err != nil {
		return errors.New("Body pars edilirken hata oluştu (SingUp)")
	}
	err = database.DB().Model(&models.Kullanici{}).Create(&gelenModel).Error
	if err != nil {
		return errors.New("database Error olusturudu (singup)")
	}
	return c.JSON(gelenModel)
}
