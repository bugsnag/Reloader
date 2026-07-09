package crypto

import (
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

// GenerateSHA generates a SHA-256 digest from the input string.
func GenerateSHA(data string) string {
	hasher := sha256.New()
	_, err := io.WriteString(hasher, data)
	if err != nil {
		logrus.Errorf("Unable to write data in hash writer %v", err)
	}
	sha := hasher.Sum(nil)
	return fmt.Sprintf("%x", sha)
}
