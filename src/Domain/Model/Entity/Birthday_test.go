package Entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBirthday(t *testing.T) {
	birthday := &Birthday{
		ID:      1,
		Name:    "test",
		Month:   1,
		Day:     1,
		Year:    2001,
		OwnerID: 1,
	}

	assert.NotNil(t, birthday)
	assert.Equal(t, uint(1), birthday.ID)
	assert.Equal(t, "test", birthday.Name)
	assert.Equal(t, uint(1), birthday.Month)
	assert.Equal(t, uint(1), birthday.Day)
	assert.Equal(t, uint(2001), birthday.Year)
	assert.Equal(t, uint(1), birthday.OwnerID)
}
