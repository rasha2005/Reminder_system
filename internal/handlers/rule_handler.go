package handlers

import (
	"reminder-system/internal/database"
	"reminder-system/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateRule(c *gin.Context) {
	var rule models.ReminderRule

	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	rule.IsActive = true

	database.DB.Create(&rule)

	audit := models.AuditLog{
		RuleID:      &rule.ID,
		Event:       "rule_created",
		TriggeredAt: time.Now(),
		Message:     "Reminder rule created",
	}
	database.DB.Create(&audit)

	c.JSON(201, rule)
}

func GetRules(c *gin.Context) {
	var rules []models.ReminderRule
	database.DB.Where("is_active = ?", true).Find(&rules)
	c.JSON(200, rules)
}

func UpdateRule(c *gin.Context) {
	id := c.Param("id")

	var rule models.ReminderRule
	if err := database.DB.First(&rule, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Rule not found"})
		return
	}

	var input models.ReminderRule
	c.ShouldBindJSON(&input)

	rule.Name = input.Name
	rule.Offset = input.Offset
	rule.Unit = input.Unit
	rule.IsActive = input.IsActive

	database.DB.Save(&rule)

	audit := models.AuditLog{
		RuleID:      &rule.ID,
		Event:       "rule_updated",
		TriggeredAt: time.Now(),
		Message:     "Reminder rule updated",
	}
	database.DB.Create(&audit)

	c.JSON(200, rule)
}

func DeleteRule(c *gin.Context) {
	id := c.Param("id")

	var rule models.ReminderRule
	if err := database.DB.First(&rule, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Rule not found"})
		return
	}

	database.DB.Delete(&rule)

	audit := models.AuditLog{
		RuleID:      &rule.ID,
		Event:       "rule_deleted",
		TriggeredAt: time.Now(),
		Message:     "Reminder rule deleted",
	}
	database.DB.Create(&audit)

	c.JSON(200, gin.H{"message": "Rule deleted"})
}

func Update_isActive(c *gin.Context) {
	id := c.Param("id")
	var rule models.ReminderRule
	if err := database.DB.First(&rule, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Rule not found"})
		return
	}

	var input struct {
		IsActive bool `json:"isActive"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&rule).
		Update("is_active", input.IsActive).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update rule"})
		return
	}

	audit := models.AuditLog{
		RuleID:      &rule.ID,
		Event:       "rule_status_changed",
		TriggeredAt: time.Now(),
		Message:     "Reminder rule active status updated",
	}
	database.DB.Create(&audit)

	c.JSON(200, gin.H{
		"message":  "Rule active status updated",
		"isActive": input.IsActive,
	})
}
