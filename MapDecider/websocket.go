package MapDecider

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

type mapDescriptor struct {
	Uid     string `json:"uid"`
	MapName string `json:"mapName"`
	Checked bool   `json:"checked"`
}

type deciderUser struct {
	Con      *websocket.Conn
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type registerRequest struct {
	Username string `json:"username"`
	Uid      string `json:"uid"`
}

type regiserResponse struct {
	Role string `json:"role"`
}

type Map struct {
	Name    string `json:"name"`
	Img     string `json:"img"`
	Checked bool   `json:"checked"`
}

var mapsStorage = [][]Map{{
	{"Blood river", "2p_blood_river_[Rem].jpg", false},
	{"Meeting of minds", "2p_meeting_of_minds.jpg", false},
	{"Quest triumph", "2p_quests_triumph.jpg", false},
	{"Outer reaches", "2p_outer_reaches.jpg", false},
	{"Shrine of excellion", "2p_shrine_of_excellion_[Rem].jpg", false},
	{"Deadly fun archeology", "2p_deadly_fun_archeology.jpg", false},
	{"Fallen city", "2p_fallen_city_[Rem].jpg", false}}, {},
}

type deciderOp struct {
	Op   string      `json:"op"`
	Body interface{} `json:"body"`
}

func RegisterWebsockets(r *gin.Engine) {

	clients := []deciderUser{}

	r.GET("/api/v1/decider/*mapId", func(ctx *gin.Context) {
		mapIdParam := ctx.Param("mapId")

		mapId, err := strconv.Atoi(strings.Split(mapIdParam, "/")[1])
		if err != nil || mapId < 0 || mapId > len(mapsStorage) {
			ctx.String(500, err.Error())
			return
		}
		c, err := websocket.Accept(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer c.CloseNow()
		for {
			_, data, err := c.Read(ctx)
			if err != nil {
				log.Println(err)
				return
			}
			op := deciderOp{}
			err = json.Unmarshal(data, &op)
			if err != nil {
				log.Println(err)
				return
			}
			switch op.Op {
			case "register":
				req := registerRequest{}
				body, err := json.Marshal(op.Body)
				if err != nil {
					log.Println(err)
					return
				}
				err = json.Unmarshal(body, &req)
				if err != nil {
					log.Println(err)
					return
				}
				user := deciderUser{
					Con:      c,
					Uid:      req.Uid,
					Username: req.Username,
				}
				clients = append(clients, user)

				resp := regiserResponse{}
				switch req.Uid {
				case "1":
					resp.Role = "user"
				case "2":
					resp.Role = "user"
				default:
					resp.Role = "observer"
				}

				respData, err := json.Marshal(resp)
				if err != nil {
					log.Println(err)
					return
				}
				c.Write(ctx, websocket.MessageText, respData)
			case "clear":
				for i := range mapsStorage[mapId] {
					mapsStorage[mapId][i].Checked = false
				}
			case "getMaps":
				data, err := json.Marshal(mapsStorage[mapId])
				if err != nil {
					log.Println(err)
					return
				}
				c.Write(ctx, websocket.MessageText, data)
			case "selectMap":
				body, err := json.Marshal(op.Body)
				if err != nil {
					log.Println(err)
					return
				}
				mapDesc := mapDescriptor{}
				err = json.Unmarshal(body, &mapDesc)
				if err != nil {
					log.Println(err)
					return
				}
				for i := range clients {
					if mapDesc.Uid == clients[i].Uid {
						if clients[i].Role == "observer" {
							continue
						}
						for Rec := range mapsStorage[mapId] {
							if mapsStorage[mapId][Rec].Name == mapDesc.MapName {
								mapsStorage[mapId][Rec].Checked = mapDesc.Checked
							}
						}
					}
				}
			}
		}
	})

	// raw, err := ctx.GetRawData()
	// if err != nil {
	// 	ctx.String(500, err.Error())
	// 	return
	// }
	// user := deciderUser{}
	// err = json.Unmarshal(raw, &user)
	// if err != nil {
	// 	ctx.String(500, err.Error())
	// 	return
	// }
	// clients = append(clients, user)
	// ctx.JSON(200, mapsStorage)

	// var client []*websocket.Conn
	// r.GET("/api/v1/decider/subscribe", func(ctx *gin.Context) {
	// 	c, err := websocket.Accept(ctx.Writer, ctx.Request, nil)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	defer c.CloseNow()
	// 	client = append(client, c)
	// 	for {
	// 		_, data, err := c.Read(ctx)
	// 		if err != nil {
	// 			log.Println(err)
	// 			return
	// 		}
	// 		fmt.Println(string(data))
	// 		c.Write(ctx, websocket.MessageText, data)
	// 	}
	// })
	// r.POST("/publish", func(ctx *gin.Context) {
	// 	if client == nil {
	// 		ctx.String(500, "no client")
	// 		return
	// 	}
	// 	data, err := ctx.GetRawData()
	// 	if err != nil {
	// 		ctx.String(500, err.Error())
	// 		return
	// 	}
	// 	for i := range client {
	// 		client[i].Write(ctx, websocket.MessageText, data)
	// 	}
	// 	ctx.String(202, "accepted")
	// })
}
