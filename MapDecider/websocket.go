package MapDecider

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func RegisterWebsockets(r *gin.Engine) {
	var client *websocket.Conn
	var err error

	r.GET("/subscribe", func(ctx *gin.Context) {
		client, err = websocket.Accept(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer client.CloseNow()
		for {
			_, data, err := client.Read(ctx)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(string(data))
			client.Write(ctx, websocket.MessageText, data)
		}
	})
	r.POST("/publish", func(ctx *gin.Context) {
		if client == nil {
			ctx.String(500, "no client")
			return
		}
		data, err := ctx.GetRawData()
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
		client.Write(ctx, websocket.MessageText, data)
		ctx.String(202, "accepted")
	})
}
