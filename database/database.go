package database

import (
	"fmt"
	neturl "net/url"
	"sync"
)

var databasesMu sync.RWMutex
var databases = make(map[string]Database)

// Database is the interface every database  must implement.
//
// How to implement?
//  1. Implement this interface.
//  2. Optionally, add a function named `WithInstance`.
//     This function should accept an existing source instance and a Config{} struct
//     and return a driver instance.
//  3. Add a test that calls source/testing.go:Test()
//  4. Add own tests for Open(), WithInstance() (when provided) and Close().
//     All other functions are tested by tests in source/testing.
//     Saves you some time and makes sure all source drivers behave the same way.
//  5. Call Register in init().
//
// Guidelines:
//   - All configuration input must come from the URL string in func Open()
//     or the Config{} struct in WithInstance. Don't os.Getenv().
//   - Drivers are supposed to be read only.
//   - Ideally don't load any contents (into memory) in Open or WithInstance.
type Database interface {
	// Open returns a new database instance configured with parameters
	// coming from the URL string.
	Open(url string, subdir string) (Database, error)

	// Close closes the underlying source instance managed by the database.
	Close() error

	// Get reads the contents of a file or resource by name.
	// If there is no file available, it must return os.ErrNotExist.
	Get(file string) (contents string, err error)

	// Set writes the contents to a file or resource.
	Set(file string, contents string) error

	// Path get the full path of the database
	Path() (path string)
}

// Open returns a new database instance.
func Open(url string, subdir string) (Database, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return nil, err
	}
	if u.Path == "" {
		return nil, fmt.Errorf(`invalid path`)
	}

	scheme := u.Scheme
	if scheme == "" {
		scheme = "file"
	}

	databasesMu.RLock()
	d, ok := databases[scheme]
	databasesMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf(`unknown database '%s' this schema isn't handled`, scheme)
	}

	return d.Open(url, subdir)
}

// Register globally registers a database.
func Register(name string, database Database) {
	databasesMu.Lock()
	defer databasesMu.Unlock()
	if database == nil {
		panic("Register database is nil")
	}
	if _, dup := databases[name]; dup {
		panic("Register called twice for database " + name)
	}
	databases[name] = database
}
