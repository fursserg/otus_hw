package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	from := "testdata/env"

	t.Run("reading envs from file", func(t *testing.T) {
		os.Setenv("FOO", "SHOULD_REPLACE")
		os.Setenv("EMPTY", "SHOULD_BE_EMPTY")
		os.Setenv("UNSET", "SHOULD_REMOVE")
		os.Setenv("HELLO", "SHOULD_REPLACE")

		vars, err := ReadDir(from)
		require.Nil(t, err)

		var expected Environment = map[string]EnvValue{}
		expected["BAR"] = EnvValue{Value: "bar", NeedRemove: false}
		expected["FOO"] = EnvValue{Value: "   foo\nwith new line", NeedRemove: false}
		expected["HELLO"] = EnvValue{Value: `"hello"`, NeedRemove: false}
		expected["UNSET"] = EnvValue{Value: "", NeedRemove: true}
		expected["EMPTY"] = EnvValue{Value: "", NeedRemove: false}

		require.Equal(t, expected, vars)
	})
}
