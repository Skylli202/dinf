package internals

import (
	"fmt"
	"io"
	"io/fs"

	"github.com/Skylli202/dinf/internals/dirs"
)

const (
	SizeFormat    = "Folder size is: %d bytes.\n"
	SizeRawFormat = "%d\n"
)

type DirSizeOpts struct {
	Recursive bool
	Raw       bool
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

	format := SizeFormat
	if opts.Raw {
		format = SizeRawFormat
	}

	fmt.Fprintf(w, format, size)
}
