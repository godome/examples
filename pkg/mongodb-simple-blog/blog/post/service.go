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
	postRepository PostRepository
}

func newPostService(m module.Module) PostService {
	return &postService{
		Provider:       provider.NewProvider(PostServiceName),
		postRepository: m.GetProvider(PostRepositoryName).(PostRepository),
	}
}

func (r *postService) GetPost(id string) (*PostEntity, error) {
	return r.postRepository.GetOne(id)
}

func (r *postService) CreatePost(name string, description string) (*PostEntity, error) {
	return r.postRepository.Create(&PostEntity{Name: name, Description: description})
}
