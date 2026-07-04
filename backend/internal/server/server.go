package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kacperkg/vellum/internal/config"
	"github.com/kacperkg/vellum/internal/database"
	"github.com/kacperkg/vellum/internal/handlers"
	"github.com/kacperkg/vellum/internal/repositories"
	"github.com/kacperkg/vellum/internal/routes"
	"github.com/kacperkg/vellum/internal/services"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config
	Router *gin.Engine
	DB     *gorm.DB
}

func New() (*Application, error) {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		return nil, err
	}

	if err := database.Migrate(db); err != nil {
		return nil, err
	}

	jwtExpiry, err := time.ParseDuration(cfg.JWTExpiry)
	if err != nil {
		return nil, err
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
	}))

	// Repositories
	userRepository := repositories.NewGormUserRepository(db)
	userSettingsRepository := repositories.NewGormUserSettingsRepository(db)

	// Services
	userService := services.NewUserService(
		userRepository,
		userSettingsRepository,
		cfg.JWTSecret,
		jwtExpiry,
	)

	// Handlers
	authHandler := handlers.NewAuthHandler(userService)

	// Routes
	routes.Register(router, authHandler)

	return &Application{
		Config: cfg,
		Router: router,
		DB:     db,
	}, nil
}

func (a *Application) Run() error {
	return a.Router.Run(":" + a.Config.Port)
}