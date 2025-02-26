package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"orders/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"orders/internal/db"
	"orders/internal/handlers"
	"orders/internal/repositories"
	"orders/internal/services"
)

type Server struct {
	Env  string
	Port string
}

func New(env string, port string) *Server {
	return &Server{
		Env:  env,
		Port: port,
	}
}

func (s Server) Run() {
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	repo := repositories.New(dbConn)

	service := services.New(repo)

	handler := handlers.New(service)

	handler.RegisterRoutes(r)

	if s.Env != config.PROD_ENV {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	server := &http.Server{
		Addr:    ":" + s.Port, // Define your address
		Handler: r,
	}

	fmt.Println("Server start at localhost:" + s.Port)
	// Run server in a goroutine so it doesn't block
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("ListenAndServe() error: %v\n", err)
		}
	}()

	// Create a channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until an interrupt signal is received
	<-quit
	fmt.Println("Shutting down server...")

	// Set a deadline for the server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown Failed: %v\n", err)
	} else {
		fmt.Println("Server exited gracefully")
	}
}
