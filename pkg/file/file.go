package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// GetSize get file size
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt returns the file name extension used by path
func GetExt(filename string) string {
	return path.Ext(filename)
}

// CheckExist checks src exist or not
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission checks permission problem
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// MkDir creates a directory named src
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	return err
}

// IsNotExistMkdir checks src is not exist and creates a directory named src
func IsNotExistMkdir(src string) error {
	if exist := CheckExist(src); !exist {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// Open opens a file
func Open(filename string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
