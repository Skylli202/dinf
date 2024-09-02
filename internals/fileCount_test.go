package internals_test

import (
	"bytes"
	"fmt"
	"go-dinf/internals"
	"testing"
	"testing/fstest"
)

func TestFileCount(t *testing.T) {
	t.Run("No file in file system", func(t *testing.T) {
		fsys := fstest.MapFS{}

		buf := bytes.Buffer{}
		internals.FileCountCmd(&buf, fsys)

		got := buf.String()
		want := fmt.Sprintf("Folder contains: %d files.\n", 0)

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("Single file in file system", func(t *testing.T) {
		fsys := fstest.MapFS{
			"foo": &fstest.MapFile{Data: []byte("123")},
		}

		buf := bytes.Buffer{}
		internals.FileCountCmd(&buf, fsys)

		got := buf.String()
		want := fmt.Sprintf("Folder contains: %d files.\n", 1)

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
}
