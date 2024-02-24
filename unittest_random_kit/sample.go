package unittest_random_kit

import (
	"math/rand"
	"time"
)

const randomStrLetterCnt = 62

var randomStrLetters = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// RandomStr
//
//	new random string by cnt
func RandomStr(cnt uint) string {
	result := make([]byte, cnt)
	rs := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range result {
		index := rs.Intn(randomStrLetterCnt)
		result[i] = randomStrLetters[index]
	}
	return string(result)
}

// RandomInt
//
//	new random int by max
func RandomInt(max int) int {
	rs := rand.New(rand.NewSource(time.Now().Unix()))
	return rs.Intn(max)
}
