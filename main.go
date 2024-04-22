package main

import (
	"log"

	"github.com/mkolchurin/crosspick_server/MapDecider"
	"github.com/mkolchurin/crosspick_server/StaticStorage"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	StaticStorage.RegisterMinio(r)
	StaticStorage.RegisterStatic(r)
	MapDecider.RegisterWebsockets(r)

	err := r.Run(":80")
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
