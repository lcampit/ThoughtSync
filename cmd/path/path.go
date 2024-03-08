package path

import (
	"fmt"
	"os"
	"path"
)

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
