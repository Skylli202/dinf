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

		fmt.Printf("[%d] %s\n", i, d.Name())
		f, err := os.Open(path.Join(dirname, d.Name()))
		// FIX: Implement test case for this error
		if err != nil {
			return 0, err
		}
		defer f.Close()

		fmt.Println("Opening a new reader...")
		r := bufio.NewReader(f)
		fmt.Println("Reader open & ready!")
		for {
			_, err := r.ReadBytes('\n')
			fmt.Println("for...", lc, err)
			if err != nil && errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				return 0, err
			}
			lc += 1
		}
		// s := bufio.NewScanner(f)
		// s.Split(bufio.ScanLines)
		// for s.Scan() {
		// 	fmt.Println("Scan")
		// 	// From bufio.ScanLines docs: The returned line may be empty.
		// 	// Might be worth to ignore empty line... Maybe create an option.
		// 	lc += 1
		// }
	}

	return lc, nil
}
