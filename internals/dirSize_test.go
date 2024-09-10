package internals_test

import (
	"bytes"
	"fmt"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/Skylli/go-dinf/internals"
	"github.com/stretchr/testify/assert"
)

func Test_DirSize(t *testing.T) {
	t.Run("No file in file system", func(t *testing.T) {
		fsys := fstest.MapFS{}

		buf := bytes.Buffer{}
		internals.DirSize(&buf, fsys, internals.DirSizeOpts{Recursive: false})
		got := buf.String()
		want := fmt.Sprintf(internals.SizeFormat, 0)
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}

		bufR := bytes.Buffer{}
		internals.DirSize(&bufR, fsys, internals.DirSizeOpts{Recursive: true})
		gotR := bufR.String()
		wantR := fmt.Sprintf(internals.SizeFormat, 0)
		if gotR != wantR {
			t.Errorf("got: %q, want: %q", gotR, wantR)
		}
	})

	t.Run("Single file in file system", func(t *testing.T) {
		fsys := fstest.MapFS{
			"foo": &fstest.MapFile{Data: []byte("123")},
		}

		buf := bytes.Buffer{}
		internals.DirSize(&buf, fsys, internals.DirSizeOpts{Recursive: false})
		got := buf.String()
		want := fmt.Sprintf(internals.SizeFormat, 3)
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}

		bufR := bytes.Buffer{}
		internals.DirSize(&bufR, fsys, internals.DirSizeOpts{Recursive: true})
		gotR := bufR.String()
		wantR := fmt.Sprintf(internals.SizeFormat, 3)
		if gotR != wantR {
			t.Errorf("got: %q, want: %q", gotR, wantR)
		}
	})

	t.Run("Directory in file system", func(t *testing.T) {
		fsys := fstest.MapFS{
			"foo": &fstest.MapFile{Data: []byte("1234")},
			"dir": &fstest.MapFile{Mode: fs.ModeDir},
		}

		buf := bytes.Buffer{}
		internals.DirSize(&buf, fsys, internals.DirSizeOpts{Recursive: false})
		got := buf.String()
		want := fmt.Sprintf(internals.SizeFormat, 4)
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}

		bufR := bytes.Buffer{}
		internals.DirSize(&bufR, fsys, internals.DirSizeOpts{Recursive: true})
		gotR := bufR.String()
		wantR := fmt.Sprintf(internals.SizeFormat, 4)
		if gotR != wantR {
			t.Errorf("got: %q, want: %q", gotR, wantR)
		}
	})

	t.Run("Directory with files in file system", func(t *testing.T) {
		fsys := fstest.MapFS{
			"foo":   &fstest.MapFile{Data: []byte("1234")},
			"dir":   &fstest.MapFile{Mode: fs.ModeDir},
			"dir/a": &fstest.MapFile{Data: []byte("1234")},
			"dir/b": &fstest.MapFile{Data: []byte("1234")},
			"dir/c": &fstest.MapFile{Data: []byte("1234")},
		}
		testcases := []struct {
			expected string
			opts     internals.DirSizeOpts
		}{
			{
				expected: fmt.Sprintf(internals.SizeFormat, 4),
				opts:     internals.DirSizeOpts{Recursive: false},
			},
			{
				expected: fmt.Sprintf(internals.SizeRawFormat, 4),
				opts:     internals.DirSizeOpts{Recursive: false, Raw: true},
			},
			{
				expected: fmt.Sprintf(internals.SizeFormat, 16),
				opts:     internals.DirSizeOpts{Recursive: true},
			},
			{
				expected: fmt.Sprintf(internals.SizeRawFormat, 16),
				opts:     internals.DirSizeOpts{Recursive: true, Raw: true},
			},
		}

		for i, tc := range testcases {
			buf := bytes.Buffer{}
			internals.DirSize(&buf, fsys, tc.opts)
			actual := buf.String()

			assert.Equal(
				t,
				tc.expected,
				actual,
				fmt.Sprintf("internals.DirSize exec case %d", i),
			)
		}
	})
}
