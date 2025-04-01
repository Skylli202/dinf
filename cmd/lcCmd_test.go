package cmd_test

import (
	"os"
	"testing"

	"github.com/Skylli202/dinf/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewLcCmd(t *testing.T) {
	dirname, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatalf("Unable to `os.MkdirTemp(...)`. Unable to execute tests. err: %v", err)
	}
	defer os.RemoveAll(dirname)

	t.Run("LcCmd should have a raw flag of type Bool and default to false", func(t *testing.T) {
		lcCmd := cmd.NewLcCmd()
		b, err := lcCmd.Flags().GetBool("raw")
		lcCmd.Execute()
		require.NoError(t, err, "LcCmd should have a `raw`, `-r` flag.")
		require.False(t, b, "LcCmd's raw flag should be `false` by default.")
	})
	t.Run("--raw should pass the flag raw to true", func(t *testing.T) {
		lcCmd := cmd.NewLcCmd()
		lcCmd.SetArgs([]string{"--raw"})
		lcCmd.Execute()
		b, err := lcCmd.Flags().GetBool("raw")
		if assert.NoError(t, err, "LcCmd should have a raw flag") {
			assert.True(t, b, "raw flag should be true if '--raw' is specified in the args")
		}
	})
	t.Run("-r should pass the flag raw to true", func(t *testing.T) {
		lcCmd := cmd.NewLcCmd()
		lcCmd.SetArgs([]string{"-r"})
		lcCmd.Execute()
		b, err := lcCmd.Flags().GetBool("raw")
		if assert.NoError(t, err, "LcCmd should have a raw flag") {
			assert.True(t, b, "raw flag should be true if '-r' is specified in the args")
		}
	})
}
