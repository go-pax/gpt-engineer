package main

import (
	"os"
	"path/filepath"
)

type DB struct {
	path string
}

func NewDB(path string) *DB {
	absPath, _ := filepath.Abs(path)
	_ = os.MkdirAll(absPath, os.ModePerm)
	return &DB{path: absPath}
}

func (db *DB) Get(key string) (string, error) {
	content, err := os.ReadFile(filepath.Join(db.path, key))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (db *DB) Set(key, val string) error {
	return os.WriteFile(filepath.Join(db.path, key), []byte(val), 0644)
}

type DBs struct {
	memory    *DB
	logs      *DB
	identity  *DB
	input     *DB
	workspace *DB
}

func NewDBs(rootPath string, projectPath string) *DBs {
	memoryPath := filepath.Join(projectPath, "memory")
	workspacePath := filepath.Join(projectPath, "workspace")
	identityPath := filepath.Join(rootPath, "identity")
	return &DBs{
		memory:    NewDB(memoryPath),
		logs:      NewDB(filepath.Join(memoryPath, "logs")),
		identity:  NewDB(identityPath),
		input:     NewDB(projectPath),
		workspace: NewDB(workspacePath),
	}
}
