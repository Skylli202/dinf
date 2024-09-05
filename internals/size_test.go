package internals_test

import (
	"bytes"
	"fmt"
	"dinf/internals"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestDinf(t *testing.T) {
	t.Run("No file in file system", func(t *testing.T) {
		fsys := fstest.MapFS{}

		buf := bytes.Buffer{}
		internals.DirSizeCmd(&buf, fsys)

		got := buf.String()
		want := fmt.Sprintf("Folder size is: %d bytes.\n", 0)

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("Single file in file system", func(t *testing.T) {
		fsys := fstest.MapFS{
			"foo": &fstest.MapFile{Data: []byte("123")},
		}

		buf := bytes.Buffer{}
		internals.DirSizeCmd(&buf, fsys)

		got := buf.String()
		want := fmt.Sprintf("Folder size is: %d bytes.\n", 3)

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("Directory in file system", func(t *testing.T) {
		fsys := fstest.MapFS{
			"foo":  &fstest.MapFile{Data: []byte("1234")},
			"dir/": &fstest.MapFile{Mode: fs.ModeDir},
		}

		buf := bytes.Buffer{}
		internals.DirSizeCmd(&buf, fsys)

		got := buf.String()
		want := fmt.Sprintf("Folder size is: %d bytes.\n", 4)

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("Directory with files in file system", func(t *testing.T) {
		fsys := fstest.MapFS{
			"foo":   &fstest.MapFile{Data: []byte("1234")},
			"dir/":  &fstest.MapFile{Mode: fs.ModeDir},
			"dir/a": &fstest.MapFile{Data: []byte("1234")},
			"dir/b": &fstest.MapFile{Data: []byte("1234")},
			"dir/c": &fstest.MapFile{Data: []byte("1234")},
		}

		buf := bytes.Buffer{}
		internals.DirSizeCmd(&buf, fsys)

		got := buf.String()
		want := fmt.Sprintf("Folder size is: %d bytes.\n", 4)

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
}
