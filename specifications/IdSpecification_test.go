package specifications

import (
	"github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIdSpecification_GetForDatabase(t *testing.T) {
	db, _ := gorm.Open(go_mocket.DRIVER_NAME, "")

	specification := IdSpecification{}

	assert.NotEmpty(t, specification.GetForDatabase(db))
}
