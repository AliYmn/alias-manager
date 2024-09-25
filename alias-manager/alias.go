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
	env, err := GetEnvironment()
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

// GetAliasFilePath returns the correct shell configuration file path based on the shell type.
func GetAliasFilePath() (string, error) {
	// Get the environment details.
	env, err := GetEnvironment()
	if err != nil {
		return "", err
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

	return filePath, nil
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

// AddAlias adds a new alias command to the respective shell file.
func AddAlias(aliasName, aliasCommand string) error {
	filePath, err := GetAliasFilePath()
	if err != nil {
		return err
	}

	// Open the file to append the new alias
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open alias file: %v", err)
	}
	defer file.Close()

	// Create the alias entry
	aliasEntry := fmt.Sprintf("alias %s='%s'\n", aliasName, aliasCommand)

	// Write the alias entry to the file
	if _, err := file.WriteString(aliasEntry); err != nil {
		return fmt.Errorf("could not write to alias file: %v", err)
	}

	fmt.Printf("Alias '%s' added successfully to %s\n", aliasName, filePath)
	return nil
}

// RemoveAlias removes an alias command from the respective shell file.
func RemoveAlias(aliasName string) error {
	filePath, err := GetAliasFilePath()
	if err != nil {
		return err
	}

	// Read the file content
	input, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read alias file: %v", err)
	}

	// Split the file content into lines
	lines := strings.Split(string(input), "\n")

	// Create a variable to track if the alias was found and removed
	aliasFound := false
	aliasPrefix := fmt.Sprintf("alias %s=", aliasName)

	// Iterate through the lines and remove the alias
	var updatedLines []string
	for _, line := range lines {
		if strings.HasPrefix(line, aliasPrefix) {
			aliasFound = true
			continue // Skip the line containing the alias to remove
		}
		updatedLines = append(updatedLines, line)
	}

	if !aliasFound {
		return fmt.Errorf("alias '%s' not found", aliasName)
	}

	// Write the updated content back to the file
	err = os.WriteFile(filePath, []byte(strings.Join(updatedLines, "\n")), 0644)
	if err != nil {
		return fmt.Errorf("could not write to alias file: %v", err)
	}

	fmt.Printf("Alias '%s' removed successfully from %s\n", aliasName, filePath)
	return nil
}
