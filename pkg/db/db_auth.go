package db

import (
	"log"

	"github.com/EdoRguez/business-manager-admin/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBAuth *gorm.DB

func DBAuthConnection(c config.Config) {
	var error error
	DBAuth, error = gorm.Open(postgres.Open(c.Db_Auth), &gorm.Config{})
	if error != nil || DBAuth == nil {
		log.Fatal(error)
	}
}
