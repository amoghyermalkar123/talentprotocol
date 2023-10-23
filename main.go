package main

import (
	"fmt"
	"os"
	"talentprotocol/api"
	assessmentpipeline "talentprotocol/assessment-pipeline"
	"talentprotocol/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	db, err := db.GetDB("127.0.0.1")
	if err != nil {
		panic(err)
	}
	api := api.Api{
		DB:                  db,
		AssessmentPipelines: assessmentpipeline.NewAssessmentPipeline(db),
	}

	r.Use(cors.Default())
	r.POST("/login", api.Login)
	r.POST("/signup", api.Signup)

	userApi := r.Group("/v1/candidate")
	{
		userApi.GET("/:candidate-email/home", api.HomePage)
		userApi.POST("/:candidate-email/:opening-id/submit", api.SubmitAssignment)
	}

	orgApi := r.Group("/v1/org")
	{
		orgApi.POST("/login", api.OrgLogin)
		orgApi.POST("/register", api.OrgSignup)
		orgApi.POST("/openings", api.CreateJobOpening)
		orgApi.GET("/:orgname/openings", api.GetAllOrgOpenings)
	}

	go api.AssessmentPipelines.AssessmentPipeline()
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
