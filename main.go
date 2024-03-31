package main

import (
	"database/sql"
	"fmt"
	"scrim-api/database"
	"scrim-api/handler"

	"github.com/gin-contrib/cors"

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
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/ping", handler.HandlerPing)
	r.POST("/register", handler.HandlerRegister)
	r.POST("/login", handler.HandlerLogin)
	r.POST("/change-role", handler.HandlerChangeRole)
	r.POST("/kick-member", handler.HandlerKickMember)

	team := r.Group("/team")
	{
		team.POST("/create/user-id/:user-id/", func(ctx *gin.Context) { handler.HandlerTeamCreate(ctx, ctx.Param("user-id")) })
		team.PUT("/join", handler.HandlerTeamJoin)
		team.GET("/member/team-id/:team-id", func(ctx *gin.Context) { handler.HandlerTeamMemberGet(ctx, ctx.Param("team-id")) })
		team.GET("/detail/team-id/:team-id", func(ctx *gin.Context) { handler.HandlerTeamDetailGet(ctx, ctx.Param("team-id")) })
		team.PUT("/invite-code/team-id/:team-id", func(ctx *gin.Context) { handler.HandlerTeamInviteCodeGet(ctx, ctx.Param("team-id")) })
		team.PUT("/", handler.HandlerTeamUpdate)
	}

	scrim := r.Group("/scrim")
	{
		scrim.POST("/offer", handler.HandlerScrimMakeOffer)
		scrim.PUT("/accept", handler.HandlerScrimAcceptOffer)
		scrim.DELETE("/cancel", handler.HandlerScrimCancelMatch)
		scrim.POST("/", handler.HandlerScrimPost)
		scrim.DELETE("/", handler.HandlerScrimDelete)
	}

	r.Run()
}

// env := os.Getenv("HOMEPATH")
// 	log.Println(env)
