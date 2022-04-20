package database

import (
	"fmt"
	"github.com/omerfruk/LibraryApp/models"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	HOST     = "db"
	DATABASE = "library"
	USER     = "library"
	PASSWORD = "library"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func ConnectDB() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=5432", HOST, USER, PASSWORD, DATABASE)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return errors.Wrapf(err, "DB'ye bağlanılamadı: "+dsn)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(2)
	return nil
}
func AutoMigrate() {
	db.AutoMigrate(&models.Kullanici{})
	db.AutoMigrate(&models.Katagori{})
	db.AutoMigrate(&models.Kitap{})
	db.AutoMigrate(&models.Favori{})
}
