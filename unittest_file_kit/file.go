package unittest_file_kit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// ReadFileAsByte
//
//	read file as byte array
func ReadFileAsByte(path string) ([]byte, error) {
	exists, err := PathExists(path)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("path not exist %v", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("path: %s read err: %v", path, err)
	}

	return data, nil
}

// ReadFileAsJson
//
//	read file as json
func ReadFileAsJson(path string, v interface{}) error {
	fileAsByte, errRead := ReadFileAsByte(path)
	if errRead != nil {
		return fmt.Errorf("path: %s , read file as err: %v", path, errRead)
	}
	err := json.Unmarshal(fileAsByte, v)
	if err != nil {
		return fmt.Errorf("path: %s , read file as json err: %v", path, err)
	}
	return nil
}

// WriteFileByByte
//
//	write bytes to file
//	path most use Abs Path
//	data []byte
//	fileMod os.FileMode(0666) or os.FileMode(0644)
//	coverage true will coverage old
func WriteFileByByte(path string, data []byte, fileMod fs.FileMode, coverage bool) error {
	if !coverage {
		exists, err := PathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	parentPath := filepath.Dir(path)
	if !PathExistsFast(parentPath) {
		err := os.MkdirAll(parentPath, fileMod)
		if err != nil {
			return fmt.Errorf("can not WriteFileByByte at new dir at mode: %v , at parent path: %v", fileMod, parentPath)
		}
	}
	err := os.WriteFile(path, data, fileMod)
	if err != nil {
		return fmt.Errorf("write data at path: %v, err: %v", path, err)
	}
	return nil
}

// WriteFileAsJson write json file
//
//	path most use Abs Path
//	v data
//	fileMod os.FileMode(0666) or os.FileMode(0644)
//	coverage true will coverage old
//	beauty will format json when write
func WriteFileAsJson(path string, v interface{}, fileMod fs.FileMode, coverage, beauty bool) error {
	if !coverage {
		exists, err := PathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if beauty {
		var str bytes.Buffer
		err := json.Indent(&str, data, "", "  ")
		if err != nil {
			return err
		}
		return WriteFileByByte(path, str.Bytes(), fileMod, coverage)
	}
	return WriteFileByByte(path, data, fileMod, coverage)
}

// WriteFileAsJsonBeauty
//
//	write json file as 0766 and beauty
func WriteFileAsJsonBeauty(path string, v interface{}, coverage bool) error {
	return WriteFileAsJson(path, v, os.FileMode(0766), coverage, true)
}
