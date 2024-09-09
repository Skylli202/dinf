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

func DirSize(w io.Writer, fsys fs.FS, opts DirSizeOpts) {
	var size int64
	if opts.Recursive {
		// FIX: This error needs to be handled properly, with testing.
		size, _ = dirs.DirSizeR(fsys)
	} else {
		// FIX: This error needs to be handled properly, with testing.
		size, _ = dirs.DirSize(fsys)
	}
	fmt.Fprintf(w, SizeFormat, size)
}
