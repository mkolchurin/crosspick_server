package router

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func spa(urlPrefix, spaDirectory string) gin.HandlerFunc {
	directory := static.LocalFile(spaDirectory, true)
	fileserver := http.FileServer(directory)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if directory.Exists(urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		} else {
			c.Request.URL.Path = "/"
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

func InitSpa(r *gin.Engine) {
	r.Use(spa("/", "./www"))
}
