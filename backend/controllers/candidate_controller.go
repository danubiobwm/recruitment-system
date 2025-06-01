package controllers

import (
	"net/http"
	"strconv"

	"github.com/danubiobwm/recruitment-system/database"
	"github.com/danubiobwm/recruitment-system/models"
	"github.com/gin-gonic/gin"
)

type UpdateCandidateInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateCandidate(c *gin.Context) {
	var candidate models.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&candidate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create candidate"})
		return
	}

	c.JSON(http.StatusCreated, candidate)
}

func GetCandidates(c *gin.Context) {
	var candidates []models.Candidate
	if err := database.DB.Find(&candidates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve candidates"})
		return
	}
	c.JSON(http.StatusOK, candidates)
}

func GetCandidate(c *gin.Context) {
	id := c.Param("id")

	candidateID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid candidate ID"})
		return
	}

	var candidate models.Candidate
	if err := database.DB.First(&candidate, candidateID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		return
	}

	c.JSON(http.StatusOK, candidate)
}

func UpdateCandidate(c *gin.Context) {
	id := c.Param("id")

	candidateID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid candidate ID"})
		return
	}

	var candidate models.Candidate
	if err := database.DB.First(&candidate, candidateID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		return
	}

	var input UpdateCandidateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Atualize somente os campos que você quer permitir
	candidate.Name = input.Name
	candidate.Email = input.Email

	if err := database.DB.Save(&candidate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update candidate"})
		return
	}

	c.JSON(http.StatusOK, candidate)
}

func DeleteCandidate(c *gin.Context) {
	id := c.Param("id")

	candidateID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid candidate ID"})
		return
	}

	if err := database.DB.Delete(&models.Candidate{}, candidateID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete candidate"})
		return
	}

	c.Status(http.StatusNoContent)
}
