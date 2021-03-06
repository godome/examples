// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	component "github.com/godome/godome/pkg/component"
	config "github.com/godome/godome/pkg/config"

	logger "github.com/godome/godome/pkg/logger"

	mock "github.com/stretchr/testify/mock"

	mongoPlugin "github.com/godome/plugins/pkg/mongo-plugin"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MongoAdapter is an autogenerated mock type for the MongoAdapter type
type MongoAdapter struct {
	mock.Mock
}

// Collection provides a mock function with given fields: name, opts
func (_m *MongoAdapter) Collection(name string, opts ...*options.CollectionOptions) mongoPlugin.MongoCollection {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongoPlugin.MongoCollection
	if rf, ok := ret.Get(0).(func(string, ...*options.CollectionOptions) mongoPlugin.MongoCollection); ok {
		r0 = rf(name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoPlugin.MongoCollection)
		}
	}

	return r0
}

// Config provides a mock function with given fields:
func (_m *MongoAdapter) Config() config.Config {
	ret := _m.Called()

	var r0 config.Config
	if rf, ok := ret.Get(0).(func() config.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Config)
		}
	}

	return r0
}

// Disconnect provides a mock function with given fields:
func (_m *MongoAdapter) Disconnect() {
	_m.Called()
}

// Logger provides a mock function with given fields:
func (_m *MongoAdapter) Logger() logger.Logger {
	ret := _m.Called()

	var r0 logger.Logger
	if rf, ok := ret.Get(0).(func() logger.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.Logger)
		}
	}

	return r0
}

// Metadata provides a mock function with given fields:
func (_m *MongoAdapter) Metadata() component.Metadata {
	ret := _m.Called()

	var r0 component.Metadata
	if rf, ok := ret.Get(0).(func() component.Metadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(component.Metadata)
		}
	}

	return r0
}
