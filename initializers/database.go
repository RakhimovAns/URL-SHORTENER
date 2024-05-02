package initializers

import (
	"errors"
	"github.com/RakhimovAns/URL-SHORTENER/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type link struct {
	Link      string         `gorm:"type:text"`
	Short     string         `gorm:"type:text"`
	DeletedAt gorm.DeletedAt `gorm:"index;"`
}

func ConnectToDB() {
	var err error
	dsn := "host=postgres user=postgres password=postgres dbname=yandex port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect with database")
	}
}

func CreateTable() {
	err := DB.AutoMigrate(&link{})
	if err != nil {
		log.Fatal("failed to migrate")
	}
}

func IsLinkExists(linkToCheck string) (bool, error) {
	var l link
	result := DB.Where("link = ?", linkToCheck).First(&l)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func AddLink(l model.Link) {
	DB.Create(&model.Link{Link: l.Link, Short: l.Short})
}

func GetAll() ([]model.Link, error) {
	var links []link
	result := DB.Find(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	var modelLinks []model.Link
	for _, l := range links {
		modelLinks = append(modelLinks, model.Link{Link: l.Link, Short: l.Short})
	}
	return modelLinks, nil
}
