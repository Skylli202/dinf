package internals

import (
	"fmt"
	"io"
	"io/fs"

	"github.com/Skylli202/dinf/internals/dirs"
)

const (
	FileCountFormat    = "Folder contains: %d files.\n"
	FileCountRawFormat = "%d\n"
)

type FileCountOpts struct {
	Recursive bool
	Raw       bool
}

func FileCount(w io.Writer, fsys fs.FS, opts FileCountOpts) {
	var fileCount int
	if opts.Recursive {
		fileCount, _ = dirs.FileCountR(fsys)
	} else {
		fileCount, _ = dirs.FileCount(fsys)
	}

	format := FileCountFormat
	if opts.Raw {
		format = FileCountRawFormat
	}

	fmt.Fprintf(w, format, fileCount)
}
