package app

import (
	"github.com/edaywalid/chat-app/configs"
	"github.com/edaywalid/chat-app/internal/db"
	"github.com/edaywalid/chat-app/internal/handlers"
	"github.com/edaywalid/chat-app/internal/middlewares"
	"github.com/edaywalid/chat-app/internal/repositories"
	"github.com/edaywalid/chat-app/internal/services"
	"github.com/edaywalid/chat-app/internal/websocket"
	"github.com/edaywalid/chat-app/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type App struct {
	Databases    *Databases
	Config       *configs.Config
	Repositories *Repoisitories
	Services     *Services
	Handlers     *Handlers
	Middlewares  *Middlewares
	Managers     *Managers
}

type (
	Databases struct {
		postgres *gorm.DB
		mongo    *mongo.Client
	}
	Repoisitories struct {
		UserRepository *repositories.UserRepository
	}
	Services struct {
		AuthService  *services.AuthService
		JwtService   *services.JwtService
		EmailService *utils.EmailService
		ChatService  *services.ChatService
		RedisService *services.RedisService
	}
	Handlers struct {
		AuthHandler      *handlers.AuthHandler
		WebSocketHandler *handlers.WebSocketHandler
	}
	Middlewares struct {
		AuthMiddleware *middlewares.AuthMiddleware
		CorsMiddleware *middlewares.CorsMiddleware
	}
	Managers struct {
		WsManager *websocket.Manager
	}
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

	mongo, err := db.InitMongo(a.Config)
	if err != nil {
		panic(err)
	}

	a.Databases = &Databases{
		postgres: postgres,
		mongo:    mongo,
	}
}

func (a *App) initRepositories() {
	a.Repositories = &Repoisitories{
		UserRepository: repositories.NewUserRepository(a.Databases.postgres),
	}
}

func (a *App) initServices() {
	a.Services = &Services{
		JwtService:   services.NewJwtService(a.Config),
		EmailService: utils.NewEmailService(a.Config),
		RedisService: services.NewRedisService(a.Config),
	}
	a.Services.AuthService = services.NewAuthService(
		a.Repositories.UserRepository,
		a.Services.JwtService,
		a.Services.EmailService,
	)
	a.Services.ChatService = services.NewChatService(
		a.Managers.WsManager,
		a.Services.RedisService,
	)
}

func (a *App) initManagers() {
	a.Managers = &Managers{
		WsManager: websocket.NewManger(),
	}
}

func (a *App) initHandlers() {
	a.Handlers = &Handlers{
		AuthHandler:      handlers.NewAuthHandler(a.Services.AuthService),
		WebSocketHandler: handlers.NewWebSocketHandler(a.Managers.WsManager, a.Services.ChatService),
	}
}

func (a *App) initMiddlewares() {
	a.Middlewares = &Middlewares{
		AuthMiddleware: middlewares.NewAuthMiddleware(a.Services.JwtService),
		CorsMiddleware: middlewares.NewCorsMiddleware(),
	}
}

func (a *App) Init() {
	a.initDatabases()
	a.initRepositories()
	a.initManagers()
	a.initServices()
	a.initHandlers()
	a.initMiddlewares()
}
