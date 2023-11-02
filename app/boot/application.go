package boot

import (
	"cc-backend-root-console/app/boot/route"
	"cc-backend-root-console/routes"
	"github.com/gin-gonic/gin"
)

type App struct {
	engine *gin.Engine
	router route.Router
}

func (app *App) Boot() {
	app.engine = gin.Default()
}

func (app *App) setupRoutes() {
	app.router = route.BuildRouterInstance(app.engine)
	app.router.ApplyRoutes(app.BuildRouterGroup())
}

func (app *App) BuildRouterGroup() route.Group {
	return routes.BuildGroupInstance()
}
