package boot

import (
	"cc-backend-root-console/app/boot/application/route"
)

type Application interface {
	Boot()
	BuildRouterGroup() route.Group
	Listen()
}

func BuildApplicationInstance() Application {
	return &App{}
}
