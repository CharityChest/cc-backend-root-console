package http

import "cc-backend-root-console/app/providers"

type Kernel interface {
	GetServiceProviders() []providers.ServiceProvider
	AddServiceProvider(providers.ServiceProvider)
}
