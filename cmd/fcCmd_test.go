package cmd_test

import (
	"bytes"
	"fmt"
	"go-dinf/cmd"
	"io"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

var emptyFS = fstest.MapFS{}

func Test_NewFcCmd(t *testing.T) {
	t.Run("FcCmd should have a recursive Flag of type Bool", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		_, err := fcCmd.Flags().GetBool("recursive")
		fcCmd.Execute()
		assert.NoError(t, err, "FcCmd should have a recursive flag")
	})

	t.Run("By default, recursive Flag should be false", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		b, err := fcCmd.Flags().GetBool("recursive")
		fcCmd.Execute()
		if assert.NoError(t, err, "FcCmd should have a recursive flag") {
			assert.False(t, b, "recursive flag should be flase by default")
		}
	})

	t.Run("--recursive should pass the flag to true", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		fcCmd.SetArgs([]string{"--recursive"})
		fcCmd.Execute()
		b, err := fcCmd.Flags().GetBool("recursive")
		if assert.NoError(t, err, "FcCmd should have a recursive flag") {
			assert.True(t, b, "recursive flag should be true if '--recursive' is specified in the args")
		}
	})

	t.Run("-R should pass the flag to true", func(t *testing.T) {
		fcCmd := cmd.NewFcCmd(emptyFS)
		fcCmd.SetArgs([]string{"-R"})
		fcCmd.Execute()
		b, err := fcCmd.Flags().GetBool("recursive")
		if assert.NoError(t, err, "FcCmd should have a recursive flag") {
			assert.True(t, b, "recursive flag should be true if '--recursive' is specified in the args")
		}
	})
}

func Test_FcCmdExecute(t *testing.T) {
	testcases := []struct {
		fsys     fstest.MapFS
		args     []string
		expected int
	}{
		{
			fsys:     fstest.MapFS{},
			args:     []string{},
			expected: 0,
		},
		{
			fsys:     fstest.MapFS{},
			args:     []string{"--recursive"},
			expected: 0,
		},
		{
			fsys: fstest.MapFS{
				"file_1": &fstest.MapFile{},
				"file_2": &fstest.MapFile{},
			},
			args:     []string{},
			expected: 2,
		},
		{
			fsys: fstest.MapFS{
				"file_1": &fstest.MapFile{},
				"file_2": &fstest.MapFile{},
			},
			args:     []string{"--recursive"},
			expected: 2,
		},
		{
			fsys: fstest.MapFS{
				"file_1":          &fstest.MapFile{},
				"file_2":          &fstest.MapFile{},
				"folder_1/file_1": &fstest.MapFile{},
			},
			args:     []string{},
			expected: 2,
		},
		{
			fsys: fstest.MapFS{
				"file_1":          &fstest.MapFile{},
				"file_2":          &fstest.MapFile{},
				"folder_1/file_1": &fstest.MapFile{},
			},
			args:     []string{"-R"},
			expected: 3,
		},
		{
			fsys: fstest.MapFS{
				"file_1":          &fstest.MapFile{},
				"file_2":          &fstest.MapFile{},
				"folder_1/file_1": &fstest.MapFile{},
			},
			args:     []string{"--recursive"},
			expected: 3,
		},
		{
			fsys: fstest.MapFS{
				"file_1":          &fstest.MapFile{},
				"file_2":          &fstest.MapFile{},
				"folder_1/file_1": &fstest.MapFile{},
				"folder_2/file_1": &fstest.MapFile{},
			},
			args:     []string{"--recursive"},
			expected: 4,
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
			args:     []string{"--recursive"},
			expected: 5,
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
					assert.Contains(
						t,
						string(out),
						fmt.Sprintf("%d files", tc.expected),
					)
				}
			},
		)
	}
}
