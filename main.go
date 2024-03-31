package main

import (
	"database/sql"
	"fmt"
	"scrim-api/database"
	"scrim-api/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	// Initialize connection constants.
	HOST     = "aws-0-ap-southeast-1.pooler.supabase.com"
	DATABASE = "postgres"
	USER     = "postgres.pkeejyrcevjrgrgljqfw"
	PASSWORD = "HSgyWmqlUfz2F7Xd"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Initialize connection string.
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", HOST, USER, PASSWORD, DATABASE)

	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")

	database.SetDB(db)

	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/ping", handler.HandlerPing)
	r.POST("/register", handler.HandlerRegister)
	r.POST("/login", handler.HandlerLogin)
	r.POST("change-role", handler.HandlerChangeRole)
	r.POST("kick-member", handler.HandlerKickMember)

	r.POST("/team/user-id/:user-id/create", func(ctx *gin.Context) { handler.HandlerTeamCreate(ctx, ctx.Param("user-id")) })
	r.PUT("/team", handler.HandlerTeamUpdate)
	r.PUT("/team/join", handler.HandlerTeamJoin)
	r.GET("team/:team-id/member", func(ctx *gin.Context) { handler.HandlerTeamMemberGet(ctx, ctx.Param("team-id")) })
	r.GET("team/:team-id/detail", func(ctx *gin.Context) { handler.HandlerTeamDetailGet(ctx, ctx.Param("team-id")) })
	r.PUT("team/:team-id/invite-code", func(ctx *gin.Context) { handler.HandlerTeamInviteCodeGet(ctx, ctx.Param("team-id")) })

	r.POST("/scrim", handler.HandlerScrimPost)
	r.POST("/scrim/offer", handler.HandlerScrimMakeOffer)
	r.PUT("/scrim/offer/accept", handler.HandlerScrimAcceptOffer)
	r.DELETE("/scrim/cancel", handler.HandlerScrimCancelMatch)
	r.DELETE("/scrim", handler.HandlerScrimDelete)

	r.Run()
}

// env := os.Getenv("HOMEPATH")
// 	log.Println(env)
