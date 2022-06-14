package post

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/godome/godome/pkg/component/adapter"
	"github.com/godome/godome/pkg/component/provider/repository"
	mongoAdapter "github.com/godome/plugins/pkg/mongo"
)

const PostRepositoryName = "PostRepository"

type PostRepository interface {
	repository.Repository
	GetOne(id string) (*PostEntity, error)
	Create(data *PostEntity) (*PostEntity, error)
}

type postRepository struct {
	repository.Repository
	postCollection mongoAdapter.MongoCollection
}

func newPostRepository(a adapter.Adapter) PostRepository {
	r := repository.NewRepository(PostRepositoryName)
	r.SetAdapter(a)

	return &postRepository{
		Repository: r,
		postCollection: r.
			GetAdapter(a.Metadata().GetName()).(mongoAdapter.MongoAdapter).
			Collection("post"),
	}
}

func (r *postRepository) GetOne(id string) (*PostEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	foundItem := new(PostEntity)
	r.postCollection.FindOne(ctx, &PostEntity{ID: id}).Decode(foundItem)
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
	r.postCollection.FindOne(ctx, &PostEntity{ID: data.ID}).Decode(foundItem)
	if foundItem != nil {
		return nil, fmt.Errorf("duplication error")
	}

	// operation
	_, err := r.postCollection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
