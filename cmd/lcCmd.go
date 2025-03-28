package cmd

import (
	"io/fs"
	"os"

	"github.com/Skylli202/dinf/internals"
	"github.com/spf13/cobra"
)

func NewLcCmd(fsys fs.FS) *cobra.Command {
	lcCmd := &cobra.Command{
		Use:        "lc",
		Aliases:    []string{"line_count", "lineCount", "LineCount", "Linecount", "linecount"},
		Example:    "dinf lc",
		SuggestFor: []string{"cl"},
		Short:      "Count lines in files of the current directory.",
		Args:       cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			w := cmd.OutOrStdout()
			opts := internals.LineCountOpts{}
			internals.LineCount(w, fsys, opts)
			return nil
		},
	}

	return lcCmd
}

func init() {
	wd, _ := os.Getwd()
	fsys := os.DirFS(wd)
	lcCmd := NewLcCmd(fsys)
	rootCmd.AddCommand(lcCmd)
}
