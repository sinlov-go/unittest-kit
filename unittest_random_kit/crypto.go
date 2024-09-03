package unittest_random_kit

import (
	"crypto/rand"
	"encoding/base64"
)

// RandomCryptoStr
// generates a random string of length using crypto/rand
// support parallel
func RandomCryptoStr(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

// RandomCryptoStrFast
// generates a random string of length using crypto/rand
// if error, use RandomStr instead
func RandomCryptoStrFast(length int) string {
	res, err := RandomCryptoStr(length)
	if err != nil {
		res = RandomStr(uint(length))
	}
	return res
}
