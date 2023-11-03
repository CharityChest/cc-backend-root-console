package boot

import (
	"cc-backend-root-console/app/boot/application/route"
	"cc-backend-root-console/routes"
	"github.com/gin-gonic/gin"
)

type App struct {
	engine *gin.Engine
	router route.Router
}

func (app *App) Boot() {
	app.engine = gin.Default()
	app.setupRoutes()
}

func (app *App) setupRoutes() {
	app.router = route.BuildRouterInstance(app.engine)
	group := app.BuildRouterGroup()
	app.router.ApplyRoutes(group)
}

func (app *App) Listen() {
	err := app.engine.Run()
	if err != nil {
		panic(err)
	}
}

func (app *App) BuildRouterGroup() route.Group {
	return routes.BuildGroupInstance()
}
