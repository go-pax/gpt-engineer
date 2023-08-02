// Package testing has the source tests.
// All source drivers must pass the Test function.
// This lives in it's own package so it stays a test dependency.
package testing

import (
	"github.com/geekr-dev/gpt-engineer/database"
	"testing"
)

const (
	TestFile     = "test.txt"
	TestContents = "contents of test file"
)

// Test runs tests against database implementations.
// It assumes that the database tests can read and write a file named test.txt.
//
// See database/file/file_test.go for an example.
func Test(t *testing.T, d database.Database) {
	TestSet(t, d)
	TestGet(t, d)
}

func TestGet(t *testing.T, d database.Database) {
	contents, err := d.Get(TestFile)
	if err != nil {
		t.Fatalf("Get: expected err to be nil, got %v", err)
	}
	if contents != TestContents {
		t.Errorf("Get: expected %s, got %s", TestContents, contents)
	}
}

func TestSet(t *testing.T, d database.Database) {
	err := d.Set(TestFile, TestContents)
	if err != nil {
		t.Fatalf("Set: expected err to be nil, got %v", err)
	}
}
