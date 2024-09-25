
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	env, err := NewEnvironment()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	aliasManager := NewAliasManager(env)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Add a new alias")
		fmt.Println("2. List all aliases")
		fmt.Println("3. Remove an alias")
		fmt.Println("4. Exit")

		fmt.Print("\nEnter choice (1-4): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter alias name: ")
			aliasName, _ := reader.ReadString('\n')
			aliasName = strings.TrimSpace(aliasName)

			fmt.Print("Enter alias command: ")
			aliasCommand, _ := reader.ReadString('\n')
			aliasCommand = strings.TrimSpace(aliasCommand)

			err := aliasManager.AddAlias(aliasName, aliasCommand)
			if err != nil {
				fmt.Println("Error adding alias:", err)
			}

		case "2":
			err := aliasManager.ListAliases()
			if err != nil {
				fmt.Println("Error listing aliases:", err)
			}

		case "3":
			fmt.Print("Enter alias name to remove: ")
			aliasName, _ := reader.ReadString('\n')
			aliasName = strings.TrimSpace(aliasName)

			err := aliasManager.RemoveAlias(aliasName)
			if err != nil {
				fmt.Println("Error removing alias:", err)
			}

		case "4":
			fmt.Println("Exiting the program. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
		}
	}
}
