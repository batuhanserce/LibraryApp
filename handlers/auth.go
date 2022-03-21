package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/LibraryApp/database"
	"github.com/omerfruk/LibraryApp/models"
	"github.com/omerfruk/LibraryApp/utils"
	"github.com/omerfruk/LibraryApp/viewmodels"
	"strconv"
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
		return c.JSON(nil)
	}
	if dbModel.Parola != utils.Sha256String(gelenModel.Parola) {
		return c.JSON(nil)
	}
	m := make(map[string]string)
	m["id"] = strconv.Itoa(int(dbModel.ID))
	m["isim"] = dbModel.Isim
	m["soyisim"] = dbModel.Soyisim
	m["eposta"] = dbModel.Eposta

	return c.JSON(m)
}

func SingUp(c *fiber.Ctx) error {
	var gelenModel models.Kullanici
	err := c.BodyParser(&gelenModel)
	if err != nil {
		return c.JSON(nil)
	}
	gelenModel.Parola = utils.Sha256String(gelenModel.Parola)
	err = database.DB().Model(&models.Kullanici{}).Create(&gelenModel).Error
	if err != nil {
		return c.JSON(nil)
	}
	return c.JSON(gelenModel)
}
