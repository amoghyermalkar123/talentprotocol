package api

import (
	"fmt"
	"net/http"
	"talentprotocol/db"
	"talentprotocol/middlewares"
	"talentprotocol/types"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Api struct {
	DB  *db.DB
	Log *zerolog.Logger
}

func (a *Api) Signup(c *gin.Context) {
	userDetails := &types.CandidateDetails{}
	if err := c.Bind(userDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	profileEmail := userDetails.Email
	if err := a.DB.AddCandidateDetails(userDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	tokenString, err := middlewares.GenerateJWT(profileEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (a *Api) Login(c *gin.Context) {
	userDetails := &types.CandidateLogin{}
	if err := c.Bind(userDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	profileEmail := userDetails.Email
	userInfo, err := a.DB.GetCandidateDetails(userDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	tokenString, err := middlewares.GenerateJWT(profileEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_details": userInfo, "token": tokenString})
}

func (a *Api) HomePage(c *gin.Context) {
	email := c.Param("candidate-email")
	jobOpenings, err := a.DB.GetJobOpeningsNotAppliedTo(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"job_openings": jobOpenings})
}

func (a *Api) SubmitAssignment(c *gin.Context) {
	submission := &types.CandidateSubmission{}
	if err := c.Bind(submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	jobOpeningID := c.Param("opening-id")
	candEmail := c.Param("candidate-email")
	asgID := c.Param("assignment-id")

	if err := a.DB.InsertCandidateSubmission(candEmail, jobOpeningID, asgID, submission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	// TODO: add ai-engine analysis to include code analysis and rating in the response
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
