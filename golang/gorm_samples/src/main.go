package main

import (
	"psgres/controllers"
	"psgres/models"

	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Init database
	models.ConnectDatabase()
	//log.Println(controllers.FindPosts)

	router := gin.Default()

  router.ForwardedByClientIP = true
  router.SetTrustedProxies([]string{"127.0.0.1"})

  router.StaticFile("/favicon.ico", "./favicon.ico")

	v1 := router.Group("/api/v1")

	v1.POST("/posts", controllers.CreateUser)
	v1.GET("/posts", controllers.FindPosts)
	//v1.GET("/posts/:id", controllers.FindPost)
	//v1.PATCH("/posts/:id", controllers.UpdatePost)
	//v1.DELETE("/posts/:id", controllers.DeletePost)

	v1.GET("/hello", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	//router.Run("0.0.0.0:8080")


	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
