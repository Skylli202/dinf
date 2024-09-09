package cmd

import (
	"dinf/internals"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

func NewSizeCmd(fsys fs.FS) *cobra.Command {
	sizeCmd := &cobra.Command{
		Use:        "size",
		Aliases:    []string{},
		Example:    `dinf size`,
		SuggestFor: []string{"szie", "siez"},
		Short:      "Compute the size (in bytes) of the files in the current directory.",
		// Long:       "",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			w := cmd.OutOrStdout()

			internals.DirSizeCmd(w, fsys)

			return nil
		},
	}

	return sizeCmd
}

func init() {
	wd, _ := os.Getwd()
	fsys := os.DirFS(wd)
	sizeCmd := NewSizeCmd(fsys)
	rootCmd.AddCommand(sizeCmd)
}
