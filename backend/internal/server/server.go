package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kacperkg/vellum/internal/config"
	"github.com/kacperkg/vellum/internal/database"
	"github.com/kacperkg/vellum/internal/routes"
	"gorm.io/gorm"
)



type Application struct {
	Config *config.Config
	Router *gin.Engine
	DB *gorm.DB
}

func New() (*Application, error) {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		return nil, err
	}

	err = database.Migrate(db)
	if err != nil {
    	return nil, err
	}

	fmt.Println("Connected to PostgreSQL")


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

	routes.Register(router)

	return &Application{
		Config: cfg,
		DB: db,
		Router: router,
	}, nil
}

func (a *Application) Run() error {
	return a.Router.Run(":" + a.Config.Port)
}