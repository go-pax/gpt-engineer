package memory

import (
	dt "github.com/geekr-dev/gpt-engineer/database/testing"
	"testing"
)

func Test(t *testing.T) {
	m := &Memory{}
	d, err := m.Open("memory://", "")
	if err != nil {
		t.Fatal(err)
	}
	dt.Test(t, d)
}
