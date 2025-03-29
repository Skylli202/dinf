package dirs

import (
	"bufio"
	"io/fs"
	"os"
)

func LineCount(fsys fs.FS) (int, error) {
	dirents, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return 0, err
	}

	lc := 0
	for _, d := range dirents {
		// Ignore directories, they don't have line per say.
		if d.IsDir() {
			continue
		}

		f, err := os.Open(d.Name())
		// FIX: Implement test case for this error
		if err != nil {
			return 0, err
		}
		defer f.Close()

		s := bufio.NewScanner(f)
		s.Split(bufio.ScanLines)
		for s.Scan() {
			// From bufio.ScanLines docs: The returned line may be empty.
			// Might be worth to ignore empty line... Maybe create an option.
			lc += 1
		}
	}

	return lc, nil
}
