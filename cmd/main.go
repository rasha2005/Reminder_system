package main

import (
	"fmt"
	"os"
	"reminder-system/internal/database"
	"reminder-system/internal/models"
	"reminder-system/internal/routes"
	"reminder-system/internal/scheduler"
	seed "reminder-system/internal/seeds"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("hello world")

	database.Connect()

	database.DB.AutoMigrate(
		&models.Task{},
		&models.ReminderRule{},
		&models.AuditLog{},
	)

	seed.SeedTasks()

	scheduler.StartScheduler()

	r := gin.Default()

	routes.SetupRuleRoutes(r)

	routes.SetupAuditRoutes(r)

	r.Run(os.Getenv("PORT"))

}
