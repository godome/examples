package main

import (
	"github.com/godome/examples/pkg/mongodb-simple-blog/blog/post"
	"github.com/godome/godome/pkg/config"
	fiberPlugin "github.com/godome/plugins/pkg/fiber-plugin"
	mongoPlugin "github.com/godome/plugins/pkg/mongo-plugin"
)

func main() {
	config.NewConfig().Set("DEBUG", "true")
	// Init (mongodb) adapter
	adapter := mongoPlugin.NewMongoAdapter("mongodb://localhost:27017", "godome-blog", true)

	// Init modules
	postModule := post.NewPostModule(adapter)

	// Init (fiber http) exposure
	fiberPlugin.
		NewFiberExposure("3000", nil).
		ExposeModule(postModule).
		Run()
}
