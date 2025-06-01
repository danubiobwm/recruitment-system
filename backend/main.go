package main

import (
	"log"

	"github.com/danubiobwm/recruitment-system/database"
	"github.com/danubiobwm/recruitment-system/models"
	"github.com/danubiobwm/recruitment-system/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	db := database.ConnectDB()

	if err := db.AutoMigrate(&models.User{}, &models.Job{}, &models.Candidate{}); err != nil {
		log.Fatalf("Erro ao migrar tabelas: %v", err)
	}

	r := gin.Default()

	routes.SetupRoutes(r)
	routes.SetupJobRoutes(r)
	routes.SetupCandidateRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
