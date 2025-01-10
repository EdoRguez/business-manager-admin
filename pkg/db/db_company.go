package db

import (
	"log"

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
