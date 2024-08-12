package app

import (
	"github.com/edaywalid/chat-app/configs"
	"github.com/edaywalid/chat-app/internal/db"
	"github.com/edaywalid/chat-app/internal/handlers"
	"github.com/edaywalid/chat-app/internal/repositories"
	"github.com/edaywalid/chat-app/internal/services"
	"github.com/edaywalid/chat-app/pkg/utils"
	"gorm.io/gorm"
)

type App struct {
	Databases    *Databases
	Config       *configs.Config
	Repositories *Repoisitories
	Services     *Services
	Handlers     *Handlers
	Middlewares  *Middlewares
}

type (
	Databases struct {
		postgres *gorm.DB
	}
	Repoisitories struct {
		UserRepository *repositories.UserRepository
	}
	Services struct {
		AuthService *services.AuthService
	}
	Handlers struct {
		AuthHandler *handlers.AuthHandler
	}
	Middlewares struct{}
)

func NewApp(path string) (*App, error) {
	config, err := configs.LoadConfig(path)
	if err != nil {
		return nil, err
	}

	app := &App{
		Config: config,
	}

	app.Init()

	return app, nil
}

func (a *App) initDatabases() {
	postgres, err := db.InitPSQL(a.Config)
	if err != nil {
		panic(err)
	}


	a.Databases = &Databases{
		postgres: postgres,
	}
}
func (a *App) initRepositories() {
	a.Repositories = &Repoisitories{
		UserRepository: repositories.NewUserRepository(a.Databases.postgres),
	}
}

func (a *App) initServices() {
	a.Services = &Services{
		AuthService: services.NewAuthService(
			a.Repositories.UserRepository,
			services.NewJwtService(a.Config),
			utils.NewEmailService(a.Config),
		),
	}
}

func (a *App) initHandlers() {
	a.Handlers = &Handlers{
		AuthHandler: handlers.NewAuthHandler(a.Services.AuthService),
	}
}

func (a *App) initMiddlewares() {
	a.Middlewares = &Middlewares{}
}

func (a *App) Init() {
	a.initDatabases()
	a.initRepositories()
	a.initServices()
	a.initHandlers()
	a.initMiddlewares()
}
