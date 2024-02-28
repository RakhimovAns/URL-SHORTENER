package initializers

import (
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
	dsn := "host=localhost user=postgres password=postgres dbname=yandex port=5432 sslmode=disable"
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

func IsLinkExists(linkToCheck string) (bool, model.Link, error) {
	var l link
	result := DB.Where("link = ?", linkToCheck).First(&l)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, model.Link{Link: l.Link, Short: l.Short}, nil
		}
		return false, model.Link{Link: l.Link, Short: l.Short}, result.Error
	}
	return true, model.Link{Link: l.Link, Short: l.Short}, nil
}

func AddLink(l model.Link) {
	DB.Create(&model.Link{Link: l.Link, Short: l.Short})
}
func UpdateLink(l model.Link) {
	DB.Model(&link{}).Where("short = ?", l.Short).Updates(link{Link: l.Link, Short: l.Short})
}
