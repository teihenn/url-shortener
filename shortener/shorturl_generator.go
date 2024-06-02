package shortener

import (
	"crypto/sha256"
)

func sha2560f(input string) []byte {
	algorighm := sha256.New()
	algorighm.Write([]byte(input))
	return algorighm.Sum(nil)
}
