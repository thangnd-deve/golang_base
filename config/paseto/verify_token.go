package paseto

import "github.com/o1egl/paseto/v2"

func VerifyToken(key string, token string) (string, error) {
	keyByte := []byte(key)
	var payload string

	// Decrypt and verify the PASETO token
	err := paseto.NewV2().Decrypt(token, keyByte, &payload, nil)
	if err != nil {
		return "", err
	}

	return payload, nil
}
