package cmd

import (
	"dinf/internals"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

func NewFcCmd(fsys fs.FS) *cobra.Command {
	fcCmd := &cobra.Command{
		Use:        "fc",
		Aliases:    []string{"file_count", "fileCount", "FileCount", "Filecount", "filecount"},
		Example:    `dinf fc`,
		SuggestFor: []string{"cf"},
		Short:      "Count the file in the current directory.",
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

			opts := internals.FileCountOpts{
				Recursive: recursive,
				Raw:       raw,
			}

			internals.FileCount(w, fsys, opts)

			return nil
		},
	}

	fcCmd.Flags().BoolP(
		"recursive",
		"R",
		false,
		"By default, fileCount (fc) command look only at the direct child of the current working directory. If the recursive flag is specified, fileCount will travers all the file tree.",
	)

	fcCmd.Flags().BoolP(
		"raw",
		"r",
		false,
		"Output only the file count instead of a human friendly sentence.",
	)

	return fcCmd
}

func init() {
	wd, _ := os.Getwd()
	fsys := os.DirFS(wd)
	fcCmd := NewFcCmd(fsys)
	rootCmd.AddCommand(fcCmd)
}
