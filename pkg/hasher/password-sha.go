package hasher

import (
	"crypto/sha1"
	"fmt"
)

const salt = "ie498grhfneo30fnf"

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
