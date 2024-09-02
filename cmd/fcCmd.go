package cmd

import (
	"go-dinf/internals"
	"os"

	"github.com/spf13/cobra"
)

var fcCmd = &cobra.Command{
	Use:        "fc",
	Aliases:    []string{"file_count"},
	Example:    `go-dinf fc`,
	SuggestFor: []string{"cf"},
	Short:      "Count the file in the current directory.",
	// Long:       "",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		fsys := os.DirFS(wd)
		internals.FileCountCmd(os.Stdout, fsys)
	},
}

func init() {
	rootCmd.AddCommand(fcCmd)
}
