package post_test

import (
	"testing"

	"github.com/godome/examples/pkg/mongodb-simple-blog/blog/post"
	"github.com/godome/examples/pkg/mongodb-simple-blog/mocks"
	"github.com/godome/godome/pkg/component"
	"github.com/godome/godome/pkg/component/adapter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceFixture struct {
	service post.PostService
	mocks   struct{ *mocks.PostRepository }
}

func newServiceFixture() *serviceFixture {
	f := new(serviceFixture)
	// init module with an "empty" adapter
	module := post.NewPostModule(adapter.NewAdapter("test"))
	// mock repository layer
	f.mocks.PostRepository = &mocks.PostRepository{}
	f.mocks.PostRepository.On("Metadata").Return(component.NewMetadata(post.PostRepositoryName, "provider"))
	module.SetProvider(f.mocks.PostRepository)

	// add providers
	f.service = module.GetProvider(post.PostServiceName).(post.PostService)

	return f
}

func TestPostService(t *testing.T) {
	f := newServiceFixture()
	assert := assert.New(t)
	testId := "test-id"

	// Testing Metadata
	assert.Equal(f.service.Metadata().GetName(), post.PostServiceName)

	// Testing GetPost
	f.mocks.PostRepository.On("GetOne", testId).Return(&post.PostEntity{ID: testId}, nil)
	p, err := f.service.GetPost(testId)
	assert.Nil(err)
	assert.Equal(p.ID, testId)

	// Testing CreatePost
	testName := "test-name"
	testDescription := "test-description"
	f.mocks.PostRepository.On("Create", mock.Anything).Return(&post.PostEntity{testId, testName, testDescription}, nil)
	np, err := f.service.CreatePost(testName, testDescription)
	assert.Nil(err)
	assert.Equal(np.ID, testId)
	assert.Equal(np.Name, testName)
	assert.Equal(np.Description, testDescription)
}
