package db

import (
	"log"

	"github.com/EdoRguez/business-manager-admin/pkg/config"
	"github.com/EdoRguez/business-manager-admin/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DBUser *gorm.DB

func DBUserConnection(c config.Config) {
	var error error
	DBUser, error = gorm.Open(postgres.Open(c.Db_Auth), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "user.",
			SingularTable: true,
		}})
	if error != nil || DBUser == nil {
		log.Fatal(error)
	}
}

func CreateUser(model *models.User) error {
	res := DBUser.Create(model)
	err := res.Error

	return err
}
