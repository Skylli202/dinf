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
