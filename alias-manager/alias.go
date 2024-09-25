package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadAliasFile reads the alias file based on the shell type.
func ReadAliasFile() ([]string, error) {
	// Get the environment details.
	err, env := GetEnvironment()
	if err != nil {
		return nil, err
	}

	// Determine the file path based on the shell.
	var filePath string
	if env.Shell == "/bin/zsh" {
		filePath = os.Getenv("HOME") + "/.zshrc"
	} else if env.Shell == "/bin/bash" {
		filePath = os.Getenv("HOME") + "/.bashrc"
	} else {
		filePath = os.Getenv("HOME") + "/.bash_aliases"
	}

	// Open the file and defer its closure.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file line by line and filter alias commands.
	var aliases []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "alias ") {
			aliases = append(aliases, line)
		}
	}

	// Check if there was a scanning error.
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return aliases, nil
}

// ListAliases prints out all alias commands from the shell config file.
func ListAliases() {
	aliases, err := ReadAliasFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(aliases) == 0 {
		fmt.Println("No alias commands found.")
		return
	}

	fmt.Println("Alias Commands:")
	for index, alias := range aliases {
		fmt.Println(index,":",alias)
	}
}
