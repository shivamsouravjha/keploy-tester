package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"segwise/config"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
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

func main() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	logger, _ := zapConfig.Build()
	zap.ReplaceGlobals(logger)
	config.Get()
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Example route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running!"})
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	GracefulShutdown(server)

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Fatal("Server failed to start", zap.Error(err))
	}
}
