package main

import (
	"errors"
	"fmt"
	"os"

	"oaicheck/cmd"
)

func main() {
	err := cmd.NewRootCmd().Execute()
	if err == nil {
		return
	}
	if !errors.Is(err, cmd.ErrCheckFailed) {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(1)
}
