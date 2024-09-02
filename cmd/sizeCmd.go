package cmd

import (
	"go-dinf/internals"
	"os"

	"github.com/spf13/cobra"
)

var sizeCmd = &cobra.Command{
	Use:        "size",
	Aliases:    []string{},
	Example:    `go-dinf size`,
	SuggestFor: []string{"szie", "siez"},
	Short:      "Compute the size (in bytes) of the files in the current directory.",
	// Long:       "",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		fsys := os.DirFS(wd)
		internals.DirSizeCmd(os.Stdout, fsys)
	},
}

func init() {
	rootCmd.AddCommand(sizeCmd)
}
