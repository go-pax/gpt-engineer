package main

import (
	"github.com/geekr-dev/gpt-engineer/db"
	"path/filepath"
)

type DB interface {
	Set(key, val string) error
	Get(key string) (string, error)
}

type DBs struct {
	memory    DB
	logs      DB
	identity  DB
	input     DB
	workspace DB
}

func NewFileDBs(rootPath string, projectPath string) DBs {
	memoryPath := filepath.Join(projectPath, "memory")
	workspacePath := filepath.Join(projectPath, "workspace")
	identityPath := filepath.Join(rootPath, "identity")
	return DBs{
		memory:    db.NewFile(memoryPath),
		logs:      db.NewFile(filepath.Join(memoryPath, "logs")),
		identity:  db.NewFile(identityPath),
		input:     db.NewFile(projectPath),
		workspace: db.NewFile(workspacePath),
	}
}
