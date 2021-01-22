package filesystem

import (
	"os"
	"path/filepath"
	"strings"
)

// Exist checks if folder or file exist
func Exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
	}
	return true, err
}

// Mkdir make directory if directory does not exist
func Mkdir(path string) error {
	if exist, _ := Exist(path); !exist {
		return os.MkdirAll(path, 0777)
	}
	return nil
}

// Delete delete file
func Delete(filename string) error {
	return os.Remove(filename)
}

// CreateFile write a config file
func CreateFile(filename string, contents string) error {
	directory := filepath.Dir(filename)

	if exists, _ := Exist(directory); !exists {
		if err := Mkdir(directory); err != nil {
			return err
		}
	}

	ext := filepath.Ext(filename)
	filename = strings.Replace(filename, ext, "", 1)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(contents)
	return nil
}
