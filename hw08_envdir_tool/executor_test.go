package main

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("checking exit status code", func(t *testing.T) {
		code := RunCmd([]string{"sleep"}, map[string]EnvValue{})
		require.Equal(t, 1, code)

		code = RunCmd([]string{"sleep", "1"}, map[string]EnvValue{})
		require.Equal(t, 0, code)
	})

	t.Run("applying environment variables", func(t *testing.T) {
		os.Setenv("FOO", "SHOULD_REPLACE")
		os.Setenv("EMPTY", "SHOULD_BE_EMPTY")
		os.Setenv("UNSET", "SHOULD_REMOVE")
		os.Setenv("HELLO", "SHOULD_REPLACE")
		os.Setenv("ADDED", "from original")

		origStdOut := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		var vars Environment = map[string]EnvValue{}
		cmd := []string{"testdata/echo.sh", "testargs"}

		vars["BAR"] = EnvValue{Value: "bar", NeedRemove: false}
		vars["FOO"] = EnvValue{Value: "   foo\nwith new line", NeedRemove: false}
		vars["HELLO"] = EnvValue{Value: `"hello"`, NeedRemove: false}
		vars["UNSET"] = EnvValue{Value: "", NeedRemove: true}
		vars["EMPTY"] = EnvValue{Value: "", NeedRemove: false}
		vars["ADDED"] = EnvValue{Value: "from original", NeedRemove: false}

		expectedOut := fmt.Sprintf(
			"HELLO is (%s)\nBAR is (%s)\nFOO is (%s)\nUNSET is (%s)\nADDED is (%s)\nEMPTY is (%s)\narguments are %s\n",
			vars["HELLO"].Value,
			vars["BAR"].Value,
			vars["FOO"].Value,
			vars["UNSET"].Value,
			vars["ADDED"].Value,
			vars["EMPTY"].Value,
			cmd[1],
		)

		RunCmd(cmd, vars)

		w.Close()
		actualOut, _ := io.ReadAll(r)

		os.Stdout = origStdOut

		require.Equal(t, expectedOut, string(actualOut))

		_, ok := os.LookupEnv("UNSET")
		require.False(t, ok)
	})
}
