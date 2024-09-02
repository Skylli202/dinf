package internals

import (
	"go-dinf/internals/dirs"
	"io"
	"io/fs"
	"text/template"
)

const DinfTemplate = "Folder size is: {{.Size}} bytes.\n"

func DirSizeCmd(w io.Writer, fsys fs.FS) {
	// FIX: This error needs to be handled properly, with testing.
	size, _ := dirs.DirSize(fsys)
	// FIX: This error needs to be handled properly, with testing.
	t, _ := template.New("DinfTemplate").Parse(DinfTemplate)
	t.Execute(w, struct{ Size int64 }{Size: size})
}
