package routes

import (
	"segwise/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	v1 := r.Group("/api")

	{
		v1.POST("/triggers", controllers.CreateTrigger)
		v1.GET("/triggers", controllers.GetTriggers)
		v1.GET("/triggers/:id", controllers.GetTriggerByID)
		v1.PUT("/triggers/:id", controllers.UpdateTrigger)
		v1.DELETE("/triggers/:id", controllers.DeleteTrigger)
		v1.POST("/triggers/:id/execute", controllers.ExecuteTrigger)

		v1.GET("/events", controllers.GetActiveEvents)
		v1.GET("/events/archived", controllers.GetArchivedEvents)
		v1.DELETE("/events/purge", controllers.PurgeOldEvents)

	}
}
