package controllers

import (
	"net/http"
	"strconv"

	"github.com/danubiobwm/recruitment-system/database"
	"github.com/danubiobwm/recruitment-system/models"
	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := database.DB
	db.Create(&job)
	c.JSON(http.StatusCreated, job)
}

func GetJobs(c *gin.Context) {
	var jobs []models.Job
	database.DB.Preload("Candidates").Find(&jobs)
	c.JSON(http.StatusOK, jobs)
}

func GetJob(c *gin.Context) {
	id := c.Param("id")
	var job models.Job
	result := database.DB.Preload("Candidates").First(&job, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}
	c.JSON(http.StatusOK, job)
}

func UpdateJob(c *gin.Context) {
	id := c.Param("id")

	jobID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var job models.Job
	if err := database.DB.First(&job, jobID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar job"})
		return
	}

	c.JSON(http.StatusOK, job)
}

func DeleteJob(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Job{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting job"})
		return
	}
	c.Status(http.StatusNoContent)
}

func ListJobs(c *gin.Context) {
	var jobs []models.Job
	if err := database.DB.Preload("Candidates").Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar vagas"})
		return
	}

	c.JSON(http.StatusOK, jobs)
}
