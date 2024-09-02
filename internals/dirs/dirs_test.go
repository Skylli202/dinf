package dirs_test

import (
	"fmt"
	"go-dinf/internals/dirs"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestDirSize(t *testing.T) {
	cases := []struct {
		fs           fstest.MapFS
		msg          string
		expectedSize int64
	}{
		{
			fs:           fstest.MapFS{},
			msg:          "Empty FS should have a size of 0.",
			expectedSize: 0,
		},
		{
			fs: fstest.MapFS{
				"foo": &fstest.MapFile{Data: []byte("123")},
			},
			msg:          "Single file FS should have a size of 3.",
			expectedSize: 3,
		},
		{
			fs: fstest.MapFS{
				"foo": &fstest.MapFile{Data: []byte("123")},
				"bar": &fstest.MapFile{Data: []byte("456")},
			},
			msg:          "Two files FS should have a size of 6.",
			expectedSize: 6,
		},
		{
			fs: fstest.MapFS{
				"foo": &fstest.MapFile{Mode: fs.ModeDir},
				"bar": &fstest.MapFile{Mode: fs.ModeDir},
			},
			msg:          "A file system with only directories should have a size of 0.",
			expectedSize: 0,
		},
		{
			fs: fstest.MapFS{
				"foo":     &fstest.MapFile{Mode: fs.ModeDir},
				"foo/foo": &fstest.MapFile{Data: []byte("123")},
				"bar":     &fstest.MapFile{Mode: fs.ModeDir},
			},
			msg:          "It should only look at direct child (i.e. ignore subdirectories files).",
			expectedSize: 0,
		},
	}

	for i, tc := range cases {
		size, err := dirs.DirSize(tc.fs)

		assert.Nil(t, err, fmt.Sprintf("DirSize should return err equals to Nil, but it is not for test case: [%d] %s", i, tc.msg))
		assert.Equal(t, tc.expectedSize, size, tc.msg)
	}
}

type (
	mMapFS struct {
		fstest.MapFS
		errors map[string]error
	}
)

func (fsys mMapFS) ReadDir(name string) ([]fs.DirEntry, error) {
	err, ok := fsys.errors[name]
	fmt.Printf("fsys.errors[%s]: %v, %t\n", name, err, ok)
	if ok {
		return make([]fs.DirEntry, 0), err
	}

	fmt.Printf("[TEST] fsys.ReadDir(%s)\n", name)

	// For now, return an empty DirEntry slices.
	dirents, err := fsys.MapFS.ReadDir(name)
	if err != nil {
		panic(fmt.Sprintf("Cannot ReadDir in-memory file system in test: %v", err))
	}

	results := make([]fs.DirEntry, len(dirents))
	// for _, v := range dirents {
	// 	e, ok := fsys.errors[v.Name()]
	// 	if ok {
	// 		results = append(results, mDirEntry{
	// 			DirEntry: v,
	// 			err:      e,
	// 		})
	// 	} else {
	// 		results = append(results, mDirEntry{v, nil})
	// 	}
	// }

	return results, nil
}

var ErrMockFSReadDir = fmt.Errorf("MockFS.ReadDir(): forced error")

func TestDirSizeErrorOnReadDir(t *testing.T) {
	cases := []struct {
		expectedError error
		fs            fs.FS
		msg           string
	}{
		{
			expectedError: ErrMockFSReadDir,
			fs: mMapFS{
				MapFS: fstest.MapFS{
					"foo": {Data: []byte("123")},
				},
				errors: map[string]error{
					".": ErrMockFSReadDir,
				},
			},
			msg: "DirSize should return the Error returned by ReadDir, with a size of 0.",
		},
	}

	for i, tc := range cases {
		size, err := dirs.DirSize(tc.fs)

		fmt.Printf("DirSize: %d, %v\n", size, err)
		assert.NotNil(t, err, fmt.Sprintf("DirSize should return an error not Nil, test case: [%d] %s", i, tc.msg))
		assert.Equal(t, tc.expectedError, err, tc.msg)
		assert.Equal(t, int64(0), size, tc.msg)
		fmt.Println("---")
	}
}

// Test FileCount and FileCountR
func TestFileCount(t *testing.T) {
	t.Run("FileCount error test cases", func(t *testing.T) {
		testcases := []struct {
			fsys          fs.FS
			expectedError error
			msg           string
			expectedCount int
		}{
			{
				fsys: mMapFS{
					MapFS: fstest.MapFS{
						"file_1":    {Data: []byte("123")},
						"folder_1/": {Mode: fs.ModeDir},
					},
					errors: map[string]error{
						".": ErrMockFSReadDir,
					},
				},
				expectedError: ErrMockFSReadDir,
				msg:           "If ReadDir fail, FileCount should return this error and a count of 1.",
				expectedCount: 0,
			},
		}

		for _, tc := range testcases {
			got, err := dirs.FileCount(tc.fsys)
			assert.NotNil(t, err, "FileCount should return an error.")
			assert.Equal(t, err, ErrMockFSReadDir, "When ReadDir within FileCount return an error, FileCount should return this error and a count of 0.")
			assert.Equal(t, tc.expectedCount, got, "When FileCount return an error, count should be equal to 0.")
		}
	})

	t.Run("FileCount error free test cases", func(t *testing.T) {
		testcases := []struct {
			fsys          fstest.MapFS
			msg           string
			expectedCount int
		}{
			{
				fsys:          fstest.MapFS{},
				msg:           "Empty directory should return a file count of 0.",
				expectedCount: 0,
			},
			{
				fsys: fstest.MapFS{
					"folder_1/": &fstest.MapFile{Mode: fs.ModeDir},
					"folder_2/": &fstest.MapFile{Mode: fs.ModeDir},
					"folder_3/": &fstest.MapFile{Mode: fs.ModeDir},
				},
				msg:           "Directory that contain only directories should return a file count of 0.",
				expectedCount: 0,
			},
			{
				fsys: fstest.MapFS{
					"file_1": &fstest.MapFile{Data: []byte("123")},
				},
				msg:           "Directory with single file should return a file count of 1.",
				expectedCount: 1,
			},
		}

		for _, tc := range testcases {
			got, err := dirs.FileCount(tc.fsys)
			assert.Nil(t, err, tc.msg)
			assert.Equal(t, tc.expectedCount, got, tc.msg)
		}
	})

	t.Run("FileCountR error free test cases", func(t *testing.T) {
		testcases := []struct {
			fsys          fstest.MapFS
			msg           string
			expectedCount int
		}{
			{
				fsys:          fstest.MapFS{},
				msg:           "Empty directory should return a file count of 0.",
				expectedCount: 0,
			},
			{
				fsys: fstest.MapFS{
					"folder_1": &fstest.MapFile{Mode: fs.ModeDir},
					"folder_2": &fstest.MapFile{Mode: fs.ModeDir},
					"folder_3": &fstest.MapFile{Mode: fs.ModeDir},
				},
				msg:           "Directory that contain only directories should return a recusive file count of 0.",
				expectedCount: 0,
			},
			{
				fsys: fstest.MapFS{
					"file_1": &fstest.MapFile{},
				},
				msg:           "Directory with single file should return a file count of 1.",
				expectedCount: 1,
			},
			{
				fsys: fstest.MapFS{
					"folder_1/file_1": &fstest.MapFile{Data: []byte("123")},
				},
				msg:           "Directory with a single subdirectory and a single file should return a file count of 1.",
				expectedCount: 1,
			},
			{
				fsys: fstest.MapFS{
					"folder_1/file_1": &fstest.MapFile{Data: []byte("123")},
					"fake_folder_2":   &fstest.MapFile{Mode: fs.ModeDir},
				},
				msg:           "Directory with two subdirectories, one with a single file and one empty, should return a file count of 1.",
				expectedCount: 1,
			},
			{
				fsys: fstest.MapFS{
					"folder_1/file_1": &fstest.MapFile{Data: []byte("123")},
					"folder_2/file_2": &fstest.MapFile{Data: []byte("123")},
				},
				msg:           "Directory with two subdirectories, with a single file in each, should return a file count of 2.",
				expectedCount: 2,
			},
			{
				fsys: fstest.MapFS{
					"folder_1/file_1": &fstest.MapFile{Data: []byte("123")},
					"folder_1/file_3": &fstest.MapFile{Data: []byte("123")},
					"folder_2/file_2": &fstest.MapFile{Data: []byte("123")},
				},
				msg:           "Directory with two subdirectories, with one or more file in each, should return a file count of 3.",
				expectedCount: 3,
			},
		}

		for _, tc := range testcases {
			got, err := dirs.FileCountR(tc.fsys)
			assert.Nil(t, err, tc.msg)
			assert.Equal(t, tc.expectedCount, got, tc.msg, tc.fsys)
		}
	})
}
