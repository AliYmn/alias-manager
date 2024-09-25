package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Main function to handle the loop and ask the user for input
func main() {
	reader := bufio.NewReader(os.Stdin)

	// Infinite loop to continuously ask user for actions
	for {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Add a new alias")
		fmt.Println("2. List all aliases")
		fmt.Println("3. Remove an alias")
		fmt.Println("4. Exit")

		// Get the user's choice
		fmt.Print("\nEnter choice (1-4): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// Perform action based on user's choice
		switch choice {
		case "1":
			// Add a new alias
			fmt.Print("Enter alias name: ")
			aliasName, _ := reader.ReadString('\n')
			aliasName = strings.TrimSpace(aliasName)

			fmt.Print("Enter alias command: ")
			aliasCommand, _ := reader.ReadString('\n')
			aliasCommand = strings.TrimSpace(aliasCommand)

			err := AddAlias(aliasName, aliasCommand)
			if err != nil {
				fmt.Println("Error adding alias:", err)
			} else {
				fmt.Println("Alias added successfully!")
			}

		case "2":
			// List all aliases
			ListAliases()
		case "3":
			// Remove an alias
			fmt.Print("Enter alias name to remove: ")
			aliasName, _ := reader.ReadString('\n')
			aliasName = strings.TrimSpace(aliasName)

			err := RemoveAlias(aliasName)
			if err != nil {
				fmt.Println("Error removing alias:", err)
			} else {
				fmt.Println("Alias removed successfully!")
			}

		case "4":
			// Exit the program
			fmt.Println("Exiting the program. Goodbye!")
			return

		default:
			// Handle invalid choices
			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
		}
	}
}
