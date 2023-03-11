package app

import (
	"os"

	"github.com/PongDev/PongDev_agnos_backend_internship_2023/routes"
	"github.com/gin-gonic/gin"
)

type App struct {
	engine *gin.Engine
}

func NewApp() *App {
	trustProxy := os.Getenv("TRUSTED_PROXY")
	ginMode := os.Getenv("GIN_MODE")

	gin.SetMode(ginMode)
	app := &App{
		engine: gin.Default(),
	}
	app.engine.SetTrustedProxies([]string{trustProxy})
	return app
}

func (app App) SetupRouter() {
	routes.SetupRouter(app.engine)
}

func (app App) Run() {
	app.engine.Run()
}
