package internals_test

import (
	"testing"
)

// NOTE: After refactoring I think FileCount is fully tested by the dinf/cmd_test package.
// This function is just a wrapper (and formatter to an io.Writer) of the [go-ding/dirs] package,
// which has its own tests.
func TestFileCount(t *testing.T) {
}
