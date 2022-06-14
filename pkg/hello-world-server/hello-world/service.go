package helloworld

import (
	"fmt"

	"github.com/godome/godome/pkg/component/provider"
)

const HelloworldServiceName = "HelloworldService"

type HelloworldService interface {
	provider.Provider
	SayHello(name string) string
}

type helloworldService struct {
	provider.Provider
}

func newHelloworldService() HelloworldService {
	return &helloworldService{
		Provider: provider.NewProvider(HelloworldServiceName),
	}
}

func (r *helloworldService) SayHello(name string) string {
	r.Logger().Info("helloworldService.SayHello has been called")
	return fmt.Sprintf("hello %s", name)
}
