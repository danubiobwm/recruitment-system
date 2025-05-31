package main

import (
	"github.com/danubiobwm/recruitment-system/config"
	"github.com/danubiobwm/recruitment-system/database"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()
	r := routes.SetupRoutes()
	r.Run(":8080")
}
