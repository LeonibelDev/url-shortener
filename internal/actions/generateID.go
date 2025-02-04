package actions

import (
	"crypto/sha256"
	"math/rand"
)

func GenerateID() (string, error) {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	id := make([]byte, 6)

	for i := range id {
		id[i] = characters[rand.Intn(len(characters))]
	}

	hash := sha256.New()
	hash.Write(id)

	return string(id), nil
}
