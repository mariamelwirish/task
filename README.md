# Task - CLI Task Manager

A simple and efficient command-line task manager built with Go. Manage your daily tasks directly from your terminal.


## Demo

```bash
$ task
task is a CLI Task Manager.

Usage:
  task [command]

Available Commands:
  add         Adds a task to your task list.
  do          Marks a task as complete.
  help        Help about any command
  list        Lists all of your tasks.

Flags:
  -h, --help   help for task

Use "task [command] --help" for more information about a command.

$ task add finish project
Added "finish project" to your task list.

$ task add clean dishes
Added "clean dishes" to your task list.

$ task list
You have the following tasks:
1. finish project
2. clean dishes

$ task do 1
Marked to mark "1" as completed.

$ task list
You have the following tasks:
1. clean dishes
```

*Note: Lines prefixed with `$` are commands typed into the terminal, other lines are program output.*

## Features

- **Add tasks** - Quickly add new tasks with `task add`
- **List tasks** - View all pending tasks with numbered IDs
- **Complete tasks** - Mark tasks as done with `task do`
- **Persistent storage** - Uses BoltDB for reliable local data storage
- **Cross-platform** - Works on Windows, macOS, and Linux

## Installation

### Prerequisites

You'll need Go installed on your system. If you don't have Go:
- **Download Go:** https://golang.org/dl/
- **Installation guide:** https://golang.org/doc/install

### Install Task

```bash
go install github.com/mariamelwirish/task@latest
```

This will install the `task` binary to your `$GOPATH/bin` directory. Make sure this directory is in your `$PATH`.

## Usage

### Adding Tasks

Add a new task to your TODO list:

```bash
task add Buy groceries
task add Finish project
```

### Listing Tasks

Display all your incomplete tasks:

```bash
task list
```

When you have no tasks:
```bash
$ task list
You have no tasks to complete! :D
```

### Completing Tasks

Mark tasks as complete by their number:

```bash
# Complete a single task
task do 1

# Complete multiple tasks at once
task do 1 3 5
```

### Getting Help

```bash
task --help           # Show main help
task add --help       # Help for add command
task list --help      # Help for list command
task do --help        # Help for do command
```

## How It Works

### Architecture

The CLI is built using:
- **[Cobra](https://github.com/spf13/cobra)** - Powerful CLI framework for Go
- **[BoltDB](https://github.com/boltdb/bolt)** - Embedded key/value database
- **[go-homedir](https://github.com/mitchellh/go-homedir)** - Cross-platform home directory detection

### Project Structure

```
task/
├── main.go           # Entry point and database initialization
├── cmd/              # Cobra command definitions
│   ├── root.go       # Root command setup
│   ├── add.go        # Add command implementation
│   ├── list.go       # List command implementation
│   └── do.go         # Do command implementation
└── db/               # Database operations
    └── tasks.go      # Task CRUD operations with BoltDB
```

### Data Storage

- Tasks are stored in a local BoltDB database
- Database location: `~/tasks.db` (in your home directory)
- Data persists across terminal sessions and system reboots

