package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/LibraryApp/database"
	"github.com/omerfruk/LibraryApp/models"
	"github.com/pkg/errors"
)

func KatagoriGetAll(c *fiber.Ctx) error {
	katagori := make([]models.Katagori, 0)
	err := database.DB().Model(&models.Katagori{}).Find(&katagori).Error
	if err != nil {
		return errors.New("Database hatasi (katagori get all )")
	}
	return c.JSON(katagori)
}

func KatagoriGetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("database hatasi (katagori Get By id )")
	}
	var katagori models.Katagori
	err = database.DB().Model(&models.Katagori{}).Where("id = ?", id).Find(&katagori).Error
	if err != nil {
		return errors.New("Database hatasi (katagori Get By id )")
	}
	return c.JSON(katagori)
}

func KatagoriAdd(c *fiber.Ctx) error {
	var katagori models.Katagori
	err := c.BodyParser(&katagori)
	if err != nil {
		return errors.New("body parser hatasi (katagori Add)")
	}
	err = database.DB().Model(&models.Katagori{}).Create(&katagori).Error
	if err != nil {
		return errors.New("database error (katagori Add)")
	}
	return c.JSON(katagori)
}

func KatagoriDelete(c *fiber.Ctx) error {
	var katagori models.Katagori
	err := c.BodyParser(&katagori)
	if err != nil {
		return errors.New("body parser hatasi (katagori delete)")
	}
	err = database.DB().Model(&models.Katagori{}).Delete(&katagori).Error
	if err != nil {
		return errors.New("database error (katagori Add)")
	}
	return c.JSON(nil)
}
