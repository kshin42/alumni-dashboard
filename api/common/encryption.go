package common

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/argon2"
)

type gParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
	hash        string
}

func HashPassword(pw string) (string, error) {
	params := &gParams{
		memory:      64 * 128,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	salt, err := generateRandomBytes(params.saltLength)
	if err != nil {
		log.Error().Msg("Error generating password salt")
		return "", err
	}

	hash := argon2.IDKey([]byte(pw), salt, params.iterations, params.memory, params.parallelism, params.keyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("%s$%s", b64Salt, b64Hash)
	return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b) //uses the crypto rand so this is actually a secure random number
	if err != nil {
		return nil, err
	}

	return b, nil
}
