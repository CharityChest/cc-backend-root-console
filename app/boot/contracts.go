package boot

import (
	"cc-backend-root-console/app/boot/route"
)

type Application interface {
	Boot()
	BuildRouterGroup() route.Group
}

func BuildApplicationInstance() Application {
	return &App{}
}
