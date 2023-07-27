package db

import (
	"os"
	"path/filepath"
)

type File struct {
	path string
}

func NewFile(path string) File {
	absPath, _ := filepath.Abs(path)
	_ = os.MkdirAll(absPath, os.ModePerm)
	return File{path: absPath}
}

func (f File) Get(key string) (string, error) {
	content, err := os.ReadFile(filepath.Join(f.path, key))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (f File) Set(key, val string) error {
	return os.WriteFile(filepath.Join(f.path, key), []byte(val), 0600)
}
