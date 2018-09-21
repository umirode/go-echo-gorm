package response

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type jsonContextMock struct {
	mock.Mock
}

func (m *jsonContextMock) JSON(code int, i interface{}) error {
	args := m.Called(code, i)

	return args.Error(0)
}

func TestGetResponseJson(t *testing.T) {
	assert.NotEmpty(t, GetResponseJson("success", "test", "test"))
}

func TestSendResponseJson(t *testing.T) {
	context := &jsonContextMock{}
	context.On("JSON", 200, mock.Anything).Return(nil)

	assert.Empty(t, SendResponseJson(context, "success", "", nil))
}
