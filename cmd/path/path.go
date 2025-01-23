package path

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type FileInfo struct {
	Path string
	Name string
}

// IsExist returns true if the path
// provided points to a currently existing file (or directory)
func IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateFolders Creates the folders described in
// the provided path
func CreateFolders(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// CreateFile Creates the file pointed by
// the provided path
func CreateFile(path string) error {
	_, err := os.Create(path)
	return err
}

// EnsurePresent checks if the directory path
// points to an existing directory, if not, it creates
// all necessary folders. After that, it checks if the
// given file exists in the last directory of directoryPath.
// If not, the file is created as well
func EnsurePresent(directoryPath, filename string) error {
	if !IsExist(directoryPath) {
		err := CreateFolders(directoryPath)
		if err != nil {
			return fmt.Errorf("error when creating folders %s: %v", directoryPath, err)
		}
	}
	fullPath := path.Join(directoryPath, filename)
	if !IsExist(fullPath) {
		err := CreateFile(fullPath)
		if err != nil {
			return fmt.Errorf("error when creating file %s: %v", fullPath, err)
		}
	}
	return nil
}

// ListAllFiles returns a slice of all files
// contained in directoryPath, recursively checking subdirectories
func ListAllFiles(directoryPath string) ([]FileInfo, error) {
	files := make([]FileInfo, 0)
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			infoStruct := FileInfo{
				Path: path,
				Name: info.Name(),
			}
			files = append(files, infoStruct)
		}
		return nil
	})
	if err != nil {
		return nil, err
	} else {
		return files, nil
	}
}
