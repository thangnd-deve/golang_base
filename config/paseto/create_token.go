package paseto

import (
	"fmt"
	"github.com/o1egl/paseto/v2"
	"time"
)

func CreatePasetoToken(key string, userId int) (string, error) {
	keyByte := []byte(key)
	now := time.Now()
	payload := fmt.Sprintf(`{"userId": "%d", "exp": %d}`, userId, now.Unix()+3600)

	// Create a new Paseto token with a version and purpose

	token, err := paseto.NewV2().Encrypt(keyByte, []byte(payload), nil)
	if err != nil {
		return "", err
	}

	return token, nil
}
