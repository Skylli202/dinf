package dirs_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Skylli202/dinf/internals/dirs"
	"github.com/stretchr/testify/require"
)

const SINGLE_LINE = "package sample"

func TestLineCount(t *testing.T) {
	t.Run("Empty dir", func(t *testing.T) {
		dirname, err := os.MkdirTemp("", "")
		if err != nil {
			t.Fatal("Unable to `os.MkdirTemp(...)`. Unable to execute tests.")
		}
		defer os.RemoveAll(dirname)

		msg := "Empty file system."
		expectedLineCount := 0

		lc, err := dirs.LineCount(dirname)

		require.Nil(t, err, fmt.Sprintf("LineCount should return err equals to Nil, but it is not for test case: [%d] %s", 999, msg))
		// assert.Nil(t, errR, fmt.Sprintf("LineCountR should return err equals to Nil, but it is not for test case: [%d] %s", i, tc.msg))
		require.Equal(t, expectedLineCount, lc, fmt.Sprintf("Line count does not match.\nTest case: %d - %q.\nTest case's file system:\n%+v\n", 999, msg, dirname))
		// assert.Equal(t, tc.expectedSizeR, sizeR, fmt.Sprintf("Line count R  does not match.\nTest case: %d - %q.\nTest case's file system:\n%+v\n", i, tc.msg, tc.fs))
	})

	t.Run("Flat file system - 1 file", func(t *testing.T) {
		dirname, err := os.MkdirTemp("", "")
		if err != nil {
			t.Fatalf("Unable to `os.MkdirTemp(...)`. Unable to execute tests. err: %v", err)
		}
		defer os.RemoveAll(dirname)

		file, err := os.CreateTemp(dirname, "")
		if err != nil {
			t.Fatalf("Unable to `os.CreateTemp(\"%s\", \"\")`. Unable to execute tests. err: %v", dirname, err)
		}
		_, err = fmt.Fprint(file, SINGLE_LINE)
		if err != nil {
			t.Fatalf("Unable to write to the test file. Unable to execute tests. err: %v", err)
		}

		expectedLineCount := 1

		lc, err := dirs.LineCount(dirname)

		require.Nil(t, err, fmt.Sprintf("LineCount should not return nil. Returned error: %v", err))
		require.Equal(t, expectedLineCount, lc, "LineCount does not match the expected line count.")
	})

	t.Run("Flat file system - 4 file", func(t *testing.T) {
		dirname, err := os.MkdirTemp("", "")
		if err != nil {
			t.Fatalf("Unable to `os.MkdirTemp(...)`. Unable to execute tests. err: %v", err)
		}
		defer os.RemoveAll(dirname)

		expectedLineCount := 0
		for range 3 {
			file, err := os.CreateTemp(dirname, "")
			if err != nil {
				t.Fatalf("Unable to `os.CreateTemp(\"%s\", \"\")`. Unable to execute tests. err: %v", dirname, err)
			}
			_, err = fmt.Fprint(file, SINGLE_LINE)
			if err != nil {
				t.Fatalf("Unable to write to the test file. Unable to execute tests. err: %v", err)
			}
			expectedLineCount += 1
		}

		// // Create an empty file
		// _, err = os.CreateTemp(dirname, "")
		// if err != nil {
		// 	t.Fatalf("Unable to `os.CreateTemp(\"%s\", \"\")`. Unable to execute tests. err: %v", dirname, err)
		// }

		lc, err := dirs.LineCount(dirname)

		require.Nil(t, err, fmt.Sprintf("LineCount should not return nil. Returned error: %v", err))
		require.Equal(t, expectedLineCount, lc, "LineCount does not match the expected line count.")
	})
}
