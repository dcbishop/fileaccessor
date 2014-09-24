package fileaccessor

import (
	"errors"
	"io/ioutil"
	"os"
)

// FileAccessor loads files
type FileAccessor interface {
	ReadFile(path string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

// FileReader reads files into a byte slice.
type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

// FileWriter writes bytes into a file given by filename.
type FileWriter interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

// LocalStorage loads files from the disk
type LocalStorage struct{}

// ReadFile load a file in from disk
func (fileAccessor LocalStorage) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm;
// otherwise WriteFile truncates it before writing.
func (fileAccessor LocalStorage) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

// Virtual load virtual files from memory
type Virtual struct {
	files map[string][]byte
}

// ReadFile returns the file as a string
func (fileAccessor Virtual) ReadFile(filename string) ([]byte, error) {
	data, ok := fileAccessor.files[filename]
	if ok != true {
		return nil, errors.New("File not found.")
	}
	return []byte(data), nil
}

// WriteFile writes data to a virtual file named by filename. Mostly ignores perm.
func (fileAccessor Virtual) WriteFile(filename string, data []byte, perm os.FileMode) error {
	if perm == os.ModeAppend {
		return errors.New("fileaccessor.Virtual Append operation currently not supported")
	}

	fileAccessor.files[filename] = data
	return nil
}
