package fileaccessor

import (
	"errors"
	"io/ioutil"
	"os"
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

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm;
// otherwise WriteFile truncates it before writing.
func (fileAccessor SystemFileAccessor) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
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

// WriteFile writes data to a virtual file named by filename. Mostly ignores perm.
func (fileAccessor MemFileAccessor) WriteFile(filename string, data []byte, perm os.FileMode) error {
	if perm == os.ModeAppend {
		return errors.New("MemFileAccessor Append operation currently not supported.")
	}

	fileAccessor.files[filename] = string(data)
	return nil
}
