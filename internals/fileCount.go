package internals

import (
	"fmt"
	"go-dinf/internals/dirs"
	"io"
	"io/fs"
)

func FileCountCmd(w io.Writer, fsys fs.FS) {
	fileCount, _ := dirs.FileCount(fsys)
	fmt.Fprintf(w, "Folder contains: %d files.\n", fileCount)
}

// NOTE: With the new way of testing the command itself I'm not too sure if that
// code is still useful, or not.
func FileCountRCmd(w io.Writer, fsys fs.FS) {
	fileCount, _ := dirs.FileCountR(fsys)
	fmt.Fprintf(w, "Folder contains: %d files.\n", fileCount)
}
