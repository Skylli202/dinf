package dirs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Given a filesystem (interface fs.FS) is read the directory
// and return the summed file size of its files.
//
// If [ReadDir] return an error, DirSize will return a size of 0 alongside the
// error returned by ReadDir. If [Info] return an error and the error is of type
// fs.ErrNotExist, then the DirEntry is ignored, if not then DirSize return a
// size of 0 alongside the error returned by Info.
func DirSize(fsys fs.FS) (int64, error) {
	dirents, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return 0, err
	}

	dirSize := int64(0)
	for _, dirent := range dirents {
		if !dirent.IsDir() {
			info, err := dirent.Info()
			// FIX: I do not know how to test this if statement
			// NOTE: From "io/fs/fs.go":
			// If the file has been removed or renamed since the directory read,
			// Info may return an error satisfying errors.Is(err, ErrNotExist).
			//
			// NOTE: If the file has been deleted, then simply ignore it and skip it.
			if errors.Is(err, fs.ErrNotExist) {
				continue
			} else if err != nil {
				return 0, err
			}
			dirSize += info.Size()
		}
	}
	return dirSize, nil
}

func DirSizeRMD(fsys fs.FS, maxdepth int) (int64, error) {
	fmt.Printf("maxdepth: %d\n", maxdepth)
	dirSize := int64(0)

	// FIX: Implement test cases for [fs.WalkDir] returning an error
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		// FIX: Implement test cases for err not nil in the WalkDirFunc

		var currentDepth int
		dirname := filepath.Dir(path)
		if dirname == "." {
			currentDepth = 0
		} else {
			currentDepth = strings.Count(dirname, string(os.PathSeparator)) + 1
		}
		fmt.Printf("path: %q - currentDepth: %d\n", path, currentDepth)

		if currentDepth > maxdepth {
			return fs.SkipDir
		}

		if !d.IsDir() {
			// FIX: Implement test cases for err not nil upon [fs.FileInfo.Info()]
			info, _ := d.Info()
			dirSize += info.Size()
		}

		return nil
	})

	return dirSize, nil
}

// DirSizeR compute the size of the root directory by summing it files' size.
// Then proceed to do it for all of the subdirectories.
func DirSizeR(fsys fs.FS) (int64, error) {
	dirSize := int64(0)
	// FIX: Implement test cases for [fs.WalkDir] returning an error
	_ = fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		// FIX: Implement test cases for err not nil in the WalkDirFunc

		if !d.IsDir() {
			// FIX: Implement test cases for err not nil upon [fs.FileInfo.Info()]
			info, _ := d.Info()
			dirSize += info.Size()
		}
		return nil
	})
	return dirSize, nil
}

// FileCountR count the files located at the root of the given file system. If
// the FS has subdirectories it will NOT traverse them. See [FileCountR] for a
// recursive function.
func FileCount(fsys fs.FS) (int, error) {
	dirents, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return 0, err
	}

	dirCounts := 0
	for _, dirent := range dirents {
		if !dirent.IsDir() {
			dirCounts += 1
		}
	}

	return dirCounts, nil
}

// FileCount count the files located in the given file system. If the FS has
// subdirectories it will traverse them. See [FileCount] for a none recursive
// function.
func FileCountR(fsys fs.FS) (int, error) {
	dirCounts := 0
	// FIX: Implement test cases for [fs.WalkDir] returning an error
	_ = fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		// FIX: Implement test cases for err not nil in the WalkDirFunc
		if !d.IsDir() {
			dirCounts += 1
		}

		return nil
	})
	return dirCounts, nil
}
