package handlers

import (
	"reminder-system/internal/database"
	"reminder-system/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAuditLogs(c *gin.Context) {
	var logs []models.AuditLog
	logType := c.Query("type")

	query := database.DB.Order("triggered_at desc")

	if logType == "status" {
		query = query.Where("event IN ?", []string{
			"rule_created",
			"rule_updated",
			"rule_deleted",
			"rule_status_changed",
		})
	}

	if logType == "trigger" {
		query = query.Where("event = ?", "triggered")
	}

	query.Find(&logs)

	c.JSON(200, logs)
}
