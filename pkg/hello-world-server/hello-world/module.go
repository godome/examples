package helloworld

import "github.com/godome/godome/pkg/component/module"

func NewHelloworldModule() module.Module {
	m := module.NewModule("helloWorld")
	// Add service
	m.SetProvider(newHelloworldService())
	// Add (fiber http server) handler
	m.SetProvider(newHelloworldHandler(m))
	return m
}
