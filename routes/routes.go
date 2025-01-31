package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	v1 := r.Group("/api")

	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Server is running!"})
		})
	}
}
