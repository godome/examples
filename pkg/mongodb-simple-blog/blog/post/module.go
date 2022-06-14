package post

import (
	"github.com/godome/godome/pkg/component/adapter"
	"github.com/godome/godome/pkg/component/module"
)

func NewPostModule(a adapter.Adapter) module.Module {
	m := module.NewModule("post")
	// Add Repository
	m.SetProvider(newPostRepository(a))
	// Add service
	m.SetProvider(newPostService(m))
	// Add (fiber http server) handler
	m.SetProvider((newPostHandler(m)))
	return m
}
