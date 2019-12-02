package password

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

const (
	salt = "*CFVdkgomy#OAhDmfp4%jlzDj%8zCqOo"
	iter = 32
	kenLen = 128
)

func New(password string) string {
	return fmt.Sprintf("%x", pbkdf2.Key([]byte(password), []byte(salt), iter, kenLen, sha256.New))
}