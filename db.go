package main

import (
	"github.com/geekr-dev/gpt-engineer/db/file"
	"path/filepath"
)

type DB interface {
	Set(key, val string) error
	Get(key string) (string, error)
	Path() string
}

type DBs struct {
	canExecute bool
	dbType     string
	memory     DB
	logs       DB
	identity   DB
	input      DB
	workspace  DB
}

func NewFileDBs(rootPath string, projectPath string) DBs {
	memoryPath := filepath.Join(projectPath, "memory")
	workspacePath := filepath.Join(projectPath, "workspace")
	identityPath := filepath.Join(rootPath, "identity")
	return DBs{
		canExecute: file.CanExecute,
		dbType:     file.DbType,
		memory:     file.New(memoryPath),
		logs:       file.New(filepath.Join(memoryPath, "logs")),
		identity:   file.New(identityPath),
		input:      file.New(projectPath),
		workspace:  file.New(workspacePath),
	}
}
