package file

import (
	"fmt"
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

// CheckNotExist check if the file exists
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission check if the file has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkdir create a directory if it does not exist
func IsNotExistMkdir(src string) error {
	if notExist := CheckNotExist(src); notExist {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir creates a directory
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	return err
}

// Open open a file according to a specific model
func Open(filename string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("[pkg.file.MustOpen] os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("[pkg.file.MustOpen] file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkdir(src)
	if err != nil {
		return nil, fmt.Errorf("[pkg.file.MustOpen] file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("[pkg.file.MustOpen] Fail to OpenFile :%v", err)
	}

	return f, nil
}
