package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/omerfruk/LibraryApp/database"
	"github.com/omerfruk/LibraryApp/models"
	"github.com/pkg/errors"
)

func KitapGetAll(c *fiber.Ctx) error {
	kitaplar := make([]models.Kitap, 0)
	err := database.DB().Model(&models.Kitap{}).Find(&kitaplar).Error
	if err != nil {
		return errors.New("Database hatasi (KitapGetAll)")
	}
	return c.JSON(kitaplar)
}

func KitapGetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("database hatasi (Kitap Get By id )")
	}
	var kitap models.Kitap
	err = database.DB().Model(&models.Kitap{}).Where("id = ?", id).Find(&kitap).Error
	if err != nil {
		return errors.New("Database hatasi (Kitap Get By id )")
	}
	return c.JSON(kitap)
}

func KitapAdd(c *fiber.Ctx) error {
	var kitap models.Kitap
	err := c.BodyParser(&kitap)
	if err != nil {
		return errors.New("body parser hatasi (kitap Add)")
	}
	err = database.DB().Model(&models.Kitap{}).Create(&kitap).Error
	if err != nil {
		return errors.New("database error (kitap Add)")
	}
	return c.JSON(kitap)
}
