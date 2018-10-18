package controllers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type getParamContextMock struct {
	mock.Mock
}

func (m *getParamContextMock) Param(name string) string {
	args := m.Called(name)

	return args.String(0)
}

func TestBaseController_getParam(t *testing.T) {
	controller := &BaseController{}

	types := [...]string{
		"uint",
		"string",
		"int",
	}

	context := &getParamContextMock{}
	for _, valueType := range types {
		context.On("Param", mock.Anything).Return("222")

		value, err := controller.getParam(context, "test", valueType)

		assert.NotEmpty(t, value)
		assert.NoError(t, err)
	}
}

func TestBaseController_getParam_EmptyError(t *testing.T) {
	controller := &BaseController{}

	types := [...]string{
		"uint",
		"string",
		"int",
	}

	context := &getParamContextMock{}
	for _, valueType := range types {
		context.On("Param", mock.Anything).Return("")

		value, err := controller.getParam(context, "test", valueType)

		assert.Empty(t, value)
		assert.Error(t, err)
	}
}
func TestBaseController_getParam_UndefinedTypeError(t *testing.T) {
	controller := &BaseController{}

	context := &getParamContextMock{}
	context.On("Param", mock.Anything).Return("222")

	value, err := controller.getParam(context, "test", "test")

	assert.Empty(t, value)
	assert.Error(t, err)
}
