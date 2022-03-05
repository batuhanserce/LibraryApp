package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/LibraryApp/database"
	"github.com/omerfruk/LibraryApp/models"
	"github.com/pkg/errors"
)

func KitapGetAll(c *fiber.Ctx) error {
	kitaplar := make([]models.Kitap, 0)
	err := database.DB().Model(&models.Kitap{}).Preload("Katagori").Find(&kitaplar).Error
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
	if kitap.ID == 0 {
		err = database.DB().Model(&models.Kitap{}).Create(&kitap).Error
		if err != nil {
			return errors.New("database error (kitap Add)")
		}
		return c.JSON(kitap)
	}
	var dbGelen models.Kitap
	err = database.DB().Where("id = ?", kitap.ID).Find(&dbGelen).Error
	if err != nil {
		return errors.New("database error (kitap Add)")
	}
	dbGelen.Adi = kitap.Adi
	dbGelen.Yazar = kitap.Yazar
	dbGelen.KatagoriID = kitap.KatagoriID

	err = database.DB().Save(&dbGelen).Error
	if err != nil {
		return errors.New("database error (kitap Add)")
	}
	return c.JSON(dbGelen)
}

func KitapDelete(c *fiber.Ctx) error {
	var kitap models.Kitap
	err := c.BodyParser(&kitap)
	if err != nil {
		return errors.New("body parser hatasi (kitap delete)")
	}
	err = database.DB().Model(&models.Kitap{}).Delete(&kitap).Error
	if err != nil {
		return errors.New("database error (kitap Add)")
	}
	return c.JSON(nil)
}
