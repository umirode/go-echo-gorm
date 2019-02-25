package Config

import (
	"sync"

	"github.com/umirode/go-rest/Config/Helper"
)

type FirebaseConfig struct {
	CloudMessagingKey string
}

var firebaseConfigOnce sync.Once
var firebaseConfig *FirebaseConfig

func GetFirebaseConfig() *FirebaseConfig {
	firebaseConfigOnce.Do(func() {
		firebaseConfig = &FirebaseConfig{
			CloudMessagingKey: Helper.GetEnv("FIREBASE_CLOUD_MESSAGING_KEY", "string").(string),
		}
	})

	return firebaseConfig
}
