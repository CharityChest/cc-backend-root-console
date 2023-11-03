package http

import (
	"cc-backend-root-console/app/providers"
)

type kernel struct {
	providers []providers.ServiceProvider
}

func (k *kernel) GetServiceProviders() []providers.ServiceProvider {
	return k.providers
}

func (k *kernel) AddServiceProvider(provider providers.ServiceProvider) {
	k.providers = append(k.providers, provider)
}

func NewKernel() Kernel {
	k := &kernel{}

	k.AddServiceProvider(&providers.RouteServiceProvider{})

	return k
}
