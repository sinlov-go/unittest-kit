package unittest_file_kit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

type TestGoldenKitFunc interface {
	GetTestBaseFolderFullPath() string
	GetTestDataFolderFullPath() string
	GetOrCreateTestDataFolderFullPath() (string, error)
	GetOrCreateTestDataFullPath(elem ...string) (string, error)
	GoldenDataSaveFast(t *testing.T, data interface{}, extraName string) error
	GoldenDataSave(t *testing.T, data []byte, extraName string, fileMod fs.FileMode) error
	GoldenDataReadAsByte(t *testing.T, extraName string) ([]byte, error)
	GoldenDataReadAsType(t *testing.T, extraName string, v interface{}) error
}

type TestGoldenKit struct {
	testBaseFolderPath           string
	currentTestDataFolderAbsPath string
	TestGoldenKitFunc            TestGoldenKitFunc
}

func (g *TestGoldenKit) GetTestBaseFolderFullPath() string {
	return g.testBaseFolderPath
}

func (g *TestGoldenKit) GetTestDataFolderFullPath() string {
	return g.currentTestDataFolderAbsPath
}

// GetOrCreateTestDataFolderFullPath
//
//	will create `testdata` folder under this package
func (g *TestGoldenKit) GetOrCreateTestDataFolderFullPath() (string, error) {
	if g.currentTestDataFolderAbsPath == "" {
		return "", fmt.Errorf("currentTestDataFolderAbsPath is empty")
	}
	if !PathExistsFast(g.currentTestDataFolderAbsPath) {
		err := Mkdir(g.currentTestDataFolderAbsPath)
		if err != nil {
			g.currentTestDataFolderAbsPath = ""
			return "", err
		}
	}
	return g.currentTestDataFolderAbsPath, nil
}

func (g *TestGoldenKit) GetOrCreateTestDataFullPath(elem ...string) (string, error) {
	basicTestDataPath, errTestDataPath := g.GetOrCreateTestDataFolderFullPath()
	if errTestDataPath != nil {
		return "", errTestDataPath
	}
	fullPath := filepath.Join(basicTestDataPath, elem[0])
	if len(elem) > 1 {
		for i := 1; i < len(elem); i++ {
			fullPath = filepath.Join(fullPath, elem[i])
		}
	}
	baseDir := filepath.Dir(fullPath)
	if !PathExistsFast(baseDir) {
		errMkdir := Mkdir(baseDir)
		if errMkdir != nil {
			return fullPath, errMkdir
		}
	}
	if !PathIsDir(baseDir) {
		return "", fmt.Errorf("GetOrCreateTestDataFullPath exist file, and can not create dir at path: %s", baseDir)
	}
	return fullPath, nil
}

// GoldenDataSaveFast
//
//	save data to golden file
//	style as: "TestFuncName/extraName.golden"
func (g *TestGoldenKit) GoldenDataSaveFast(t *testing.T, data interface{}, extraName string) error {
	marshal, errJson := json.Marshal(data)
	if errJson != nil {
		t.Fatal(errJson)
	}
	return g.GoldenDataSave(t, marshal, extraName, os.FileMode(0o666))
}

// GoldenDataSave
//
//	save data to golden file
//	style as: "TestFuncName/extraName.golden"
func (g *TestGoldenKit) GoldenDataSave(t *testing.T, data []byte, extraName string, fileMod fs.FileMode) error {
	testDataFolderFullPath, err := g.GetOrCreateTestDataFolderFullPath()
	if err != nil {
		return fmt.Errorf("try goldenDataSave err: %v", err)
	}
	testDataFolder := filepath.Join(testDataFolderFullPath, t.Name())
	if !PathExistsFast(testDataFolder) {
		errMk := Mkdir(testDataFolder)
		if errMk != nil {
			t.Fatal(errMk)
		}
	}
	savePath := filepath.Join(testDataFolderFullPath, t.Name(), fmt.Sprintf("%s.golden", extraName))
	var str bytes.Buffer
	err = json.Indent(&str, data, "", "  ")
	if err != nil {
		return err
	}
	err = WriteFileByByte(savePath, str.Bytes(), fileMod, true)
	if err != nil {
		return fmt.Errorf("try goldenDataSave at path: %s err: %v", savePath, err)
	}
	return nil
}

// GoldenDataReadAsByte
//
//	read golden file as byte
//	style as: "TestFuncName/extraName.golden"
func (g *TestGoldenKit) GoldenDataReadAsByte(t *testing.T, extraName string) ([]byte, error) {
	savePath := filepath.Join(g.currentTestDataFolderAbsPath, t.Name(), fmt.Sprintf("%s.golden", extraName))
	if !PathExistsFast(savePath) {
		return nil, fmt.Errorf("goldenDataReadAsByte not exist: %s", savePath)
	}
	fileAsByte, err := ReadFileAsByte(savePath)
	if err != nil {
		return nil, fmt.Errorf("try goldenDataReadAsByte err: %v", err)
	}
	return fileAsByte, nil
}

// GoldenDataReadAsType
//
//	read golden file as type
//	style as: "TestFuncName/extraName.golden"
func (g *TestGoldenKit) GoldenDataReadAsType(t *testing.T, extraName string, v interface{}) error {
	readAsByte, err := g.GoldenDataReadAsByte(t, extraName)
	if err != nil {
		t.Fatal(err)
	}
	return json.Unmarshal(readAsByte, v)
}

func NewTestGoldenKit(testFolderPath string) *TestGoldenKit {
	if !PathExistsFast(testFolderPath) {
		panic(fmt.Errorf("NewTestGoldenKit testFolderPath not exist: %s", testFolderPath))
	}

	testDataPath := filepath.Join(testFolderPath, "testdata")

	goldenKit := TestGoldenKit{
		testBaseFolderPath:           testFolderPath,
		currentTestDataFolderAbsPath: testDataPath,
	}
	return &goldenKit
}
