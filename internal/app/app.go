package app

import (
	"github.com/edaywalid/chat-app/configs"
	"github.com/edaywalid/chat-app/internal/db"
	"github.com/edaywalid/chat-app/internal/handlers"
	"github.com/edaywalid/chat-app/internal/repositories"
	"github.com/edaywalid/chat-app/internal/services"
	"gorm.io/gorm"
)

type App struct {
	DB           *gorm.DB
	Config       *configs.Config
	Repositories *Repoisitories
	Services     *Services
	Handlers     *Handlers
	Middlewares  *Middlewares
}

type (
	Repoisitories struct {
	}
	Services struct {
	}
	Handlers struct {
	}
	Middlewares struct{}
)

func NewApp(path string) (*App, error) {
	config, err := configs.LoadConfig(path)
	if err != nil {
		return nil, err
	}

	db, err := db.InitDB(config)
	if err != nil {
		return nil, err
	}
	app := &App{
		DB:     db,
		Config: config,
	}

	app.Init()

	return app, nil
}

func (a *App) initRepositories() {
	a.Repositories = &Repoisitories{
	}
}

func (a *App) initServices() {
	a.Services = &Services{
	}
}

func (a *App) initHandlers() {
	a.Handlers = &Handlers{
	}
}

func (a *App) initMiddlewares() {
	a.Middlewares = &Middlewares{}
}

func (a *App) Init() {
	a.initRepositories()
	a.initServices()
	a.initHandlers()
	a.initMiddlewares()
}
