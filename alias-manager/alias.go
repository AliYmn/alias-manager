package main

import (
	"fmt"
	"os"
	"strings"
)

// AliasManager struct handles alias-related operations (add, remove, list).
type AliasManager struct {
	Env *Environment
}

// NewAliasManager creates a new AliasManager for managing alias operations.
func NewAliasManager(env *Environment) *AliasManager {
	return &AliasManager{Env: env}
}

// ListAliases lists all alias commands.
func (am *AliasManager) ListAliases() error {
	aliases, err := am.Env.ReadAliasFile()
	if err != nil {
		return err
	}

	if len(aliases) == 0 {
		fmt.Println("No alias commands found.")
		return nil
	}

	fmt.Println("Alias Commands:")
	for index, alias := range aliases {
		fmt.Println(index, "-" , alias)
	}
	return nil
}

// AddAlias adds a new alias command to the respective shell file.
func (am *AliasManager) AddAlias(aliasName, aliasCommand string) error {
	aliasEntry := fmt.Sprintf("alias %s='%s'\n", aliasName, aliasCommand)
	err := am.Env.WriteToAliasFile(aliasEntry)
	if err != nil {
		return err
	}

	fmt.Printf("Alias '%s' added successfully to %s\n", aliasName, am.Env.FilePath)
	return nil
}

// RemoveAlias removes an alias command from the respective shell file.
func (am *AliasManager) RemoveAlias(aliasName string) error {
	input, err := os.ReadFile(am.Env.FilePath)
	if err != nil {
		return fmt.Errorf("could not read alias file: %v", err)
	}

	lines := strings.Split(string(input), "\n")
	aliasFound := false
	aliasPrefix := fmt.Sprintf("alias %s=", aliasName)

	var updatedLines []string
	for _, line := range lines {
		if strings.HasPrefix(line, aliasPrefix) {
			aliasFound = true
			continue
		}
		updatedLines = append(updatedLines, line)
	}

	if !aliasFound {
		return fmt.Errorf("alias '%s' not found", aliasName)
	}

	err = os.WriteFile(am.Env.FilePath, []byte(strings.Join(updatedLines, "\n")), 0644)
	if err != nil {
		return fmt.Errorf("could not write to alias file: %v", err)
	}

	fmt.Printf("Alias '%s' removed successfully from %s\n", aliasName, am.Env.FilePath)
	return nil
}
