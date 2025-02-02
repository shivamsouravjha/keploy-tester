package routes

import (
	"segwise/controllers"
	"segwise/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	v1 := r.Group("/api")

	{
		// Authentication
		v1.POST("/register", controllers.RegisterUser)
		v1.POST("/login", controllers.LoginUser)

		// Protected Routes
		auth := v1.Group("/")
		auth.Use(middleware.JWTAuthMiddleware())

		// Trigger management
		auth.POST("/triggers", controllers.CreateTrigger)
		auth.GET("/triggers", controllers.GetTriggers)
		auth.GET("/triggers/:id", controllers.GetTriggerByID)
		auth.PUT("/triggers/:id", controllers.UpdateTrigger)
		auth.DELETE("/triggers/:id", controllers.DeleteTrigger)

		// Trigger Execution
		auth.POST("/triggers/:id/execute", controllers.ExecuteTrigger)

		// Event Logs
		auth.GET("/events", controllers.GetActiveEvents)
		auth.GET("/events/archived", controllers.GetArchivedEvents)
		auth.DELETE("/events/purge", controllers.PurgeOldEvents)

		// **Testing API Endpoints**
		auth.POST("/triggers/test/scheduled", controllers.TestScheduledTrigger)
		auth.POST("/triggers/test/api", controllers.TestAPITrigger)

	}
}
