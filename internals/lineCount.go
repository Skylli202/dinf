package internals

import (
	"io"
	"io/fs"
)

type LineCountOpts struct{}

func LineCount(writer io.Writer, fsys fs.FS, opts LineCountOpts) {
}
