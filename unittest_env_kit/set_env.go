package unittest_env_kit

import (
	"os"
	"strconv"
	"testing"
)

// SetEnvStr
//
//	set env by key and val
func SetEnvStr(t *testing.T, key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		t.Fatalf("set env key [%v] string err: %v", key, err)
	}
}

// SetEnvBool
//
//	set env by key and val
//
//nolint:golint,unused
func SetEnvBool(t *testing.T, key string, val bool) {
	var err error
	if val {
		err = os.Setenv(key, "true")
	} else {
		err = os.Setenv(key, "false")
	}
	if err != nil {
		t.Fatalf("set env key [%v] bool err: %v", key, err)
	}
}

// SetEnvU64
//
//	set env by key and val
//
//nolint:golint,unused
func SetEnvU64(t *testing.T, key string, val uint64) {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		t.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
}

// SetEnvInt64
//
//	set env by key and val
//
//nolint:golint,unused
func SetEnvInt64(t *testing.T, key string, val int64) {
	err := os.Setenv(key, strconv.FormatInt(val, 10))
	if err != nil {
		t.Fatalf("set env key [%v] int64 err: %v", key, err)
	}
}
