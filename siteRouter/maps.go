package siteRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/mkolchurin/crosspick_server/db"
)

func AddMaps(r *gin.Engine) {
	r.GET("/api/v1/maps", func(ctx *gin.Context) {
		maps, err := db.GetMaps(0, -1)
		if err != nil {
			processError(ctx, err)
			return
		}
		ctx.JSON(200, maps)
	})

}
