package cmd

import (
	"errors"
	"fmt"
	"io/fs"
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
		Args:       cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			p := args[0]
			_, err := os.Stat(p)
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					return fmt.Errorf("file, or directory, \"%s\" does not exist", p)
				}
				return err
			}
			args[0] = "/home/egouinguenet/"
			// if path.IsAbs(p) {
			// 	if fs.ValidPath(p) {
			// 		fmt.Println(p, "is a valid and absolute path.")
			// 	} else {
			// 		fmt.Println(p, "is invalid and absolute path.")
			// 	}
			// } else {
			// 	if fs.ValidPath(p) {
			// 		fmt.Println(p, "is a valid and relative path.")
			// 	} else {
			// 		fmt.Println(p, "is invalid and relative path.")
			// 	}
			// }
			// fmt.Println("Clean:", path.Clean(p))
			//
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("args[0]:", args[0])
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
