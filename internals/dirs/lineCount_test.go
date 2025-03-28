package dirs_test

import (
	"fmt"
	"testing"
	"testing/fstest"

	"github.com/Skylli202/dinf/internals/dirs"
	"github.com/stretchr/testify/assert"
)

func TestLineCount(t *testing.T) {
	cases := []struct {
		fs                 fstest.MapFS
		msg                string
		expectedLineCount  int
		expectedLineCountR int
	}{
		{
			fs:                 fstest.MapFS{},
			msg:                "Empty file system.",
			expectedLineCount:  0,
			expectedLineCountR: 0,
		},
	}

	for i, tc := range cases {
		lc, err := dirs.LineCount(tc.fs)
		// lcR, errR := dirs.LineCountR(fsys fs.FS)

		assert.Nil(t, err, fmt.Sprintf("LineCount should return err equals to Nil, but it is not for test case: [%d] %s", i, tc.msg))
		// assert.Nil(t, errR, fmt.Sprintf("LineCountR should return err equals to Nil, but it is not for test case: [%d] %s", i, tc.msg))
		assert.Equal(t, tc.expectedLineCount, lc, fmt.Sprintf("Line count does not match.\nTest case: %d - %q.\nTest case's file system:\n%+v\n", i, tc.msg, tc.fs))
		// assert.Equal(t, tc.expectedSizeR, sizeR, fmt.Sprintf("Line count R  does not match.\nTest case: %d - %q.\nTest case's file system:\n%+v\n", i, tc.msg, tc.fs))
	}
}
