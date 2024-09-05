package cmd_test

import (
	"bytes"
	"fmt"
	"dinf/cmd"
	"dinf/internals"
	"io"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

var emptyFS = fstest.MapFS{}

func Test_NewFcCmd(t *testing.T) {
	t.Run("FcCmd should have a recursive flag of type Bool", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		_, err := fcCmd.Flags().GetBool("recursive")
		fcCmd.Execute()
		assert.NoError(t, err, "FcCmd should have a recursive flag")
	})
	t.Run("By default, recursive flag should be false", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		b, err := fcCmd.Flags().GetBool("recursive")
		fcCmd.Execute()
		if assert.NoError(t, err, "FcCmd should have a recursive flag") {
			assert.False(t, b, "recursive flag should be flase by default")
		}
	})
	t.Run("--recursive should pass the flag recursive to true", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		fcCmd.SetArgs([]string{"--recursive"})
		fcCmd.Execute()
		b, err := fcCmd.Flags().GetBool("recursive")
		if assert.NoError(t, err, "FcCmd should have a recursive flag") {
			assert.True(t, b, "recursive flag should be true if '--recursive' is specified in the args")
		}
	})
	t.Run("-R should pass the flag recursive to true", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		fcCmd.SetArgs([]string{"-R"})
		fcCmd.Execute()
		b, err := fcCmd.Flags().GetBool("recursive")
		if assert.NoError(t, err, "FcCmd should have a recursive flag") {
			assert.True(t, b, "recursive flag should be true if '-R' is specified in the args")
		}
	})

	t.Run("FcCmd should have a raw flag of type Bool", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		_, err := fcCmd.Flags().GetBool("raw")
		fcCmd.Execute()
		assert.NoError(t, err)
	})
	t.Run("By default, raw flag should be false", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		b, err := fcCmd.Flags().GetBool("raw")
		fcCmd.Execute()
		if assert.NoError(t, err, "FcCmd should have a raw flag") {
			assert.False(t, b, "raw flag should be flase by default")
		}
	})
	t.Run("--raw should pass the flag raw to true", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		fcCmd.SetArgs([]string{"--raw"})
		fcCmd.Execute()
		b, err := fcCmd.Flags().GetBool("raw")
		if assert.NoError(t, err, "FcCmd should have a raw flag") {
			assert.True(t, b, "raw flag should be true if '--raw' is specified in the args")
		}
	})
	t.Run("-r should pass the flag raw to true", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		fcCmd.SetArgs([]string{"-r"})
		fcCmd.Execute()
		b, err := fcCmd.Flags().GetBool("raw")
		if assert.NoError(t, err, "FcCmd should have a raw flag") {
			assert.True(t, b, "raw flag should be true if '-r' is specified in the args")
		}
	})
}

