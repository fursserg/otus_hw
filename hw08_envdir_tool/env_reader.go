package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("ReadDir error: %w", err)
	}

	vars := map[string]EnvValue{}

	for _, file := range files {
		fileName := file.Name()

		if !strings.Contains(fileName, "=") && !file.IsDir() {
			data, err := os.ReadFile(dir + "/" + fileName)
			if err != nil {
				return nil, fmt.Errorf("ReadDir read file %s error: %w", fileName, err)
			}

			_, existsEnv := os.LookupEnv(fileName)
			env := EnvValue{NeedRemove: existsEnv}

			if len(data) > 0 {
				data = bytes.SplitN(data, []byte("\n"), 2)[0]
				data = bytes.TrimRight(data, "\t ")
				data = bytes.ReplaceAll(data, []byte{0}, []byte("\n"))

				env.Value = string(data)
				env.NeedRemove = false
			}

			vars[fileName] = env
		}
	}

	// Place your code here
	return vars, nil
}
