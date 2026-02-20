package seed

import (
	"fmt"
	"time"

	"reminder-system/internal/database"
	"reminder-system/internal/models"
)

func SeedTasks() {
	var count int64

	database.DB.Model(&models.Task{}).Count(&count)

	
	if count > 0 {
		fmt.Println("Tasks already seeded")
		return
	}

	tasks := []models.Task{
		{
			Title:       "Submit report",
			Description: "Submit monthly report",
			DueDate:     time.Now().Add(30 * time.Minute),
		},
		{
			Title:       "Doctor appointment",
			Description: "Annual health check",
			DueDate:     time.Now().Add(2 * time.Hour),
		},
		{
			Title:       "Project deadline",
			Description: "Complete backend module",
			DueDate:     time.Now().Add(24 * time.Hour),
		},
		{
			Title:       "Buy groceries",
			Description: "Milk, eggs, bread",
			DueDate:     time.Now().Add(48 * time.Hour),
		},
		{
			Title:       "Call client",
			Description: "Discuss new contract",
			DueDate:     time.Now().Add(1 * time.Hour),
		},
	}

	database.DB.Create(&tasks)

	fmt.Println("5 sample tasks seeded successfully")
}