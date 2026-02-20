package routes

import (
	"reminder-system/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRuleRoutes(r *gin.Engine) {
	r.POST("/rules", handlers.CreateRule)
	r.GET("/rules", handlers.GetRules)
	r.PUT("/rules/:id", handlers.UpdateRule)
	r.DELETE("/rules/:id", handlers.DeleteRule)
	r.PATCH("/rules/:id", handlers.Update_isActive)
}
