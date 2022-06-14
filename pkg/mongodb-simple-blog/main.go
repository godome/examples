package main

import (
	"github.com/godome/examples/pkg/mongodb-simple-blog/blog/post"
	"github.com/godome/godome/pkg/config"
	fiberExposure "github.com/godome/plugins/pkg/fiber"
	mongoAdapter "github.com/godome/plugins/pkg/mongo"
)

func main() {
	config.NewConfig().Set("DEBUG", "true")
	// Init (mongodb) adapter
	adapter := mongoAdapter.NewMongoAdapter("mongodb://localhost:27017", "godome-blog", true)

	// Init modules
	postModule := post.NewPostModule(adapter)

	// Init (fiber http) exposure
	fiberExposure.
		NewFiberExposure("3000", nil).
		ExposeModule(postModule).
		Run()
}
