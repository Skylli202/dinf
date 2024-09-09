package internals

import (
	"dinf/internals/dirs"
	"fmt"
	"io"
	"io/fs"
)

const SizeFormat = "Folder size is: %d bytes.\n"

func DirSizeCmd(w io.Writer, fsys fs.FS) {
	// FIX: This error needs to be handled properly, with testing.
	size, _ := dirs.DirSize(fsys)
	// FIX: This error needs to be handled properly, with testing.
	fmt.Fprintf(w, SizeFormat, size)
}
