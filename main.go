package main

import (
	"log"

	"github.com/mkolchurin/crosspick_server/db"
	"github.com/mkolchurin/crosspick_server/siteRouter"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	siteRouter.AddMaps(r)
	siteRouter.AddMinio(r)
	siteRouter.AddSpa(r)
	siteRouter.AddWsDecider(r)

	err := r.Run(":9988")
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
