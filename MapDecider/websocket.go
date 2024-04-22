package MapDecider

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func RegisterWebsockets(r *gin.Engine) {
	var client []*websocket.Conn
	r.GET("/subscribe", func(ctx *gin.Context) {
		c, err := websocket.Accept(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer c.CloseNow()
		client = append(client, c)
		for {
			_, data, err := c.Read(ctx)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(string(data))
			c.Write(ctx, websocket.MessageText, data)
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
		for i := range client {
			client[i].Write(ctx, websocket.MessageText, data)
		}
		ctx.String(202, "accepted")
	})
}
