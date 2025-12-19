package main

import (
	"go-matrix-processor/internal/handlers"
	"go-matrix-processor/internal/middleware"
	"go-matrix-processor/internal/service"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Inicializar Fiber con configuración personalizada
	app := fiber.New(fiber.Config{
		AppName: "Matrix QR Service",
	})

	// Middleware de logging
	app.Use(logger.New())

	// Enable CORS (Critical for Frontend)
	app.Use(cors.New())

	// Initialize layers (Manual Dependency Injection)
	// Initialize layers (Manual Dependency Injection)
	matrixService := service.NewMatrixService()
	matrixHandler := handlers.NewMatrixHandler(matrixService)

	authService := service.NewAuthService()
	authHandler := handlers.NewAuthHandler(authService)

	// Define routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Public Routes
	v1.Post("/login", authHandler.Login)

	// Protected Routes (Apply Middleware)
	v1.Post("/process", middleware.Protected(), matrixHandler.ProcessMatrix)

	// Puerto desde variable de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Iniciando servidor Go en el puerto %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}
