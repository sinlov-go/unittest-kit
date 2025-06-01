package unittest_random_kit

import (
	"math/rand"
	"time"
)

const randomStrLetterCnt = 62

var randomStrLetters = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// RandomStr
// support parallel
// new random string by cnt, seed by time.Now().UnixNano()
// and string only with [0-9A-Za-z]
func RandomStr(cnt uint) string {
	mu.Lock()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	defer mu.Unlock()

	result := make([]byte, cnt)
	for i := range result {
		index := rng.Intn(randomStrLetterCnt)
		result[i] = randomStrLetters[index]
	}
	return string(result)
}

// RandomInt
// support parallel
// new random int by max, seed by time.Now().UnixNano()
// nolint: predeclared
func RandomInt(max int) int {
	mu.Lock()
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	defer mu.Unlock()
	return rs.Intn(max)
}
