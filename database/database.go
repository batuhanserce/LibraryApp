package database

import (
	"fmt"

	"github.com/omerfruk/LibraryApp/models"
	"github.com/omerfruk/LibraryApp/utils"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	HOST     = "pg_db"
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

func DefaultUserCreate() {
	admin := models.Kullanici{
		Isim:    "Admin",
		Soyisim: "Admin",
		Eposta:  "admin@mail.com",
		Parola:  utils.Sha256String("admin123"),
	}
	err := DB().Where("eposta = ?", admin.Eposta).FirstOrCreate(&admin).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	batu := models.Kullanici{
		Isim:    "Batuhan",
		Soyisim: "Serçe",
		Eposta:  "batuhan@mail.com",
		Parola:  utils.Sha256String("batuhan123"),
	}
	err = DB().Where("eposta = ?", batu.Eposta).FirstOrCreate(&batu).Error
	if err != nil {
		fmt.Println(err.Error())
	}

}
