package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

// Environment holds the operating system type and shell information.
type Environment struct {
	OSType string
	Shell  string
}

// GetEnvironment retrieves the current environment's OS and shell details.
// If the OS is not macOS (darwin) or Linux, it returns an error.
// Returns: error if the OS is unsupported, and the Environment struct.
func GetEnvironment() (Environment, error) {
	env := Environment{
		OSType: runtime.GOOS,
		Shell:  os.Getenv("SHELL"),
	}
	if env.OSType != "darwin" && env.OSType != "linux" {
		err := errors.New(fmt.Sprintf("Unsupported OS: %s", env.OSType))
		return env, err
	}
	return env, nil
}
