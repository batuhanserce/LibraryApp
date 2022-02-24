package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/omerfruk/LibraryApp/database"
	"github.com/omerfruk/LibraryApp/models"
	"github.com/pkg/errors"
)

func FavoriGetAll(c *fiber.Ctx) error {
	kullaniciID, err := c.ParamsInt("kullanici_id")
	if err != nil {
		return errors.New("Pars int hatasi (Favori get all)")
	}
	favoriler := make([]models.Favori, 0)
	err = database.DB().Model(&models.Favori{}).Where("kullanici_id = ? ", kullaniciID).Preload("Kitap").Find(&favoriler).Error
	if err != nil {
		return errors.New("database error (favori get all)")
	}
	return c.JSON(favoriler)
}

func FavoriAdd(c *fiber.Ctx) error {
	var favori models.Favori
	err := c.BodyParser(&favori)
	if err != nil {
		return errors.New("body parse hatasi (favori Add)")
	}
	err = database.DB().Model(&models.Favori{}).Create(&favori).Error
	if err != nil {
		return errors.New("database error (favori Add)")
	}
	return c.JSON(favori)
}
