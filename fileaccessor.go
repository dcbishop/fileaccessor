package fileaccessor

import (
	"errors"
	"io/ioutil"
)

// FileAccessor loads files
type FileAccessor interface {
	ReadFile(path string)
}

// SystemFileAccessor loads files from the disk
type SystemFileAccessor struct{}

// ReadFile load a file in from disk
func (fileAccessor SystemFileAccessor) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// MemFileAccessor load virtual files from memory
type MemFileAccessor struct {
	files map[string]string
}

// ReadFile returns the file as a string
func (fileAccessor MemFileAccessor) ReadFile(filename string) ([]byte, error) {
	data, ok := fileAccessor.files[filename]
	if ok != true {
		return nil, errors.New("File not found.")
	}
	return []byte(data), nil
}
