package main

import (
	"fmt"
	"os"
	"talentprotocol/api"
	"talentprotocol/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	db, err := db.GetDB("db")
	if err != nil {
		panic(err)
	}
	api := api.Api{
		DB: db,
	}

	r.Use(cors.Default())

	r.POST("/login", api.Login)
	r.POST("/signup", api.Signup)

	userApi := r.Group("/v1/candidate")
	{
		userApi.GET("/:candidate-email/home", api.HomePage)
	}

	orgApi := r.Group("/v1/org")
	{
		orgApi.POST("/login", api.OrgLogin)
		orgApi.POST("/register", api.OrgSignup)
		orgApi.POST("/openings", api.CreateJobOpening)
	}

	godotenv.Load()
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
