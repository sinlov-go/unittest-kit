package unittest_random_kit

import (
	"math/rand"
	"time"
)

const randomStrLetterCnt = 62

var randomStrLetters = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// RandomStr
// new random string by cnt, seed by time.Now().UnixNano()
// and is 0-9 A-Z a-z
func RandomStr(cnt uint) string {
	result := make([]byte, cnt)
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range result {
		index := rs.Intn(randomStrLetterCnt)
		result[i] = randomStrLetters[index]
	}
	return string(result)
}

// RandomInt
// new random int by max, seed by time.Now().UnixNano()
func RandomInt(max int) int {
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rs.Intn(max)
}
