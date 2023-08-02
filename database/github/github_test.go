package github

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	st "github.com/geekr-dev/gpt-engineer/database/testing"
)

var GithubTestSecret = "" // username:token

func init() {
	secrets, err := os.ReadFile(".github_test_secrets")
	if err == nil {
		GithubTestSecret = string(bytes.TrimSpace(secrets))
	}
}

func Test(t *testing.T) {
	if len(GithubTestSecret) == 0 {
		t.Skip("test requires .github_test_secrets")
	}

	g := &Github{}
	d, err := g.Open("github://"+GithubTestSecret+"@test-dump/gpt-engineer-output/test#main", "")
	if err != nil {
		t.Fatal(err)
	}

	st.Test(t, d)
}

func TestUnauthenticatedSet(t *testing.T) {
	g := &Github{}
	d, err := g.Open("github://test-dump/gpt-engineer-output/test", "")
	if err != nil {
		t.Fatal(err)
	}

	err = d.Set("will_fail.txt", "contents")
	assert.ErrorIs(t, err, ErrNoAccess)
}

func TestUnauthenticatedGet(t *testing.T) {
	g := &Github{}
	d, err := g.Open("github://test-dump/gpt-engineer-output/test", "")
	if err != nil {
		t.Fatal(err)
	}

	content, err := d.Get(st.TestFile)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, st.TestContents, content)
}

func TestUnknownRefGet(t *testing.T) {
	g := &Github{}
	d, err := g.Open("github://test-dump/gpt-engineer-output/test#doesntexist", "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = d.Get(st.TestFile)
	assert.ErrorIs(t, err, ErrInvalidRef)
}
