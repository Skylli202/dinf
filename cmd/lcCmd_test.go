package cmd_test

import (
	"os"
	"testing"

	"github.com/Skylli202/dinf/cmd"
	"github.com/stretchr/testify/assert"
)

func Test_NewLcCmd(t *testing.T) {
	dirname, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatalf("Unable to `os.MkdirTemp(...)`. Unable to execute tests. err: %v", err)
	}
	defer os.RemoveAll(dirname)

	t.Run("LcCmd should have a raw flag of type Bool", func(t *testing.T) {
		lcCmd := cmd.NewLcCmd(dirname)
		_, err := lcCmd.Flags().GetBool("raw")
		lcCmd.Execute()
		assert.NoError(t, err)
	})
	// t.Run("By default, raw flag should be false", func(t *testing.T) {
	// 	fcCmd := cmd.NewFcCmd(emptyFS)
	// 	b, err := fcCmd.Flags().GetBool("raw")
	// 	fcCmd.Execute()
	// 	if assert.NoError(t, err, "FcCmd should have a raw flag") {
	// 		assert.False(t, b, "raw flag should be flase by default")
	// 	}
	// })
	// t.Run("--raw should pass the flag raw to true", func(t *testing.T) {
	// 	fcCmd := cmd.NewFcCmd(emptyFS)
	// 	fcCmd.SetArgs([]string{"--raw"})
	// 	fcCmd.Execute()
	// 	b, err := fcCmd.Flags().GetBool("raw")
	// 	if assert.NoError(t, err, "FcCmd should have a raw flag") {
	// 		assert.True(t, b, "raw flag should be true if '--raw' is specified in the args")
	// 	}
	// })
	// t.Run("-r should pass the flag raw to true", func(t *testing.T) {
	// 	fcCmd := cmd.NewFcCmd(emptyFS)
	// 	fcCmd.SetArgs([]string{"-r"})
	// 	fcCmd.Execute()
	// 	b, err := fcCmd.Flags().GetBool("raw")
	// 	if assert.NoError(t, err, "FcCmd should have a raw flag") {
	// 		assert.True(t, b, "raw flag should be true if '-r' is specified in the args")
	// 	}
	// })
}
