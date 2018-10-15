package response

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

type jsonContextMock struct {
	mock.Mock
}

func (m *jsonContextMock) JSON(code int, i interface{}) error {
	args := m.Called(code, i)

	return args.Error(0)
}

func TestSendResponseJson(t *testing.T) {
	context := &jsonContextMock{}
	context.On("JSON", 200, mock.Anything).Return(nil)

	assert.Empty(t, SendResponseJson(context, http.StatusOK, nil))
}
