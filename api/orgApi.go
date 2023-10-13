package api

import (
	"fmt"
	"net/http"
	"talentprotocol/middlewares"
	"talentprotocol/types"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *Api) OrgSignup(c *gin.Context) {
	orgDetails := &types.Organization{}
	if err := c.Bind(orgDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	profileEmail := orgDetails.OrgEmail
	if err := a.DB.InsertOrganization(orgDetails); err != nil {
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

func (a *Api) OrgLogin(c *gin.Context) {
	orgDetails := &types.OrgLogin{}
	if err := c.Bind(orgDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	profileEmail := orgDetails.Email
	userInfo, err := a.DB.GetOrgDetails(orgDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	tokenString, err := middlewares.GenerateJWT(profileEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"org_details": userInfo, "token": tokenString})
}

func (a *Api) CreateJobOpening(c *gin.Context) {
	opening := &types.JobOpening{}
	if err := c.Bind(opening); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	opening.JobPostedAt = time.Now()
	if err := a.DB.CreateJobOpening(opening); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (a *Api) GetOrgAssignmentByID(c *gin.Context) {
	assignmentId := c.Param("opening-id")

	assignment, err := a.DB.GetAssignmentForOrgOpening(assignmentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"assignment": assignment})
}

func (a *Api) CreateOrgAssignment(c *gin.Context) {
	assignment := &types.OrgAssignments{}
	if err := c.Bind(assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": fmt.Errorf("failed request operation: %v", err).Error()})
		return
	}

	openingID := c.Param("opening-id")
	err := a.DB.InsertOrgAssignment(assignment, openingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": fmt.Errorf("failed db operation: %v", err).Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "succcess"})
}
