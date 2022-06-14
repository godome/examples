package helloworld_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	helloworld "github.com/godome/examples/pkg/hello-world-server/hello-world"
	"github.com/godome/examples/pkg/hello-world-server/mocks"
	"github.com/godome/godome/pkg/component"
	fiberPlugin "github.com/godome/plugins/pkg/fiber-plugin"

	"github.com/stretchr/testify/assert"
)

type handlerFixture struct {
	server fiberPlugin.FiberExposure
	mocks  struct{ *mocks.HelloworldService }
}

func newHandlerFixture() *handlerFixture {
	f := new(handlerFixture)
	// init module
	module := helloworld.NewHelloworldModule()
	// mocks
	f.mocks.HelloworldService = &mocks.HelloworldService{}
	f.mocks.HelloworldService.On("Metadata").Return(component.NewMetadata(helloworld.HelloworldServiceName, "provider"))
	module.SetProvider(f.mocks.HelloworldService)
	// exposure
	f.server = fiberPlugin.NewFiberExposure("", nil).ExposeModule(module)
	return f
}

func TestHandler(t *testing.T) {
	f := newHandlerFixture()
	assert := assert.New(t)

	// When try to get /test-name
	// It should say hello with the given mocked name
	f.mocks.HelloworldService.On("SayHello", "test-name").Return("hello test-name from mocked service")
	res, err := f.server.Test(httptest.NewRequest("GET", "/test-name", nil), -1)
	assert.Equal(err, nil)
	assert.NotEqual(res, nil)
	body, err := ioutil.ReadAll(res.Body)
	assert.Equal(err, nil)
	assert.Equal(string(body), "hello test-name from mocked service")
}
