package unittest_file_kit

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// PathExists
//
//	path exists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// PathExistsFast
//
//	path exists fast
func PathExistsFast(path string) bool {
	exists, _ := PathExists(path)
	return exists
}

// PathIsDir
//
//	path is dir
func PathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// RmDir
//
//	remove dir by path
//
//nolint:golint,unused
func RmDir(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}
	exists, err := PathExists(path)
	if err != nil {
		return err
	}
	if exists {
		return os.RemoveAll(path)
	}
	return fmt.Errorf("remove dir not exist path: %s , use force can cover this err", path)
}

// Mkdir
//
// will use FetchDefaultFolderFileMode()
func Mkdir(path string) error {
	err := os.MkdirAll(path, FetchDefaultFolderFileMode())
	if err != nil {
		return fmt.Errorf("fail MkdirAll at path: %s , err: %v", path, err)
	}
	return nil
}

// FetchDefaultFolderFileMode
// if in windows, will return os.FileMode(0o766), windows not support umask
//
//	use umask to get folder file mode
//
// if not windows, will use umask, will return os.FileMode(0o777) - umask
// not support umask will use os.FileMode(0o777)
func FetchDefaultFolderFileMode() fs.FileMode {
	switch runtime.GOOS {
	case "windows":
		return os.FileMode(0o766)
	default:
		umaskCode, err := getUmask()
		if err != nil {
			return os.FileMode(0o777)
		}
		if len(umaskCode) > 3 {
			umaskCode = umaskCode[len(umaskCode)-3:]
		}
		umaskOct, errParseUmask := strconv.ParseInt(umaskCode, 8, 64)
		if errParseUmask != nil {
			return os.FileMode(0o777)
		}
		defaultFOlderCode := 0o777
		// nolint: unconvert
		nowOct := int(defaultFOlderCode) - int(umaskOct)
		return os.FileMode(nowOct)
	}
}

func getUmask() (string, error) {
	cmd := exec.Command("sh", "-c", "umask")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(&out)
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}

	return "", fmt.Errorf("no output from umask command")
}
