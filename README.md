
# Alias Manager

This is a simple Go-based Alias Manager that allows you to:

- Add new alias commands
- List all alias commands
- Remove alias commands

## How to use

1. Clone the repository and navigate to the project directory.

```bash
git clone https://github.com/AliYmn/alias-manager.git
cd alias-manager
```

2. Build the Go program.

```bash
go build -o alias-manager
```

3. Run the program to manage your alias commands.

```bash
./alias-manager
```

## Features

- **Add Alias**: Adds a new alias to the shell configuration file (`.zshrc`, `.bashrc`, or `.bash_aliases`).
- **List Aliases**: Lists all alias commands in your shell configuration file.
- **Remove Alias**: Removes an alias from the shell configuration file.

## File Structure

- `main.go`: Main entry point of the program that handles user interactions.
- `alias.go`: Contains functions for managing alias commands (Add, List, Remove).
- `environment.go`: Contains functions for detecting the operating system and shell type.

## Requirements

- Go 1.18 or later
- MacOS or Linux operating system

## How It Works

The program checks the current shell (`zsh`, `bash`, or others) and interacts with the corresponding shell configuration file to manage alias commands.

- **Shell Detection**: The program uses the `SHELL` environment variable to determine the current shell.
- **Alias Commands**: Alias commands are added, listed, or removed by modifying the shell configuration files.

