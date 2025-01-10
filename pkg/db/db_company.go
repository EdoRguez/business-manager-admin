package db

import (
	"log"

	"github.com/EdoRguez/business-manager-admin/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBCompany *gorm.DB

func DBCompanyConnection(c config.Config) {
	var error error
	DBCompany, error = gorm.Open(postgres.Open(c.Db_Company), &gorm.Config{})
	if error != nil || DBCompany == nil {
		log.Fatal(error)
	}
}
