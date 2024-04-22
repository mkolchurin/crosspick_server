package main

import (
	"crosspick/MapDecider"
	"crosspick/StaticStorage"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	StaticStorage.RegisterMinio(r)
	StaticStorage.RegisterStatic(r)
	MapDecider.RegisterWebsockets(r)

	err := r.Run(":9988")
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
