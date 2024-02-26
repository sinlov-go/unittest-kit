package env_kit

import (
	"log"
	"os"
	"strconv"
	"strings"
)

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

func SetEnvStr(key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		log.Fatalf("set env key [%v] string err: %v", key, err)
	}
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

func SetEnvStringSlice(key string, val []string) {
	err := os.Setenv(key, strings.Join(val, ","))
	if err != nil {
		log.Fatalf("set env key [%v] string slice err: %v", key, err)
	}
}

// FetchOsEnvStringSlice
//
//	fetch os env split by `,` and trim space
//	return not found will return []string(nil).
func FetchOsEnvStringSlice(key string) []string {
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

// FetchOsEnvArray
//
//	fetch os env split by `,` and trim space
//	return not found will return []string(nil).
//
// Deprecated: use FetchOsEnvStringSlice instead
func FetchOsEnvArray(key string) []string {
	return FetchOsEnvStringSlice(key)
}

func SetEnvInt(key string, val int) {
	err := os.Setenv(key, strconv.Itoa(val))
	if err != nil {
		log.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
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

func SetEnvIntSlice(key string, val []int) {
	var strVal []string
	for _, item := range val {
		strVal = append(strVal, strconv.Itoa(item))
	}
	err := os.Setenv(key, strings.Join(strVal, ","))
	if err != nil {
		log.Fatalf("set env key [%v] int slice err: %v", key, err)
	}
}

// FetchOsEnvIntSlice
//
//	fetch os env split by `,` and trim space
//	return not found will return []int(nil).
func FetchOsEnvIntSlice(key string) []int {
	if os.Getenv(key) == "" {
		return nil
	}
	envValue := os.Getenv(key)
	splitVal := strings.Split(envValue, ",")
	if len(splitVal) == 0 {
		return nil
	}
	var outVal []int
	for _, item := range splitVal {
		outNum, err := strconv.Atoi(strings.TrimSpace(item))
		if err != nil {
			return nil
		}
		outVal = append(outVal, outNum)
	}
	return outVal
}

func SetEnvInt64(key string, val int64) {
	err := os.Setenv(key, strconv.FormatInt(val, 10))
	if err != nil {
		log.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
}

// FetchOsEnvInt64
//
//	fetch os env by key.
//	return not found will return devValue.
//	if not parse to uint, return devValue
func FetchOsEnvInt64(key string, devValue int64) int64 {
	if os.Getenv(key) == "" {
		return devValue
	}
	parsed, err := strconv.ParseInt(os.Getenv(key), 0, 64)
	if err != nil {
		return 0
	}
	return parsed
}

func SetEnvInt64Slice(key string, val []int64) {
	var strVal []string
	for _, item := range val {
		strVal = append(strVal, strconv.FormatInt(item, 10))
	}
	err := os.Setenv(key, strings.Join(strVal, ","))
	if err != nil {
		log.Fatalf("set env key [%v] int64 slice err: %v", key, err)
	}
}

// FetchOsEnvInt64Slice
//
//	fetch os env split by `,` and trim space
//	return not found will return []uint(nil).
func FetchOsEnvInt64Slice(key string) []int64 {
	if os.Getenv(key) == "" {
		return nil
	}
	envValue := os.Getenv(key)
	splitVal := strings.Split(envValue, ",")
	if len(splitVal) == 0 {
		return nil
	}
	var outVal []int64
	for _, item := range splitVal {
		outNum, err := strconv.ParseInt(strings.TrimSpace(item), 0, 64)
		if err != nil {
			return nil
		}
		outVal = append(outVal, outNum)
	}
	return outVal
}

func SetEnvUint(key string, val uint) {
	err := os.Setenv(key, strconv.FormatUint(uint64(val), 10))
	if err != nil {
		log.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
}

// FetchOsEnvUint
//
//	fetch os env by key.
//	return not found will return devValue.
//	if not parse to uint, return devValue
func FetchOsEnvUint(key string, devValue uint) uint {
	if os.Getenv(key) == "" {
		return devValue
	}
	parsed, err := strconv.ParseUint(os.Getenv(key), 0, 64)
	if err != nil {
		return 0
	}
	return uint(parsed)
}

func SetEnvUintSlice(key string, val []uint) {
	var strVal []string
	for _, item := range val {
		strVal = append(strVal, strconv.FormatUint(uint64(item), 10))
	}
	err := os.Setenv(key, strings.Join(strVal, ","))
	if err != nil {
		log.Fatalf("set env key [%v] uint slice err: %v", key, err)
	}
}

// FetchOsEnvUintSlice
//
//	fetch os env split by `,` and trim space
//	return not found will return []uint(nil).
func FetchOsEnvUintSlice(key string) []uint {
	if os.Getenv(key) == "" {
		return nil
	}
	envValue := os.Getenv(key)
	splitVal := strings.Split(envValue, ",")
	if len(splitVal) == 0 {
		return nil
	}
	var outVal []uint
	for _, item := range splitVal {
		outNum, err := strconv.ParseUint(strings.TrimSpace(item), 0, 64)
		if err != nil {
			return nil
		}
		outVal = append(outVal, uint(outNum))
	}
	return outVal
}

func SetEnvUint64(key string, val uint64) {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		log.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
}

// Deprecated: use SetEnvUint64 instead
func SetEnvU64(key string, val uint64) {
	SetEnvUint64(key, val)
}

// FetchOsEnvUint64
//
//	fetch os env by key.
//	return not found will return devValue.
//	if not parse to uint64, return devValue
func FetchOsEnvUint64(key string, devValue uint64) uint64 {
	if os.Getenv(key) == "" {
		return devValue
	}
	parsed, err := strconv.ParseUint(os.Getenv(key), 0, 64)
	if err != nil {
		return 0
	}
	return parsed
}

func SetEnvUint64Slice(key string, val []uint64) {
	var strVal []string
	for _, item := range val {
		strVal = append(strVal, strconv.FormatUint(item, 10))
	}
	err := os.Setenv(key, strings.Join(strVal, ","))
	if err != nil {
		log.Fatalf("set env key [%v] uint64 slice err: %v", key, err)
	}
}

// FetchOsEnvUint64Slice
//
//	fetch os env split by `,` and trim space
//	return not found will return []uint(nil).
func FetchOsEnvUint64Slice(key string) []uint64 {
	if os.Getenv(key) == "" {
		return nil
	}
	envValue := os.Getenv(key)
	splitVal := strings.Split(envValue, ",")
	if len(splitVal) == 0 {
		return nil
	}
	var outVal []uint64
	for _, item := range splitVal {
		outNum, err := strconv.ParseUint(strings.TrimSpace(item), 0, 64)
		if err != nil {
			return nil
		}
		outVal = append(outVal, outNum)
	}
	return outVal
}
