package unittest_random_kit

import (
	"sync"
)

var (
	mu sync.Mutex
)

func init() {
	// initialize the random number generator
	mu.Lock()
	defer mu.Unlock()
}
