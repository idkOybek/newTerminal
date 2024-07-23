package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	_ "github.com/idkOybek/docs"
	"github.com/idkOybek/internal/handlers"
	"github.com/idkOybek/internal/logger"
	"github.com/idkOybek/internal/middleware"
	"github.com/idkOybek/internal/repository"
	"github.com/idkOybek/internal/services"
	"github.com/idkOybek/internal/utils"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title New Terminal API
// @version 1.0
// @description This is a new terminal API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host verified-gorilla-yearly.ngrok-free.app
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logger.InitLogger()

	err := godotenv.Load()
	if err != nil {
		logger.ErrorLogger.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		logger.ErrorLogger.Fatalf("Could not connect to the database: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	fiscalRepo := repository.NewFiscalRepository(db)
	terminalRepo := repository.NewTerminalRepository(db)

	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	fiscalService := services.NewFiscalService(fiscalRepo)
	terminalService := services.NewTerminalService(terminalRepo, fiscalService)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	fiscalHandler := handlers.NewFiscalHandler(fiscalService)
	terminalHandler := handlers.NewTerminalHandler(terminalService)

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(utils.JSONMiddleware)
	r.Use(middleware.CORSMiddleware)
	r.Use(utils.LoggerMiddleware)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/auth", authHandler.AuthRoutes())
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Mount("/users", userHandler.Routes())
			r.Mount("/fiscal", fiscalHandler.Routes())
			r.Mount("/terminal", terminalHandler.Routes())
		})
	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	logger.InfoLogger.Printf("Server starting on port %s", port)
	logger.ErrorLogger.Fatal(http.ListenAndServe(":"+port, r))
}
