package helloworld_test

import (
	"testing"

	helloworld "github.com/godome/examples/pkg/hello-world-server/hello-world"
	"github.com/stretchr/testify/assert"
)

type serviceFixture struct {
	service helloworld.HelloworldService
}

func newServiceFixture() *serviceFixture {
	f := new(serviceFixture)
	// init helloworld module
	module := helloworld.NewHelloworldModule()
	f.service = module.GetProvider(helloworld.ServiceType).(helloworld.HelloworldService)
	return f
}

func TestHelloworld(t *testing.T) {
	f := newServiceFixture()
	assert := assert.New(t)

	// When try to get /test-name
	// It shoud say hello with the given name
	greet := f.service.SayHello("john doe")
	assert.Equal(greet, "hello john doe")
}
