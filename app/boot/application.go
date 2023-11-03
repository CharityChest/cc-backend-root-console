package boot

import (
	"cc-backend-root-console/app/boot/application/route"
	"cc-backend-root-console/routes"
	"github.com/gin-gonic/gin"
)

type app struct {
	engine *gin.Engine
	router route.Router
}

func (app *app) Boot() {
	app.engine = gin.Default()
	app.setupRoutes()
}

func (app *app) setupRoutes() {
	app.router = route.BuildRouterInstance(app.engine)
	group := app.BuildRouterGroup()
	app.router.ApplyRoutes(group)
}

func (app *app) Listen() {
	err := app.engine.Run()
	if err != nil {
		panic(err)
	}
}

func (app *app) BuildRouterGroup() route.Group {
	return routes.BuildGroupInstance()
}
