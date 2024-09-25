package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Environment struct handles OS and shell details, as well as shell config file management.
type Environment struct {
	OSType   string
	Shell    string
	FilePath string
}

// NewEnvironment initializes the Environment struct and determines the correct shell configuration file.
func NewEnvironment() (*Environment, error) {
	env := &Environment{
		OSType: runtime.GOOS,
		Shell:  os.Getenv("SHELL"),
	}

	// Determine the file path based on the shell type
	switch env.Shell {
	case "/bin/zsh":
		env.FilePath = os.Getenv("HOME") + "/.zshrc"
	case "/bin/bash":
		env.FilePath = os.Getenv("HOME") + "/.bashrc"
	default:
		env.FilePath = os.Getenv("HOME") + "/.bash_aliases"
	}

	// Validate OS support
	if env.OSType != "darwin" && env.OSType != "linux" {
		return nil, errors.New(fmt.Sprintf("Unsupported OS: %s", env.OSType))
	}

	return env, nil
}

// ReadAliasFile reads and returns all alias commands from the shell configuration file.
func (e *Environment) ReadAliasFile() ([]string, error) {
	file, err := os.Open(e.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var aliases []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "alias ") {
			aliases = append(aliases, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return aliases, nil
}

// WriteToAliasFile writes the provided content to the shell configuration file.
func (e *Environment) WriteToAliasFile(content string) error {
	file, err := os.OpenFile(e.FilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open alias file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("could not write to alias file: %v", err)
	}

	return nil
}
