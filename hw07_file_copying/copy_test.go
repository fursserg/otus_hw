package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	from := "testdata/input.txt"
	to := "testdata/copy_of_input.txt"

	fileOrig, err := os.ReadFile(from)
	if err != nil {
		t.Fatalf("couldn't read file: %v", err)
	}

	t.Run("copying full file", func(t *testing.T) {
		Copy(from, to, 0, 0)

		fileCopy, err := os.ReadFile(to)
		if err != nil {
			t.Fatalf("couldn't read file: %v", err)
		}

		require.Equal(t, fileOrig, fileCopy)

		err = os.Remove(to)
		if err != nil {
			fmt.Printf("removing copied file error: %v\n", err)
		}
	})

	t.Run("copying file with offset", func(t *testing.T) {
		offset := int64(150)
		Copy(from, to, offset, 0)

		fileCopy, err := os.ReadFile(to)
		if err != nil {
			t.Fatalf("couldn't read file: %v", err)
		}

		require.Equal(t, fileOrig[offset:], fileCopy)

		err = os.Remove(to)
		if err != nil {
			fmt.Printf("removing copied file error: %v\n", err)
		}
	})

	t.Run("copying file with limit", func(t *testing.T) {
		limit := int64(150)
		Copy(from, to, 0, limit)

		fileCopy, err := os.ReadFile(to)
		if err != nil {
			t.Fatalf("couldn't read file: %v", err)
		}

		require.Equal(t, fileOrig[:limit], fileCopy)

		err = os.Remove(to)
		if err != nil {
			fmt.Printf("removing copied file error: %v\n", err)
		}
	})

	t.Run("copying file with offset and limit", func(t *testing.T) {
		limit := int64(150)
		offset := int64(150)
		Copy(from, to, offset, limit)

		fileCopy, err := os.ReadFile(to)
		if err != nil {
			t.Fatalf("couldn't read file: %v", err)
		}

		require.Equal(t, fileOrig[offset:offset+limit], fileCopy)

		err = os.Remove(to)
		if err != nil {
			fmt.Printf("removing copied file error: %s\n", err)
		}
	})
}
