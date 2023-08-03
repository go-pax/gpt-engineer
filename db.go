package main

import (
	"github.com/geekr-dev/gpt-engineer/database"
	"github.com/geekr-dev/gpt-engineer/database/file"
	_ "github.com/geekr-dev/gpt-engineer/database/github"
	_ "github.com/geekr-dev/gpt-engineer/database/memory"
)

type DBs struct {
	canExecute bool
	dbType     string
	memory     database.Database
	logs       database.Database
	identity   database.Database
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

	identityDB, err := database.Open("file://./", "identity")
	if err != nil {
		return DBs{}, err
	}

	return DBs{
		canExecute: file.CanExecute,
		memory:     memoryDB,
		logs:       logDB,
		input:      inputDB,
		workspace:  workspaceDB,
		identity:   identityDB,
	}, nil
}
