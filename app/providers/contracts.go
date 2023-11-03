package providers

import (
	"cc-backend-root-console/app/boot/application/application/abstracts"
)

type Handler func(abstracts.Application)

func BuildEmptyHandlerInstance() Handler {
	return func(app abstracts.Application) {}
}

type ServiceProvider interface {
	Handle(Handler, abstracts.Application) Handler
}
