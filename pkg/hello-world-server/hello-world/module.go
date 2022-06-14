package helloworld

import "github.com/godome/godome/pkg/module"

const moduleName = "helloWorld"

func NewHelloworldModule() module.Module {
	m := module.NewModule(moduleName)
	// Add service
	m.SetProvider(newHelloworldService(m))
	// Add (fiber http server) handler
	m.SetProvider(newHelloworldHandler(m))
	return m
}
