package main

import (
	"log"

	"github.com/EdoRguez/business-manager-admin/cmd"
	"github.com/EdoRguez/business-manager-admin/pkg/config"
	"github.com/EdoRguez/business-manager-admin/pkg/db"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db.DBCompanyConnection(c)
	db.DBUserConnection(c)

	cmd.Execute()
}
