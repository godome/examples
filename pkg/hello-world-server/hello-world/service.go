package helloworld

import (
	"fmt"

	"github.com/godome/godome/pkg/module"
	"github.com/godome/godome/pkg/provider"
)

const ServiceType provider.ProviderType = "HelloworldService"

type HelloworldService interface {
	provider.Provider
	SayHello(name string) string
}

type helloworldService struct {
	module       module.Module
	providerType provider.ProviderType
}

func newHelloworldService(m module.Module) HelloworldService {
	return &helloworldService{
		module:       m,
		providerType: ServiceType,
	}
}

func (r *helloworldService) GetType() provider.ProviderType {
	return r.providerType
}

func (r *helloworldService) SayHello(name string) string {
	r.module.Logger().Info("helloworldService has been called")
	return fmt.Sprintf("hello %s", name)
}
