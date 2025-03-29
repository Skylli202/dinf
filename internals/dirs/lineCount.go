package dirs

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
)

func LineCount(dirname string) (int, error) {
	fsys := os.DirFS(dirname)
	dirents, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return 0, err
	}

	lc := 0
	for i, d := range dirents {
		// Ignore directories, they don't have line per say.
		if d.IsDir() {
			continue
		}

		p := path.Join(dirname, d.Name())
		fmt.Printf("[%d] %s\n", i, p)
		f, err := os.Open(p)
		// FIX: Implement test case for this error
		if err != nil {
			return 0, err
		}
		defer f.Close()

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
	}

	return lc, nil
}
