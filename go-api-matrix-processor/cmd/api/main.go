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

	// Habilitar CORS (Crítico para el Frontend)
	app.Use(cors.New())

	// Inicializar capas (Inyección de dependencias manual)
	matrixService := service.NewMatrixService()
	matrixHandler := handlers.NewMatrixHandler(matrixService)

	authService := service.NewAuthService()
	authHandler := handlers.NewAuthHandler(authService)

	// Definir rutas
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Rutas públicas
	v1.Post("/login", authHandler.Login)

	// Rutas protegidas (Aplicar middleware)
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
