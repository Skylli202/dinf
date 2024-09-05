package cmd

import (
	"go-dinf/internals"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

func NewFcCmd(fsys fs.FS) *cobra.Command {
	fcCmd := &cobra.Command{
		Use:        "fc",
		Aliases:    []string{"file_count", "fileCount", "FileCount", "Filecount", "filecount"},
		Example:    `go-dinf fc`,
		SuggestFor: []string{"cf"},
		Short:      "Count the file in the current directory.",
		// Long:       "",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			R, err := cmd.Flags().GetBool("recursive")
			if err != nil {
				return err
			}

			if R {
				internals.FileCountRCmd(cmd.OutOrStdout(), fsys)
			} else {
				internals.FileCountCmd(cmd.OutOrStdout(), fsys)
			}

			return nil
		},
	}

	fcCmd.Flags().BoolP(
		"recursive",
		"R",
		false,
		"By default, fileCount (fc) command look only at the direct child of the current working directory. If the recursive flag is specified, fileCount will travers all the file tree.",
	)

	return fcCmd
}

func init() {
	wd, _ := os.Getwd()
	fsys := os.DirFS(wd)
	fcCmd := NewFcCmd(fsys)
	rootCmd.AddCommand(fcCmd)
}
