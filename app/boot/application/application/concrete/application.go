package concrete

import (
	"cc-backend-root-console/app/boot/application/route"
	"cc-backend-root-console/app/http"
	"cc-backend-root-console/app/providers"
	"github.com/gin-gonic/gin"
	"reflect"
)

type appInstance struct {
	engine    *gin.Engine
	router    route.Router
	container map[interface{}]interface{}
	kernel    http.Kernel
}

func (app *appInstance) Boot() {
	app.setupKernel()
	app.engine = gin.Default()
	app.processServiceProviders()
	app.setupRoutes()
}

func (app *appInstance) setupKernel() {
	app.kernel = http.NewKernel()
}

func (app *appInstance) processServiceProviders() {
	r := providers.BuildEmptyHandlerInstance()
	for _, provider := range app.kernel.GetServiceProviders() {
		r = provider.Handle(r, app)

		if r == nil {
			break
		}
	}
}

func (app *appInstance) setupRoutes() {
	app.router = route.BuildRouterInstance(app.engine)
	group := app.Get(route.GroupInterfaceName).(route.Group)
	app.router.ApplyRoutes(group)
}

func (app *appInstance) Listen() {
	err := app.engine.Run()
	if err != nil {
		panic(err)
	}
}

func (app *appInstance) Get(key interface{}) interface{} {
	return app.container[app.buildContainerKey(key)]
}

func (app *appInstance) Set(key interface{}, value interface{}) {
	app.container[app.buildContainerKey(key)] = value
}

func (app *appInstance) buildContainerKey(key interface{}) string {
	switch key.(type) {
	case string:
		return key.(string)
	case int:
		return key.(string)
	case float64:
		return key.(string)
	default:
		e := reflect.TypeOf(key).Elem()
		return e.PkgPath() + "." + e.Name()
	}
}

func BuildAppInstance() *appInstance {
	return &appInstance{container: make(map[interface{}]interface{})}
}
