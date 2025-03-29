package dirs_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Skylli202/dinf/internals/dirs"
	"github.com/stretchr/testify/assert"
)

func TestLineCount(t *testing.T) {
	t.Run("Empty dir", func(t *testing.T) {
		emptyDirName, err := os.MkdirTemp("", "")
		if err != nil {
			t.Fatal("Unable to `os.MkdirTemp(...)`. Unable to execute tests.")
		}
		defer os.RemoveAll(emptyDirName)

		msg := "Empty file system."
		expectedLineCount := 0

		fs := os.DirFS(emptyDirName)
		lc, err := dirs.LineCount(fs)

		assert.Nil(t, err, fmt.Sprintf("LineCount should return err equals to Nil, but it is not for test case: [%d] %s", 999, msg))
		// assert.Nil(t, errR, fmt.Sprintf("LineCountR should return err equals to Nil, but it is not for test case: [%d] %s", i, tc.msg))
		assert.Equal(t, expectedLineCount, lc, fmt.Sprintf("Line count does not match.\nTest case: %d - %q.\nTest case's file system:\n%+v\n", 999, msg, emptyDirName))
		// assert.Equal(t, tc.expectedSizeR, sizeR, fmt.Sprintf("Line count R  does not match.\nTest case: %d - %q.\nTest case's file system:\n%+v\n", i, tc.msg, tc.fs))
	})
	// dname, err := os.MkdirTemp("", "testlinecount")
	// if err != nil {
	// 	t.Fatal("Unable to `os.MkdirTemp(...)`. Unable to execute tests.")
	// }
	// defer os.RemoveAll(dname)

	// file1, err := os.CreateTemp("", "")
	// cases := []struct {
	// 	fs                 fstest.MapFS
	// 	msg                string
	// 	expectedLineCount  int
	// 	expectedLineCountR int
	// }{
	// 	{
	// 		fs:                 fstest.MapFS{},
	// 		msg:                "Empty file system.",
	// 		expectedLineCount:  0,
	// 		expectedLineCountR: 0,
	// 	},
	// 	{
	// 		fs: fstest.MapFS{
	// 			"foo": &fstest.MapFile{Data: []byte("123")},
	// 		},
	// 		msg:                "A single line in a single file (no subdirectories).",
	// 		expectedLineCount:  1,
	// 		expectedLineCountR: 1,
	// 	},
	// }
	// for i, tc := range cases {
	// 	lc, err := dirs.LineCount(tc.fs)
	// 	// lcR, errR := dirs.LineCountR(fsys fs.FS)
	//
	// 	assert.Nil(t, err, fmt.Sprintf("LineCount should return err equals to Nil, but it is not for test case: [%d] %s", i, tc.msg))
	// 	// assert.Nil(t, errR, fmt.Sprintf("LineCountR should return err equals to Nil, but it is not for test case: [%d] %s", i, tc.msg))
	// 	assert.Equal(t, tc.expectedLineCount, lc, fmt.Sprintf("Line count does not match.\nTest case: %d - %q.\nTest case's file system:\n%+v\n", i, tc.msg, tc.fs))
	// 	// assert.Equal(t, tc.expectedSizeR, sizeR, fmt.Sprintf("Line count R  does not match.\nTest case: %d - %q.\nTest case's file system:\n%+v\n", i, tc.msg, tc.fs))
	// }
}
