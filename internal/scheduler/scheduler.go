package scheduler

import (
	"fmt"
	"time"

	"reminder-system/internal/database"
	"reminder-system/internal/models"
)

func StartScheduler() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			checkReminders()
		}
	}()

	fmt.Println("Scheduler started...")
}

func checkReminders() {
	var tasks []models.Task
	var rules []models.ReminderRule

	database.DB.Find(&tasks)
	database.DB.Where("is_active = ?", true).Find(&rules)

	now := time.Now()

	for _, task := range tasks {
		for _, rule := range rules {
			triggerTime := calculateTriggerTime(task.DueDate, rule.Offset, rule.Unit)

			var count int64
			database.DB.Model(&models.AuditLog{}).
				Where("task_id = ? AND rule_id = ? AND event = ?", task.ID, rule.ID, "triggered").
				Count(&count)

			if count > 0 {
				continue
			}

			if now.After(triggerTime) &&
				now.Before(task.DueDate) {
				message := fmt.Sprintf("Reminder: Task '%s' is due at %s", task.Title, task.DueDate.Format(time.RFC1123))
				fmt.Println(message)

				audit := models.AuditLog{
					TaskID:      &task.ID,
					RuleID:      &rule.ID,
					Event:       "triggered",
					TriggeredAt: now,
					Message:     message,
				}
				database.DB.Create(&audit)
			}
		}
	}
}

func calculateTriggerTime(due time.Time, offset int, unit string) time.Time {
	switch unit {
	case "minutes":
		return due.Add(-time.Duration(offset) * time.Minute)
	case "hours":
		return due.Add(-time.Duration(offset) * time.Hour)
	case "days":
		return due.AddDate(0, 0, -offset)
	default:
		return due
	}
}