func Test_FcCmdExecute(t *testing.T) {
	testcases := []struct {
		fsys     fstest.MapFS
		expected string
		args     []string
	}{
		{
			fsys:     fstest.MapFS{},
			expected: fmt.Sprintf(internals.FileCountFormat, 0),
			args:     []string{},
		},
		{
			fsys:     fstest.MapFS{},
			expected: fmt.Sprintf(internals.FileCountFormat, 0),
			args:     []string{"--recursive"},
		},
		{
			fsys: fstest.MapFS{
				"file_1": &fstest.MapFile{},
				"file_2": &fstest.MapFile{},
			},
			expected: fmt.Sprintf(internals.FileCountFormat, 2),
			args:     []string{},
		},
		{
			fsys: fstest.MapFS{
				"file_1": &fstest.MapFile{},
				"file_2": &fstest.MapFile{},
			},
			expected: fmt.Sprintf(internals.FileCountFormat, 2),
			args:     []string{"--recursive"},
		},
		{
			fsys: fstest.MapFS{
				"file_1":          &fstest.MapFile{},
				"file_2":          &fstest.MapFile{},
				"folder_1/file_1": &fstest.MapFile{},
			},
			expected: fmt.Sprintf(internals.FileCountFormat, 2),
			args:     []string{},
		},
		{
			fsys: fstest.MapFS{
				"file_1":          &fstest.MapFile{},
				"file_2":          &fstest.MapFile{},
				"folder_1/file_1": &fstest.MapFile{},
			},
			expected: fmt.Sprintf(internals.FileCountFormat, 3),
			args:     []string{"-R"},
		},
		{
			fsys: fstest.MapFS{
				"file_1":          &fstest.MapFile{},
				"file_2":          &fstest.MapFile{},
				"folder_1/file_1": &fstest.MapFile{},
			},
			expected: fmt.Sprintf(internals.FileCountFormat, 3),
			args:     []string{"--recursive"},
		},
		{
			fsys: fstest.MapFS{
				"file_1":          &fstest.MapFile{},
				"file_2":          &fstest.MapFile{},
				"folder_1/file_1": &fstest.MapFile{},
				"folder_2/file_1": &fstest.MapFile{},
			},
			expected: fmt.Sprintf(internals.FileCountFormat, 4),
			args:     []string{"--recursive"},
		},
		{
			fsys: fstest.MapFS{
				"file_1":                       &fstest.MapFile{},
				"file_2":                       &fstest.MapFile{},
				"folder_1/file_1":              &fstest.MapFile{},
				"folder_1/sub_folder_1/file_1": &fstest.MapFile{},
				"folder_1/sub_folder_2":        &fstest.MapFile{Mode: fs.ModeDir},
				"folder_2/file_1":              &fstest.MapFile{},
				"folder_2/sub_folder_1":        &fstest.MapFile{Mode: fs.ModeDir},
			},
			expected: fmt.Sprintf(internals.FileCountFormat, 5),
			args:     []string{"--recursive"},
		},
		{
			fsys: fstest.MapFS{
				"file_1":                       &fstest.MapFile{},
				"file_2":                       &fstest.MapFile{},
				"folder_1/file_1":              &fstest.MapFile{},
				"folder_1/sub_folder_1/file_1": &fstest.MapFile{},
				"folder_1/sub_folder_2":        &fstest.MapFile{Mode: fs.ModeDir},
				"folder_2/file_1":              &fstest.MapFile{},
				"folder_2/sub_folder_1":        &fstest.MapFile{Mode: fs.ModeDir},
			},
			expected: fmt.Sprintf(internals.FileCountRawFormat, 5),
			args:     []string{"--recursive", "--raw"},
		},
		{
			fsys: fstest.MapFS{
				"file_1":                       &fstest.MapFile{},
				"file_2":                       &fstest.MapFile{},
				"folder_1/file_1":              &fstest.MapFile{},
				"folder_1/sub_folder_1/file_1": &fstest.MapFile{},
				"folder_1/sub_folder_2":        &fstest.MapFile{Mode: fs.ModeDir},
				"folder_2/file_1":              &fstest.MapFile{},
				"folder_2/sub_folder_1":        &fstest.MapFile{Mode: fs.ModeDir},
			},
			expected: fmt.Sprintf(internals.FileCountRawFormat, 5),
			args:     []string{"-rR"},
		},
		{
			fsys: fstest.MapFS{
				"file_1":                       &fstest.MapFile{},
				"file_2":                       &fstest.MapFile{},
				"folder_1/file_1":              &fstest.MapFile{},
				"folder_1/sub_folder_1/file_1": &fstest.MapFile{},
				"folder_1/sub_folder_2":        &fstest.MapFile{Mode: fs.ModeDir},
				"folder_2/file_1":              &fstest.MapFile{},
				"folder_2/sub_folder_1":        &fstest.MapFile{Mode: fs.ModeDir},
			},
			expected: fmt.Sprintf(internals.FileCountRawFormat, 2),
			args:     []string{"-r"},
		},
	}

	for i, tc := range testcases {
		t.Run(
			fmt.Sprintf("FileCount exec case %d - args: %v", i, tc.args),
			func(t *testing.T) {
				fcCmd := cmd.NewFcCmd(tc.fsys)
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
