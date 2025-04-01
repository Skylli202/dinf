package dirs

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path"
)

// Count the number of line into a single file.
// A successful line count returns err == nil.
// If an error is encountered it returns lc == 0.
func lineCountFile(f *os.File) (lc int, err error) {
	r := bufio.NewReader(f)
	for {
		b, err := r.ReadBytes('\n')
		if err != nil && errors.Is(err, io.EOF) {
			if len(b) > 0 {
				lc += 1
			}
			break
		} else if err != nil {
			return 0, err
		}
		lc += 1
	}
	return lc, nil
}

// LineCount returns the number of line in the given path.
// LineCount works for file path or directory path.
// If an error occurs, LineCount return a line count of 0
// alongside the error.
func LineCount(dirname string) (int, error) {
	stat, err := os.Stat(dirname)
	if err != nil {
		return 0, nil
	}

	// dirname actually is a file, not a directory
	if !stat.IsDir() {
		f, err := os.Open(dirname)
		if err != nil {
			return 0, nil
		}
		defer f.Close()

		lc, err := lineCountFile(f)
		if err != nil {
			return 0, err
		}

		return lc, nil
	}

	// dirname is a directory
	dirents, err := os.ReadDir(dirname)
	if err != nil {
		return 0, err
	}

	lc := 0
	for _, d := range dirents {
		// Ignore directories, they don't have line per say.
		if d.IsDir() {
			continue
		}

		p := path.Join(dirname, d.Name())
		f, err := os.Open(p)
		// FIX: Implement test case for this error
		if err != nil {
			return 0, err
		}
		defer f.Close()

		fc, err := lineCountFile(f)
		if err != nil {
			return 0, nil
		}
		lc += fc
	}

	return lc, nil
}
