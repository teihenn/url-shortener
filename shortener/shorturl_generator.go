package shortener

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/itchyny/base58-go"
)

func sha2560f(input string) []byte {
	algorighm := sha256.New()
	algorighm.Write([]byte(input))
	return algorighm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
