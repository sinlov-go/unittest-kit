[![ci](https://github.com/sinlov-go/unittest-kit/actions/workflows/ci.yml/badge.svg)](https://github.com/sinlov-go/unittest-kit/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov-go/unittest-kit?label=go.mod)](https://github.com/sinlov-go/unittest-kit)
[![GoDoc](https://godoc.org/github.com/sinlov-go/unittest-kit?status.png)](https://godoc.org/github.com/sinlov-go/unittest-kit)
[![goreportcard](https://goreportcard.com/badge/github.com/sinlov-go/unittest-kit)](https://goreportcard.com/report/github.com/sinlov-go/unittest-kit)

[![GitHub license](https://img.shields.io/github/license/sinlov-go/unittest-kit)](https://github.com/sinlov-go/unittest-kit)
[![codecov](https://codecov.io/gh/sinlov-go/unittest-kit/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov-go/unittest-kit)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/sinlov-go/unittest-kit)](https://github.com/sinlov-go/unittest-kit/tags)
[![GitHub release)](https://img.shields.io/github/v/release/sinlov-go/unittest-kit)](https://github.com/sinlov-go/unittest-kit/releases)

## for what

- this project used to golang unit test case.

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov-go/unittest-kit)](https://github.com/sinlov-go/unittest-kit/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/sinlov-go/unittest-kit.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/sinlov-go/unittest-kit
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov-go/unittest-kit | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## Features

- [x] `env_kit` package
    - `env_kit.FetchOsEnv*` and `env_kit.SetEnv*` for env get or set
    - `env_kit.FindAllEnv4Print`, `env_kit.FindAllEnvByPrefix`, `env_kit.FindAllEnv4PrintAsSortJust` for find print env string
    - `env_kit.MustHasEnvSetByArray` use in unit test env setting check 
- [x] `unittest_file_kit`
    - `unittest_file_kit.TestGoldenKit` for golden data test
- [x] `unittest_random_kit`
    - `unittest_random_kit.RandomStr` for random string
    - `unittest_random_kit.RandomInt` for random int
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## env

- minimum go version: go 1.19
- change `go 1.19`, `^1.19`, `1.19.12-bullseye`, `1.19.12` to new go version

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |
| https://github.com/sebdah/goldie    | v2.5.3  |

- more libs see [go.mod](https://github.com/sinlov-go/unittest-kit/blob/main/go.mod)

## usage

```go
package foo_test

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"path/filepath"
	"runtime"
)

const (
	keyEnvDebug  = "CI_DEBUG"
	keyEnvCiNum  = "CI_NUMBER"
	keyEnvCiKey  = "CI_KEY"
	keyEnvCiKeys = "CI_KEYS"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""
	testGoldenKit      *unittest_file_kit.TestGoldenKit

	envDebug  = false
	envCiNum  = 0
	envCiKey  = ""
	envCiKeys []string
)

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()

	envDebug = env_kit.FetchOsEnvBool(keyEnvDebug, false)
	envCiNum = env_kit.FetchOsEnvInt(keyEnvCiNum, 0)
	envCiKey = env_kit.FetchOsEnvStr(keyEnvCiKey, "")
	envCiKeys = env_kit.FetchOsEnvArray(keyEnvCiKeys)

	testGoldenKit = unittest_file_kit.NewTestGoldenKit(testBaseFolderPath)
}

// test case basic tools start
// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// test case basic tools end
```

- can use env as

```go
package foo_test

import (
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/sinlov-go/unittest-kit/unittest_env_kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvKeys(t *testing.T) {
	// mock EnvKeys
	const keyEnvs = "ENV_KEYS"
	t.Logf("~> mock EnvKeys")

	unittest_env_kit.SetEnvBool(t, keyEnvDebug, true)

	unittest_env_kit.SetEnvInt64(t, keyEnvCiNum, 2)

	unittest_env_kit.SetEnvStr(t, keyEnvCiKey, "foo")

	// do EnvKeys
	t.Logf("~> do EnvKeys")

	// verify EnvKeys

	assert.True(t, env_kit.FetchOsEnvBool(keyEnvDebug, false))
	assert.Equal(t, 2, env_kit.FetchOsEnvInt(keyEnvCiNum, 0))
	assert.Equal(t, "foo", env_kit.FetchOsEnvStr(keyEnvCiKey, ""))
	envArray := env_kit.FetchOsEnvArray(keyEnvs)
	assert.Nil(t, envArray)

	unittest_env_kit.SetEnvStr(t, keyEnvs, "foo, bar,My ")

	envArray = env_kit.FetchOsEnvArray(keyEnvs)

	assert.NotNil(t, envArray)
	assert.Equal(t, "foo", envArray[0])
	assert.Equal(t, "bar", envArray[1])
	assert.Equal(t, "My", envArray[2])

	t.Logf("~> verify EnvKeys: \n%s", env_kit.FindAllEnvByPrefix4Print("CI_"))
	envByPrefix := env_kit.FindAllEnvByPrefix("CI_")
	t.Logf("~> print findAllEnvByPrefix: %v", envByPrefix)
}
```

- then in unit test can use as golden data

```go
package foo_test

func Test_goldenDataSaveFast(t *testing.T) {
	// mock _goldenDataSaveFast
	const extraName = "_goldenDataSaveFast"
	type testData struct {
		Name string
	}
	var fooData = testData{
		Name: "foo",
	}
	t.Logf("~> mock _goldenDataSaveFast")
	err := testGoldenKit.GoldenDataSaveFast(t, fooData, extraName)
	if err != nil {
		t.Fatal(err)
	}
	// do _goldenDataSaveFast
	t.Logf("~> do _goldenDataSaveFast")
	var readData testData
	err = testGoldenKit.GoldenDataReadAsType(t, extraName, &readData)
	if err != nil {
		t.Fatal(err)
	}
	// verify _goldenDataSaveFast
	assert.Equal(t, fooData.Name, readData.Name)
}
```

# dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run at env dev use cmd/main.go
$ make dev
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```
