package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"segwise/docs"
	"segwise/helpers"
	_ "segwise/utils"

	_ "segwise/clients/postgres"

	_ "segwise/clients/redis"
	"segwise/config"
	"segwise/routes"
	"syscall"
	"time"

	_ "segwise/docs" // Import Swagger docs

	swaggerFiles "github.com/swaggo/files" // Make sure this alias is correct

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.uber.org/zap"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, trell-auth-token, trell-app-version-int, creator-space-auth-token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func GracefulShutdown(server *http.Server) {
	stopper := make(chan os.Signal, 1)
	// Listen for interrupt and SIGTERM signals
	signal.Notify(stopper, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stopper
		zap.L().Info("Shutting down gracefully...")

		// Create a context with a timeout for shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Shut down the server
		if err := server.Shutdown(ctx); err != nil {
			zap.L().Error("Server shutdown failed", zap.Error(err))
			return
		}
		zap.L().Info("Server exited gracefully")
	}()
}

// @title Event Trigger API
// @version 1.0
// @description This API allows users to create and manage event triggers
// @host localhost:4000
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.Get()
	port := config.Get().ServerPort
	if port == "" {
		port = "4000"
	}
	router := gin.Default()

	router.Use(CORSMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.Routes(router)
	go helpers.StartScheduler()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	docs.SwaggerInfo.Host = config.Get().Host

	GracefulShutdown(server)

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Fatal("Server failed to start", zap.Error(err))
	}
}
