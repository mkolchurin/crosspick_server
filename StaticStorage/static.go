package StaticStorage

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterStatic(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/index.html")
	})
	r.GET("/:file", func(ctx *gin.Context) {
		path := ctx.Param("file")
		f, err := os.Open("./www/" + path)
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
		defer f.Close()
		fi, err := f.Stat()
		if err != nil {
			ctx.String(500, err.Error())
			return
		}

		nameSplit := strings.Split(fi.Name(), ".")
		if len(nameSplit) < 2 {
			ctx.String(500, "file not found")
			return
		}
		ctx.DataFromReader(200, fi.Size(), ("text/" + nameSplit[1]), f, nil)
	})
}
