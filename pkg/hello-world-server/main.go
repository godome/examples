package main

import (
	helloworld "github.com/godome/examples/pkg/hello-world-server/hello-world"
	fiberExposure "github.com/godome/plugins/pkg/fiber"
)

func main() {
	// Init module
	helloworldModule := helloworld.NewHelloworldModule()

	// Init (fiber http) exposure
	fiberExposure.
		NewFiberExposure("3000").
		ExposeModule(helloworldModule).
		Run()
}
