package cmd_test

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/Skylli202/go-dinf/cmd"
	"github.com/Skylli202/go-dinf/internals"
	"github.com/stretchr/testify/assert"
)

func Test_NewSizeCmd(t *testing.T) {
	t.Run(
		"SizeCmd should have a recursive flag of type Bool, and its default value should be false",
		func(t *testing.T) {
			sizeCmd := cmd.NewSizeCmd(emptyFS)
			b, err := sizeCmd.Flags().GetBool("recursive")
			sizeCmd.Execute()
			if assert.NoError(t, err, "FcCmd should have a recursive flag") {
				assert.False(t, b, "recursive flag should be flase by default")
			}
		})
	t.Run("--recursive should pass the flag recursive to true", func(t *testing.T) {
		sizeCmd := cmd.NewSizeCmd(emptyFS)
		sizeCmd.SetArgs([]string{"--recursive"})
		sizeCmd.Execute()
		b, err := sizeCmd.Flags().GetBool("recursive")
		if assert.NoError(t, err, "SizeCmd should have a recursive flag") {
			assert.True(t, b, "recursive flag should be true if '--recursive' is specified in the args")
		}
	})
	t.Run("-R should pass the flag recursive to true", func(t *testing.T) {
		sizeCmd := cmd.NewSizeCmd(emptyFS)
		sizeCmd.SetArgs([]string{"-R"})
		sizeCmd.Execute()
		b, err := sizeCmd.Flags().GetBool("recursive")
		if assert.NoError(t, err, "SizeCmd should have a recursive flag") {
			assert.True(t, b, "recursive flag should be true if '-R' is specified in the args")
		}
	})

	t.Run(
		"SizeCmd should have a raw flag of type Bool, and its default value should be false",
		func(t *testing.T) {
			sizeCmd := cmd.NewSizeCmd(emptyFS)
			b, err := sizeCmd.Flags().GetBool("raw")
			sizeCmd.Execute()
			if assert.NoError(t, err, "FcCmd should have a raw flag") {
				assert.False(t, b, "raw flag should be flase by default")
			}
		})
	t.Run("--raw should pass the flag raw to true", func(t *testing.T) {
		sizeCmd := cmd.NewSizeCmd(emptyFS)
		sizeCmd.SetArgs([]string{"--raw"})
		sizeCmd.Execute()
		b, err := sizeCmd.Flags().GetBool("raw")
		if assert.NoError(t, err, "SizeCmd should have a raw flag") {
			assert.True(t, b, "raw flag should be true if '--raw' is specified in the args")
		}
	})
	t.Run("-R should pass the flag raw to true", func(t *testing.T) {
		sizeCmd := cmd.NewSizeCmd(emptyFS)
		sizeCmd.SetArgs([]string{"-r"})
		sizeCmd.Execute()
		b, err := sizeCmd.Flags().GetBool("raw")
		if assert.NoError(t, err, "SizeCmd should have a raw flag") {
			assert.True(t, b, "raw flag should be true if '-R' is specified in the args")
		}
	})
}

type testcase struct {
	fsys     fs.FS
	expected string
	args     []string
}

func Test_SizeCmdExecute(t *testing.T) {
	testcases := []testcase{
		// 0
		{
			fsys:     fstest.MapFS{},
			expected: fmt.Sprintf(internals.SizeFormat, 0),
			args:     []string{},
		},
		// 1
		{
			fsys: fstest.MapFS{
				"file_1": &fstest.MapFile{},
				"file_2": &fstest.MapFile{},
			},
			expected: fmt.Sprintf(internals.SizeFormat, 0),
			args:     []string{},
		},
		// 2
		{
			fsys: fstest.MapFS{
				"file_1": &fstest.MapFile{Data: []byte("123")},
				"file_2": &fstest.MapFile{Data: []byte("123")},
			},
			expected: fmt.Sprintf(internals.SizeFormat, 6),
			args:     []string{},
		},
		// 3
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
		// 4
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
			expected: fmt.Sprintf(internals.SizeFormat, 5),
			args:     []string{"--recursive"},
		},
		// 5
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
			expected: fmt.Sprintf(internals.SizeRawFormat, 2),
			args:     []string{"--raw"},
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
						tc,
					)
				}
			},
		)
	}
}

func (tc testcase) String() string {
	s := fmt.Sprintf("Cmd args are: [%s]\n", strings.Join(tc.args, ", "))
	s = s + mapToString(tc.fsys)
	return s
}

func mapToString(fsys fs.FS) (s string) {
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		i, e := d.Info()
		if e != nil {
			s = s + fmt.Sprintf("%s\n", d)
		}
		s = s + fmt.Sprintf("%s %v %s\n", i.Mode(), i.Size(), path)
		return nil
	})
	return
}
