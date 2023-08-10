package main

import (
	"embed"
	"github.com/go-pax/gpt-engineer/database"
	"github.com/go-pax/gpt-engineer/database/file"
	_ "github.com/go-pax/gpt-engineer/database/github"
	_ "github.com/go-pax/gpt-engineer/database/memory"
)

//go:embed identity/*
var identityFS embed.FS

type DBs struct {
	canExecute bool
	dbType     string
	memory     database.Database
	logs       database.Database
	identity   embed.FS
	input      database.Database
	workspace  database.Database
}

func NewDBs(projectPath string, prompt string) (DBs, error) {
	memoryDB, err := database.Open(projectPath, "memory")
	if err != nil {
		return DBs{}, err
	}
	logDB, err := database.Open(projectPath, "memory/logs")
	if err != nil {
		return DBs{}, err
	}
	workspaceDB, err := database.Open(projectPath, "workspace")
	if err != nil {
		return DBs{}, err
	}
	inputDB, err := database.Open(projectPath, "")
	if err != nil {
		return DBs{}, err
	}
	if prompt != "" {
		if err := inputDB.Set("main_prompt", prompt); err != nil {
			return DBs{}, err
		}
	}

	return DBs{
		canExecute: file.CanExecute,
		memory:     memoryDB,
		logs:       logDB,
		input:      inputDB,
		workspace:  workspaceDB,
		identity:   identityFS,
	}, nil
}
