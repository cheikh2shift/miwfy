package file_test

import (
	"slices"
	"strings"
	"testing"

	file "github.com/cheikh2shift/miwfy/osfile"
)

type testFile struct {
	buf []byte
}

func (t *testFile) Write(p []byte) (n int, err error) {
	t.buf = p
	return 0, nil
}

// =================
// Test function
func TestGet(t *testing.T) {

	// construct the "test file"
	f := strings.NewReader("Hello World")

	d, _ := file.Get(f)

	comp := []byte("Hello World")

	if !slices.Equal(d, comp) {
		t.Errorf("failed, expecting %s got %s", comp, d)
	}

}

func TestSave(t *testing.T) {

	// construct the "test file"
	f := &testFile{}

	file.Save(
		f,
		"Hello World",
	)

	comp := []byte("Hello World")

	if !slices.Equal(f.buf, comp) {
		t.Errorf("failed, expecting %s got %s", comp, f.buf)
	}

}
