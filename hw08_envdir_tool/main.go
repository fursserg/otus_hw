package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("you need to pass at least two arguments, passed: %d\n", len(os.Args))
		return
	}

	vars, err := ReadDir(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	RunCmd(os.Args[2:], vars)
}
