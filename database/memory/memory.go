package memory

import (
	"fmt"
	"github.com/geekr-dev/gpt-engineer/database"
)

const (
	CanExecute = true
)

func init() {
	database.Register("memory", &Memory{})
}

type Memory struct {
	subDirectory string
	db           map[string]string
}

func (m *Memory) Open(url, subDir string) (database.Database, error) {
	md := &Memory{
		subDirectory: subDir,
		db:           make(map[string]string),
	}
	return md, nil
}

func (m *Memory) Get(key string) (string, error) {
	content, exists := m.db[key]
	if !exists {
		return "", fmt.Errorf("key doesn't exist, %s", key)
	}
	return content, nil
}

func (m *Memory) Set(key, val string) error {
	m.db[key] = val
	return nil
}

func (m *Memory) Path() string {
	return ""
}

// Close is part of source.Driver interface implementation.
// Closes the file system if possible.
func (m *Memory) Close() error {
	return nil
}
