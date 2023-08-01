package file

import (
	"github.com/geekr-dev/gpt-engineer/database"
	"io"
	"io/fs"
	nurl "net/url"
	"os"
	"path"
	"path/filepath"
)

const (
	CanExecute = true
)

func init() {
	database.Register("file", &File{})
}

type File struct {
	path         string
	url          string
	subDirectory string
	fsys         fs.FS
}

func (f File) Open(url, subDir string) (database.Database, error) {
	p, err := parseURL(url)
	if err != nil {
		return nil, err
	}
	if subDir != "" {
		p = path.Join(p, subDir)
	}
	nf := &File{
		url:  url,
		path: p,
		fsys: os.DirFS(p),
	}
	return nf, nil
}

func parseURL(url string) (string, error) {
	u, err := nurl.Parse(url)
	if err != nil {
		return "", err
	}
	// concat host and path to restore full path
	// host might be `.`
	p := u.Opaque
	if len(p) == 0 {
		p = u.Host + u.Path
	}

	if len(p) == 0 {
		// default to current directory if no path
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		p = wd

	} else if p[0:1] == "." || p[0:1] != "/" {
		// make path absolute if relative
		abs, err := filepath.Abs(p)
		if err != nil {
			return "", err
		}
		p = abs
	}
	return p, nil
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

func (f File) Path() string {
	return f.path
}

// Close is part of source.Driver interface implementation.
// Closes the file system if possible.
func (f File) Close() error {
	c, ok := f.fsys.(io.Closer)
	if !ok {
		return nil
	}
	return c.Close()
}
