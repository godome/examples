package post

import (
	"github.com/godome/godome/pkg/component/module"
	"github.com/godome/godome/pkg/component/provider"
)

const PostServiceName = "PostService"

type PostService interface {
	provider.Provider
	GetPost(id string) (*PostEntity, error)
	CreatePost(name string, description string) (*PostEntity, error)
}

type postService struct {
	provider.Provider
	module module.Module
}

func newPostService(m module.Module) PostService {
	return &postService{
		Provider: provider.NewProvider(PostServiceName),
		module:   m,
	}
}

func (r *postService) getPostRepository() PostRepository {
	return r.module.GetProvider(PostRepositoryName).(PostRepository)
}

func (r *postService) GetPost(id string) (*PostEntity, error) {
	return r.getPostRepository().GetOne(id)
}

func (r *postService) CreatePost(name string, description string) (*PostEntity, error) {
	return r.getPostRepository().Create(&PostEntity{Name: name, Description: description})
}
