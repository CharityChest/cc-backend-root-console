package providers

import (
	"cc-backend-root-console/app/boot/application/application/abstracts"
	"cc-backend-root-console/app/boot/application/route"
	"cc-backend-root-console/routes"
)

type RouteServiceProvider struct {
}

func (r *RouteServiceProvider) Handle(next Handler, app abstracts.Application) Handler {
	group := routes.BuildGroupInstance()
	app.Set(route.GroupInterfaceName, group)

	return next
}
