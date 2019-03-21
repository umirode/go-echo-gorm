package Config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFirebaseConfig(t *testing.T) {
	err := os.Setenv("FIREBASE_CLOUD_MESSAGING_KEY", "test")
	assert.Nil(t, err)

	config := GetFirebaseConfig()

	assert.NotNil(t, config)
	assert.Equal(t, "test", config.CloudMessagingKey)
}
