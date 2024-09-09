package internals

import (
	"dinf/internals/dirs"
	"fmt"
	"io"
	"io/fs"
)

const SizeFormat = "Folder size is: %d bytes.\n"

type DirSizeOpts struct {
	Recursive bool
}

func DirSizeCmd(w io.Writer, fsys fs.FS, opts DirSizeOpts) {
	// FIX: This error needs to be handled properly, with testing.
	var size int64
	if opts.Recursive {
		size, _ = dirs.DirSizeR(fsys)
	} else {
		size, _ = dirs.DirSize(fsys)
	}
	// FIX: This error needs to be handled properly, with testing.
	fmt.Fprintf(w, SizeFormat, size)
}
