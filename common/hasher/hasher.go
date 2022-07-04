package hasher

import (
	"crypto/rand"
	"github.com/matthewhartstonge/argon2"
	"io"
)

func Hash(str string, salt []byte) ([]byte, error) {
	argon := argon2.DefaultConfig()

	r, err := argon.Hash([]byte(str), salt)

	return r.Encode(), err
}

func VerifyHash(str string, hash string) (bool, error) {
	return argon2.VerifyEncoded([]byte(str), []byte(hash))
}

func GenerateSaltBytes(saltLen uint) ([]byte, error) {
	salt := make([]byte, saltLen)
	_, err := io.ReadFull(rand.Reader, salt)

	return salt, err
}
