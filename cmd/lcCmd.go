package cmd

import (
	"os"

	"github.com/Skylli202/dinf/internals"
	"github.com/spf13/cobra"
)

func NewLcCmd(dirname string) *cobra.Command {
	lcCmd := &cobra.Command{
		Use:        "lc",
		Aliases:    []string{"line_count", "lineCount", "LineCount", "Linecount", "linecount"},
		Example:    "dinf lc",
		SuggestFor: []string{"cl"},
		Short:      "Count lines in files of the current directory.",
		Args:       cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			w := cmd.OutOrStdout()

			raw, err := cmd.Flags().GetBool("raw")
			if err != nil {
				return err
			}
			opts := internals.LineCountOpts{
				Raw: raw,
			}
			internals.LineCount(w, dirname, opts)
			return nil
		},
	}

	lcCmd.Flags().BoolP(
		"raw",
		"r",
		false,
		"Output only the line count instead of a human friendly sentence.",
	)

	return lcCmd
}

func init() {
	wd, _ := os.Getwd()
	lcCmd := NewLcCmd(wd)
	rootCmd.AddCommand(lcCmd)
}
