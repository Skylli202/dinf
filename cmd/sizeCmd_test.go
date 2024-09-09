package cmd_test

import (
	"bytes"
	"dinf/cmd"
	"dinf/internals"
	"fmt"
	"io"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func Test_NewSizeCmd(t *testing.T) {
	t.Run("", func(t *testing.T) {})
}

func Test_SizeCmdExecute(t *testing.T) {
	testcases := []struct {
		fsys     fstest.MapFS
		expected string
		args     []string
	}{
		{
			fsys:     fstest.MapFS{},
			expected: fmt.Sprintf(internals.SizeFormat, 0),
			args:     []string{},
		},
		{
			fsys: fstest.MapFS{
				"file_1": &fstest.MapFile{},
				"file_2": &fstest.MapFile{},
			},
			expected: fmt.Sprintf(internals.SizeFormat, 0),
			args:     []string{},
		},
		{
			fsys: fstest.MapFS{
				"file_1": &fstest.MapFile{Data: []byte("123")},
				"file_2": &fstest.MapFile{Data: []byte("123")},
			},
			expected: fmt.Sprintf(internals.SizeFormat, 6),
			args:     []string{},
		},
		{
			fsys: fstest.MapFS{
				"file_1":                       &fstest.MapFile{Data: []byte("1")},
				"file_2":                       &fstest.MapFile{Data: []byte("1")},
				"folder_1/file_1":              &fstest.MapFile{Data: []byte("1")},
				"folder_1/sub_folder_1/file_1": &fstest.MapFile{Data: []byte("1")},
				"folder_1/sub_folder_2":        &fstest.MapFile{Mode: fs.ModeDir},
				"folder_2/file_1":              &fstest.MapFile{Data: []byte("1")},
				"folder_2/sub_folder_1":        &fstest.MapFile{Mode: fs.ModeDir},
			},
			expected: fmt.Sprintf(internals.SizeFormat, 2),
			args:     []string{},
		},
	}

	for i, tc := range testcases {
		t.Run(
			fmt.Sprintf("Size exec case %d - args: %v", i, tc.args),
			func(t *testing.T) {
				fcCmd := cmd.NewSizeCmd(tc.fsys)
				b := bytes.NewBufferString("")
				fcCmd.SetOut(b)
				fcCmd.SetArgs(tc.args)

				fcCmd.Execute()

				out, err := io.ReadAll(b)
				if assert.NoError(t, err) {
					assert.Equal(
						t,
						tc.expected,
						string(out),
					)
				}
			},
		)
	}
}
