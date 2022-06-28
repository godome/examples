package main

import (
	helloworld "github.com/godome/examples/pkg/hello-world-server/hello-world"
	fiberPlugin "github.com/godome/plugins/pkg/fiber-plugin"
)

func main() {
	// Init module
	helloworldModule := helloworld.NewHelloworldModule()

	// Init (fiber http) exposure
	fiberPlugin.
		NewFiberExposure("3000", nil).
		ExposeModule(helloworldModule).
		Run()
}
