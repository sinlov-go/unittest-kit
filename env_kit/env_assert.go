package env_kit

import (
	"os"
	"testing"
)

// MustHasEnvSetByArray
// unit test helper env settings
// if not found env in envs will return true
// nolint: thelper
func MustHasEnvSetByArray(t *testing.T, envs []string) bool {
	for _, item := range envs {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env key: %s, full envs need set: %v", item, envs)
			return true
		}
	}
	return false
}
