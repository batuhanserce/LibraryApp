package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/LibraryApp/database"
	"github.com/omerfruk/LibraryApp/models"
	"github.com/omerfruk/LibraryApp/utils"
	"github.com/pkg/errors"
)

func KullaniciGetAll(c *fiber.Ctx) error {
	kullanicilar := make([]models.Kullanici, 0)
	err := database.DB().Model(&models.Kullanici{}).Preload("Kitaplari").Find(&kullanicilar).Error
	if err != nil {
		return errors.New("Database hatasi (kullaniciGetAll)")
	}
	for i := range kullanicilar {
		kullanicilar[i].Parola = ""
	}
	return c.JSON(kullanicilar)
}

func KullaniciGetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("database hatasi (kullanici Get By id )")
	}
	var kullanici models.Kullanici
	err = database.DB().Model(&models.Kullanici{}).Where("id = ?", id).Find(&kullanici).Error
	if err != nil {
		return errors.New("Database hatasi (kullanici Get By id )")
	}
	kullanici.Parola = ""
	return c.JSON(kullanici)
}

func KullaniciAdd(c *fiber.Ctx) error {
	var kullanici models.Kullanici
	err := c.BodyParser(&kullanici)
	if err != nil {
		fmt.Println(err)
		return errors.New("body parser hatasi (kullanici Add)")
	}
	if kullanici.ID == 0 {
		err = database.DB().Model(&models.Kullanici{}).Create(&kullanici).Error
		if err != nil {
			return errors.New("database error (kullanici Add)")
		}
		return c.JSON(kullanici)
	}
	var dbGelen models.Kullanici
	err = database.DB().Where("id = ?", kullanici.ID).Find(&dbGelen).Error
	if err != nil {
		return errors.New("database error (kullanici Add)")
	}

	dbGelen.Isim = kullanici.Isim
	dbGelen.Soyisim = kullanici.Soyisim
	dbGelen.Eposta = kullanici.Eposta
	if kullanici.Parola != "" {
		dbGelen.Parola = utils.Sha256String(kullanici.Parola)
	}
	dbGelen.Kitaplari = kullanici.Kitaplari

	err = database.DB().Save(&dbGelen).Error
	if err != nil {
		return errors.New("database error (kullanici Add)")
	}
	return c.JSON(dbGelen)
}

func KullaniciDelete(c *fiber.Ctx) error {
	var kullanici models.Kullanici
	err := c.BodyParser(&kullanici)
	if err != nil {
		return errors.New("body parser hatasi (kullanici delete)")
	}
	err = database.DB().Model(&models.Kullanici{}).Delete(&kullanici).Error
	if err != nil {
		return errors.New("database error (kullanici Add)")
	}
	return c.JSON(nil)
}
