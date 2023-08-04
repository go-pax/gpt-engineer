package file

import (
	dt "github.com/go-pax/gpt-engineer/database/testing"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	tmpDir := t.TempDir()
	f := &File{}
	d, err := f.Open("file://"+tmpDir, "")
	if err != nil {
		t.Fatal(err)
	}
	dt.Test(t, d)
}

func TestOpen(t *testing.T) {
	tmpDir := t.TempDir()
	if !filepath.IsAbs(tmpDir) {
		t.Fatal("expected tmpDir to be absolute path")
	}

	mustWriteFile(t, tmpDir, "main_prompt", "create a snake game in go that uses the keyboard for input")

	f := &File{}
	_, err := f.Open("file://"+tmpDir, "") // absolute path
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenWithRelativePath(t *testing.T) {
	tmpDir := t.TempDir()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// rescue working dir after we are done
		if err := os.Chdir(wd); err != nil {
			t.Log(err)
		}
	}()

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(tmpDir, "foo"), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	mustWriteFile(t, filepath.Join(tmpDir, "foo"), "file", "some content")

	f := &File{}

	// dir: foo
	d, err := f.Open("file://foo", "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = d.Get("file")
	if err != nil {
		t.Fatalf("expected file in dir %v for foo", tmpDir)
	}

	// dir: ./foo
	d, err = f.Open("file://./foo", "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = d.Get("file")
	if err != nil {
		t.Fatalf("expected first file in working dir %v for ./foo", tmpDir)
	}
}

func TestOpenDefaultsToCurrentDirectory(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	f := &File{}
	d, err := f.Open("file://", "")
	if err != nil {
		t.Fatal(err)
	}

	if d.(*File).path != wd {
		t.Fatal("expected database to default to current directory")
	}
}

func TestClose(t *testing.T) {
	tmpDir := t.TempDir()

	f := &File{}
	d, err := f.Open("file://"+tmpDir, "subdir")
	if err != nil {
		t.Fatal(err)
	}

	if d.Close() != nil {
		t.Fatal("expected nil")
	}
}

func mustWriteFile(t testing.TB, dir, file string, body string) {
	if err := os.WriteFile(path.Join(dir, file), []byte(body), 0600); err != nil {
		t.Fatal(err)
	}
}
