package dirs

import (
	"errors"
	"io/fs"
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
