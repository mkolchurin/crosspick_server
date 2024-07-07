package main

import (
	"log"

	"github.com/mkolchurin/crosspick_server/appconfig"
	"github.com/mkolchurin/crosspick_server/database"
	"github.com/mkolchurin/crosspick_server/router"

	"github.com/gin-gonic/gin"
)

const (
	configPath = "config.yaml"
	ginPort    = ":9988"
)

func main() {
	cfg, err := appconfig.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load app config with error '%s'", err)
	}
	err = database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database with error '%s'", err)
	}

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	router.InitWebSocketsDecider(r)

	router.InitMaps(r)
	router.InitS3(r, cfg)
	router.InitSpa(r)

	err = r.Run(ginPort)
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
