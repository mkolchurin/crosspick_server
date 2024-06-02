package main

import (
	"log"

	"github.com/mkolchurin/crosspick_server/database"
	"github.com/mkolchurin/crosspick_server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	router.AddMaps(r)
	router.AddMinio(r)
	router.AddSpa(r)
	router.AddWsDecider(r)

	err := r.Run(":9988")
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
