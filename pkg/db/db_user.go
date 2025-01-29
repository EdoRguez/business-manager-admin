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

var DBUser *gorm.DB

func DBUserConnection(c config.Config) {
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

	DBUser, error = gorm.Open(postgres.Open(c.Db_Auth), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "auth.",
			SingularTable: true,
		},
		Logger: newLogger})
	if error != nil || DBUser == nil {
		log.Fatal(error)
	}
}

func CreateUser(model *models.User) error {
	res := DBUser.Create(model)
	err := res.Error

	return err
}

func GetUserByEmail(email string) *models.User {
	var model *models.User
	trimmedEmail := strings.TrimSpace(email)
	res := DBUser.First(&model, "LOWER(email) = LOWER(?)", trimmedEmail)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return model
}
