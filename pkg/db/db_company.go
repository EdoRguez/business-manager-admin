package db

import (
	"log"
	"strings"

	"github.com/EdoRguez/business-manager-admin/pkg/config"
	"github.com/EdoRguez/business-manager-admin/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DBCompany *gorm.DB

func DBCompanyConnection(c config.Config) {
	var error error
	DBCompany, error = gorm.Open(postgres.Open(c.Db_Company), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "company.",
			SingularTable: true,
		}})
	if error != nil || DBCompany == nil {
		log.Fatal(error)
	}
}

func CreateCompany(model *models.Company) error {
	res := DBCompany.Create(model)
	err := res.Error

	return err
}

func GetCompanyByName(name string) *models.Company {
	var model *models.Company
	trimmedName := strings.TrimSpace(name)
	DBCompany.First(&model, "LOWER(name) = LOWER(?)", trimmedName)
	return model
}

func GetCompanyByNameURL(url string) *models.Company {
	var model *models.Company
	trimmedUrl := strings.TrimSpace(url)
	DBCompany.First(&model, "LOWER(url) = LOWER(?)", trimmedUrl)
	return model
}
