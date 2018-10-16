package main

import (
	"os"
	"testing"
)

func TestCmdAdd(t *testing.T) {
	os.Args = []string{"ltpcli", "add", "--debug", "/proc/self/exe"}
	if err := rootCmd.Execute(); err != nil {
		t.Fatal("Error executing: ", err)
	}
}
