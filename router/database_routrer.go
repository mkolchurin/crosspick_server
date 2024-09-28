package router

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkolchurin/crosspick_server/database"
)

func InitDatabaseRouter(r *gin.Engine) {
	r.GET("/api/v1/db/:op", func(ctx *gin.Context) {
		operation := ctx.Param("op")
		// idString := ctx.Param("id")
		switch operation {
		case "deciders":
			deciders, err := database.GetDeciders()
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, deciders)
		case "maps":
			maps, err := database.GetMaps(0, -1)
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, maps)
		case "roles":
			roles, err := database.GetRoles()
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, roles)
		case "users":
			users, err := database.GetUsers()
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, users)
		default:
			ctx.JSON(400, gin.H{"error": "Unknown operation"})
		}
	})
	r.POST("/api/v1/db/:op", func(ctx *gin.Context) {
		operation := ctx.Param("op")
		switch operation {
		case "deciders":
			var decider database.Deciders
			err := ctx.BindJSON(&decider)
			if err != nil {
				processError(ctx, err)
				return
			}
			err = database.CreateDecider(decider.Title, decider.Description, decider.CreatorId, decider.Maps)
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, gin.H{"message": "Decider created"})
		case "maps":
			var maps database.Maps
			err := ctx.BindJSON(&maps)
			if err != nil {
				processError(ctx, err)
				return
			}
			err = database.InsertMapByStruct(&maps)
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, gin.H{"message": "Map created"})
		case "roles":
			var role database.Roles
			err := ctx.BindJSON(&role)
			if err != nil {
				processError(ctx, err)
				return
			}
			err = database.CreateRole(role.Name)
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, gin.H{"message": "Role created"})
		case "users":
			var user database.Users
			err := ctx.BindJSON(&user)
			if err != nil {
				processError(ctx, err)
				return
			}
			err = database.CreateUser(user.Username, user.Password, user.Email, user.Roles)
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, gin.H{"message": "User created"})
		default:
			ctx.JSON(400, gin.H{"error": "Unknown operation"})
		}
	})
	r.DELETE("/api/v1/db/:op/:id", func(ctx *gin.Context) {
		operation := ctx.Param("op")
		idString := ctx.Param("id")
		switch operation {
		case "deciders":
			id, err := strconv.ParseUint(idString, 10, 64)
			if err != nil {
				processError(ctx, err)
				return
			}
			err = database.DeleteDecider(uint(id))
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, gin.H{"message": "Decider deleted"})
		case "maps":
			err := database.DeleteMap(idString)
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, gin.H{"message": "Map deleted"})
		case "roles":
			err := database.DeleteRole(idString)
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, gin.H{"message": "Role deleted"})
		case "users":
			err := database.DeleteUser(idString)
			if err != nil {
				processError(ctx, err)
				return
			}
			ctx.JSON(200, gin.H{"message": "User deleted"})
		default:
			ctx.JSON(400, gin.H{"error": "Unknown operation"})

		}
	})

}
