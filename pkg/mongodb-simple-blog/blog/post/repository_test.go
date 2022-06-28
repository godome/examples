package post_test

import (
	"testing"

	"github.com/godome/examples/pkg/mongodb-simple-blog/blog/post"
	"github.com/godome/examples/pkg/mongodb-simple-blog/mocks"
	"github.com/godome/godome/pkg/component"
	mongoPlugin "github.com/godome/plugins/pkg/mongo-plugin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type repositoryFixture struct {
	repository post.PostRepository
	mocks      struct {
		Adapter     *mocks.MongoAdapter
		MongoMocker *mtest.T
	}
}

func newRepositoryFixture(t *testing.T) *repositoryFixture {
	f := new(repositoryFixture)

	// mocks
	f.mocks.Adapter = &mocks.MongoAdapter{}
	f.mocks.MongoMocker = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	// init module
	f.mocks.Adapter.On("Metadata").Return(component.NewMetadata(mongoPlugin.AdapterName, "adapter"))
	module := post.NewPostModule(f.mocks.Adapter)

	// add providers
	f.repository = module.GetProvider(post.PostRepositoryName).(post.PostRepository)

	return f
}

func TestPostRepository(t *testing.T) {
	assert := assert.New(t)

	// GetOne
	//

	// When item is not found
	//
	f := newRepositoryFixture(t)
	f.mocks.MongoMocker.Run("test", func(mt *mtest.T) {
		f.mocks.Adapter.On("Collection", "post").Return(mt.Coll)
		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1,
				"DBName.CollectionName",
				mtest.FirstBatch,
				nil,
			),
		)
		items, err := f.repository.GetOne("unique-test-id")
		assert.Equal(err.Error(), "post with id [unique-test-id] is not found")
		assert.Nil(items)
	})
	f.mocks.MongoMocker.Close()

	// When item is found
	//
	f = newRepositoryFixture(t)
	f.mocks.MongoMocker.Run("test", func(mt *mtest.T) {
		f.mocks.Adapter.On("Collection", "post").Return(mt.Coll)
		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1,
				"DBName.CollectionName",
				mtest.FirstBatch,
				bson.D{{Key: "id", Value: "unique-test-id"}},
			),
		)
		item, err := f.repository.GetOne("unique-test-id")
		assert.NotNil(item)
		assert.Equal(item.ID, "unique-test-id")
		assert.Nil(err)
	})
	f.mocks.MongoMocker.Close()

	// Create
	//

	// When the new item id is already exist
	//
	f = newRepositoryFixture(t)
	f.mocks.MongoMocker.Run("", func(mt *mtest.T) {
		f.mocks.Adapter.On("Collection", "post").Return(mt.Coll)
		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1,
				"DBName.CollectionName",
				mtest.FirstBatch,
				bson.D{{Key: "id", Value: "unique-test-id"}},
			),
		)
		item, err := f.repository.Create(&post.PostEntity{})
		assert.Nil(item)
		assert.Equal(err.Error(), "duplication error")
	})
	f.mocks.MongoMocker.Close()

	// When mongo has some error on InsertOne
	//
	f = newRepositoryFixture(t)
	f.mocks.MongoMocker.Run("", func(mt *mtest.T) {
		f.mocks.Adapter.On("Collection", "post").Return(mt.Coll)
		mt.AddMockResponses(
			nil, // FindOne mock response
			mtest.CreateWriteErrorsResponse(
				mtest.WriteError{
					Index:   1,
					Code:    11000,
					Message: "some mongo error",
				},
			), // InsertOne mock response
		)
		item, err := f.repository.Create(&post.PostEntity{})
		assert.Nil(item)
		assert.Equal(err.Error(), "write exception: write errors: [some mongo error]")
	})
	f.mocks.MongoMocker.Close()

	// When the new item is created successfully
	//
	f = newRepositoryFixture(t)
	f.mocks.MongoMocker.Run("", func(mt *mtest.T) {
		f.mocks.Adapter.On("Collection", "post").Return(mt.Coll)
		mt.AddMockResponses(
			nil,                           // FindOne mock response
			mtest.CreateSuccessResponse(), // InsertOne mock response
		)
		item, err := f.repository.Create(&post.PostEntity{})

		assert.NoError(err)
		assert.NotZero(item.ID)
	})
	f.mocks.MongoMocker.Close()
}
