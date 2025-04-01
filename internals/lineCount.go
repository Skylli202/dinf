package internals

import (
	"fmt"
	"io"

	"github.com/Skylli202/dinf/internals/dirs"
)

const (
	LineCountFormat    = "Folder contains: %d lines.\n"
	LineCountRawFormat = "%d\n"
)

type LineCountOpts struct {
	Raw bool
}

func LineCount(writer io.Writer, dirname string, opts LineCountOpts) {
	lc, _ := dirs.LineCount(dirname)

	format := LineCountFormat
	if opts.Raw {
		format = LineCountRawFormat
	}

	fmt.Fprintf(writer, format, lc)
}
