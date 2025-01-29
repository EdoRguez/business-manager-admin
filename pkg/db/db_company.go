package db

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/EdoRguez/business-manager-admin/pkg/config"
	"github.com/EdoRguez/business-manager-admin/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DBCompany *gorm.DB

func DBCompanyConnection(c config.Config) {
	var error error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	DBCompany, error = gorm.Open(postgres.Open(c.Db_Company), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "company.",
			SingularTable: true,
		},
		Logger: newLogger})
	if error != nil || DBCompany == nil {
		log.Fatal(error)
	}
}

func CreateCompany(model *models.Company) error {
	res := DBCompany.Create(model)
	err := res.Error

	return err
}

func GetCompanyById(id int64) *models.Company {
	var model *models.Company
	res := DBCompany.First(&model, id)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return model
}

func GetCompanyByName(name string) *models.Company {
	var model *models.Company
	trimmedName := strings.TrimSpace(name)
	res := DBCompany.First(&model, "LOWER(name) = LOWER(?)", trimmedName)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return model
}

func GetCompanyByNameURL(url string) *models.Company {
	var model *models.Company
	trimmedUrl := strings.TrimSpace(url)
	res := DBCompany.First(&model, "LOWER(name_format_url) = LOWER(?)", trimmedUrl)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return model
}
