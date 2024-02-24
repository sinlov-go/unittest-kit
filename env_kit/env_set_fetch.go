package env_kit

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func SetEnvStr(key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		log.Fatalf("set env key [%v] string err: %v", key, err)
	}
}

func SetEnvBool(key string, val bool) {
	var err error
	if val {
		err = os.Setenv(key, "true")
	} else {
		err = os.Setenv(key, "false")
	}
	if err != nil {
		log.Fatalf("set env key [%v] bool err: %v", key, err)
	}
}

func SetEnvU64(key string, val uint64) {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		log.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
}

// FetchOsEnvBool
//
//	fetch os env by key.
//	if not found will return devValue.
//	return env not same as true (will be lowercase, so TRUE is same)
func FetchOsEnvBool(key string, devValue bool) bool {
	if os.Getenv(key) == "" {
		return devValue
	}
	return strings.ToLower(os.Getenv(key)) == "true"
}

// FetchOsEnvInt
//
//	fetch os env by key.
//	return not found will return devValue.
//	if not parse to int, return devValue
func FetchOsEnvInt(key string, devValue int) int {
	if os.Getenv(key) == "" {
		return devValue
	}
	outNum, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return devValue
	}

	return outNum
}

// FetchOsEnvStr
//
//	fetch os env by key.
//	return not found will return devValue.
func FetchOsEnvStr(key, devValue string) string {
	if os.Getenv(key) == "" {
		return devValue
	}
	return os.Getenv(key)
}

// FetchOsEnvArray
//
//	fetch os env split by `,` and trim space
//	return not found will return []string(nil).
func FetchOsEnvArray(key string) []string {
	var devValueStr []string
	if os.Getenv(key) == "" {
		return devValueStr
	}
	envValue := os.Getenv(key)
	splitVal := strings.Split(envValue, ",")
	if len(splitVal) == 0 {
		return devValueStr
	}
	for _, item := range splitVal {
		devValueStr = append(devValueStr, strings.TrimSpace(item))
	}

	return devValueStr
}
