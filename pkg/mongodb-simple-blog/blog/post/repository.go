package post

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/godome/godome/pkg/component/adapter"
	"github.com/godome/godome/pkg/component/provider/repository"
	mongoPlugin "github.com/godome/plugins/pkg/mongo-plugin"
)

const PostRepositoryName = "PostRepository"

type PostRepository interface {
	repository.Repository
	GetOne(id string) (*PostEntity, error)
	Create(data *PostEntity) (*PostEntity, error)
}

type postRepository struct {
	repository.Repository
}

func newPostRepository(a adapter.Adapter) PostRepository {
	r := repository.NewRepository(PostRepositoryName)
	r.SetAdapter(a)

	return &postRepository{
		Repository: r,
	}
}

func (r *postRepository) getPostCollection() mongoPlugin.MongoCollection {
	return r.
		Repository.
		GetAdapter(mongoPlugin.AdapterName).(mongoPlugin.MongoAdapter).
		Collection("post")
}

func (r *postRepository) GetOne(id string) (*PostEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	foundItem := new(PostEntity)
	r.getPostCollection().FindOne(ctx, &PostEntity{ID: id}).Decode(foundItem)
	if foundItem.ID == "" {
		return nil, fmt.Errorf("post with id [%s] is not found", id)
	}
	return foundItem, nil
}

func (r *postRepository) Create(data *PostEntity) (*PostEntity, error) {
	data.ID = fmt.Sprintf("%d", rand.Int())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// avoid duplications
	foundItem := new(PostEntity)
	r.getPostCollection().FindOne(ctx, &PostEntity{ID: data.ID}).Decode(foundItem)
	if foundItem.ID != "" {
		return nil, fmt.Errorf("duplication error")
	}

	// operation
	_, err := r.getPostCollection().InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
