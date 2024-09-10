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
		Aliases:    []string{"ds", "dirsize", "DirSize", "Dirsize", "dirSize"},
		Example:    `dinf size`,
		SuggestFor: []string{"szie", "siez", "sd"},
		Short:      "Compute the size (in bytes) of the files in the current directory.",
		// Long:       "",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			w := cmd.OutOrStdout()

			recursive, err := cmd.Flags().GetBool("recursive")
			if err != nil {
				return err
			}

			raw, err := cmd.Flags().GetBool("raw")
			if err != nil {
				return err
			}

			internals.DirSize(w, fsys, internals.DirSizeOpts{
				Recursive: recursive,
				Raw:       raw,
			})

			return nil
		},
	}

	sizeCmd.Flags().BoolP(
		"recursive",
		"R",
		false,
		"By default, size command look only at the direct child of the current working directory. If the recursive flag is specified, size will traverse all the file tree.",
	)

	sizeCmd.Flags().BoolP(
		"raw",
		"r",
		false,
		"",
	)

	return sizeCmd
}

func init() {
	wd, _ := os.Getwd()
	fsys := os.DirFS(wd)
	sizeCmd := NewSizeCmd(fsys)
	rootCmd.AddCommand(sizeCmd)
}
