package database

import (
	"log"

	"github.com/mkolchurin/crosspick_server/appconfig"
)

func init() {
	cfg, err := appconfig.GetConfig("../config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	Connect(cfg)
	db.AutoMigrate(&Roles{}, &Users{}, &Deciders{}, &Maps{})
}
