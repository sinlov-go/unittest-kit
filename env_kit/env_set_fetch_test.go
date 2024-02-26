package env_kit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetFetchEnvBool(t *testing.T) {
	const key = "FOO"
	const val = true
	SetEnvBool(key, val)

	gotResult := FetchOsEnvBool(key, false)
	assert.Equal(t, val, gotResult)
}

func TestSetFetchEnvStr(t *testing.T) {
	const key = "FOO"
	const val = "bar"
	SetEnvStr(key, val)

	gotResult := FetchOsEnvStr(key, "")
	assert.Equal(t, val, gotResult)
}

func TestSetFetchEnvStringSlice(t *testing.T) {
	const key = "FOO"
	val := []string{"bar", "foo"}
	SetEnvStringSlice(key, val)

	gotResult := FetchOsEnvStringSlice(key)
	assert.Equal(t, val, gotResult)
}

func TestSetFetchEnvInt(t *testing.T) {
	const key = "FOO"
	const val = 123
	SetEnvInt(key, val)

	gotResult := FetchOsEnvInt(key, 0)
	assert.Equal(t, val, gotResult)
}

func TestSetFetchEnvIntSlice(t *testing.T) {
	const key = "FOO"
	val := []int{123, 456}
	SetEnvIntSlice(key, val)

	gotResult := FetchOsEnvIntSlice(key)
	assert.Equal(t, val, gotResult)
}

func TestSetFetchEnvInt64(t *testing.T) {
	const key = "FOO"
	const val = int64(123)
	SetEnvInt64(key, val)

	gotResult := FetchOsEnvInt64(key, 0)
	assert.Equal(t, val, gotResult)
}

func TestSetFetchEnvInt64Slice(t *testing.T) {
	const key = "FOO"
	val := []int64{123, 456}
	SetEnvInt64Slice(key, val)

	gotResult := FetchOsEnvInt64Slice(key)
	assert.Equal(t, val, gotResult)
}

func TestSetFetchEnvUint64(t *testing.T) {
	const key = "FOO"
	const val = uint64(123)
	SetEnvUint64(key, val)

	gotResult := FetchOsEnvUint64(key, 0)
	assert.Equal(t, val, gotResult)
}

func TestSetFetchEnvUint64Slice(t *testing.T) {
	const key = "FOO"
	val := []uint64{123, 456}
	SetEnvUint64Slice(key, val)

	gotResult := FetchOsEnvUint64Slice(key)
	assert.Equal(t, val, gotResult)
}

func TestSetFetchEnvUint(t *testing.T) {
	const key = "FOO"
	const val = uint(123)
	SetEnvUint(key, val)

	gotResult := FetchOsEnvUint(key, 0)
	assert.Equal(t, val, gotResult)
}

func TestSetFetchEnvUintSlice(t *testing.T) {
	const key = "FOO"
	val := []uint{123, 456}
	SetEnvUintSlice(key, val)

	gotResult := FetchOsEnvUintSlice(key)
	assert.Equal(t, val, gotResult)
}
