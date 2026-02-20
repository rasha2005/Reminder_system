package routes

import (
	"reminder-system/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuditRoutes(r *gin.Engine) {
	r.GET("/audit", handlers.GetAuditLogs)
}
