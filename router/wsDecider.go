package router

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mkolchurin/crosspick_server/decider"
	"nhooyr.io/websocket"
)

/*
{
	"type": "create"/"join"/"update"/"post",
	"user": {
		"uid": "userUniqID",
		"username": "user1"
	},
	"body": { // depends on "type"
		..any..

	}
}
*/

const (
	RoleUser     = "user"
	RoleObserver = "observer"
	RoleAdmin    = "admin"
	/* create decider */
	RequestTypeCreate = "create"
	/* join on existed decider */
	RequestTypeJoin = "join"
	/* get info  */
	RequestTypeUpdate = "update"
	RequestTypeLeave  = "leave"
	RequestTypePost   = "post"
)

type UserDescriptor struct {
	UID  string `json:"uid"`
	Name string `json:"username"`
}

/* message from/to websocket */
type WsMessage struct {
	Type string         `json:"type"`
	User UserDescriptor `json:"user"`
	Body interface{}    `json:"body"`
}

/* message type: create */
// type WsBodyCreate struct {
// 	DeciderUid string `json:"deciderUid"`
// }

/* message type: join */
type WsBodyJoin struct {
	PartyUid int    `json:"partyUid"`
	Role     string `json:"role"`
}

/* message type: leave */
type WsBodyLeave struct {
	PartyUid int `json:"partyUid"`
}

/* message type: update */
type WsBodyUpdate struct {
	PartyUid int `json:"partyUid"`
}

/* message type: post */
type WsBodyPost struct {
	PartyUid int `json:"partyUid"`
	/* 0 - maps; 1 - races; 2 - bans */
	Stage   int    `json:"stage"`
	MapName string `json:"mapName"`
}

type WsResponse struct {
	PartyUid int    `json:"partyUid"`
	Result   string `json:"result"`
}

func (req WsMessage) processRequest(deciderUid string) (*WsMessage, error) {
	reqBody, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	switch req.Type {
	case RequestTypeCreate:
		partyUid, err := decider.CreateParty(deciderUid, req.User.UID)
		return &WsMessage{Type: RequestTypeUpdate, User: req.User, Body: WsResponse{PartyUid: partyUid, Result: err.Error()}}, err
	case RequestTypeJoin:
		bodyJoin := WsBodyJoin{}
		err = json.Unmarshal(reqBody, &bodyJoin)
		if err != nil {
			return nil, err
		}
		err := decider.JoinParty(deciderUid, bodyJoin.PartyUid, req.User.UID, bodyJoin.Role)
		return &WsMessage{Type: RequestTypeUpdate, User: req.User, Body: WsResponse{PartyUid: bodyJoin.PartyUid, Result: err.Error()}}, err
	case RequestTypeLeave:
		bodyLeave := WsBodyLeave{}
		err = json.Unmarshal(reqBody, &bodyLeave)
		if err != nil {
			return nil, err
		}
		err := decider.LeaveParty(deciderUid, bodyLeave.PartyUid, req.User.UID)
		return &WsMessage{Type: RequestTypeUpdate, User: req.User, Body: WsResponse{PartyUid: bodyLeave.PartyUid, Result: err.Error()}}, err
	case RequestTypeUpdate:
		bodyUpdate := WsBodyUpdate{}
		err = json.Unmarshal(reqBody, &bodyUpdate)
		if err != nil {
			return nil, err
		}
		party, err := decider.GetParty(deciderUid, bodyUpdate.PartyUid)
		return &WsMessage{Type: RequestTypeUpdate, User: req.User, Body: party}, err
	case RequestTypePost:
		bodyPost := WsBodyPost{}
		err = json.Unmarshal(reqBody, &bodyPost)
		if err != nil {
			return nil, err
		}
		return nil, errors.ErrUnsupported
	default:
		log.Println("unsupported message")
	}
	return nil, errors.ErrUnsupported
}

func processError(ctx *gin.Context, err error) {
	log.Println(err)
	ctx.String(500, err.Error())
}

func AddWsDecider(r *gin.Engine) {
	r.GET("/api/v1/decider/*deciderUID", func(ctx *gin.Context) {
		deciderId := ctx.Param("deciderUID")
		wsConnection, err := websocket.Accept(ctx.Writer, ctx.Request, nil)
		if err != nil {
			processError(ctx, err)
			return
		}
		defer wsConnection.CloseNow()
		for {
			wsMesType, dataBytes, err := wsConnection.Read(ctx)
			if err != nil {
				processError(ctx, err)
				return
			}
			log.Printf("recv t:%d, msg %s", wsMesType, string(dataBytes))
			message := WsMessage{}
			err = json.Unmarshal(dataBytes, &message)
			if err != nil {
				processError(ctx, err)
				return
			}
			response, err := message.processRequest(deciderId)
			if err != nil {
				processError(ctx, err)
				return
			}
			dataBytesResp, err := json.Marshal(response)
			if err != nil {
				processError(ctx, err)
				return
			}
			wsConnection.Write(ctx, websocket.MessageText, dataBytesResp)
		}
	})
}
