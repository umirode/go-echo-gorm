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

        assert.NotEmpty(t, controller.getParam(context, "test", valueType))
    }
}
