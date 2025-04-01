package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/Skylli202/dinf/internals"
	"github.com/spf13/cobra"
)

func NewLcCmd() *cobra.Command {
	lcCmd := &cobra.Command{
		Use:        "lc",
		Aliases:    []string{"line_count", "lineCount", "LineCount", "Linecount", "linecount"},
		Example:    "dinf lc ./path/to/dirOrFile",
		SuggestFor: []string{"cl"},
		Short:      "Count lines in a directory, or a single file.",
		Args:       cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := os.Stat(args[0])
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					return fmt.Errorf("file, or directory, \"%s\" does not exist", args[0])
				}
				return err
			}
			w := cmd.OutOrStdout()

			raw, err := cmd.Flags().GetBool("raw")
			if err != nil {
				return err
			}
			opts := internals.LineCountOpts{
				Raw: raw,
			}

			internals.LineCount(w, args[0], opts)

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
	lcCmd := NewLcCmd()
	rootCmd.AddCommand(lcCmd)
}
