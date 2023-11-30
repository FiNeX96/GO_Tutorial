package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Task struct to represent individual tasks
type Task struct {
	ID          int
	Description string
	Done        bool
	Timestamp   int64
}

type TaskManager struct {
	db *sql.DB
}


func main() {
		// Initialize the task manager
		taskManager, err := NewTaskManager("tasks.sqlite3")
		if err != nil {
			fmt.Println("Error initializing task manager:", err)
			return
		}

	// Define command-line flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addDesc := addCmd.String("description", "", "Task description")

	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	completeID := completeCmd.Int("id", 0, "Task ID to mark as completed")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.Int("id", 0, "Task ID to delete")

	// Set the command-line arguments for parsing
	flag.CommandLine.Parse(os.Args[1:])

	// Parse the command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: add, complete, list, or delete")
		return
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		taskManager.AddTask(*addDesc)

	case "complete":
		completeCmd.Parse(os.Args[2:])
		taskManager.CompleteTask(*completeID)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		taskManager.DeleteTask(*deleteID)
	case "list":
		taskManager.ListTasks()
	default:
		fmt.Println("Invalid command. Please use add, complete, list, or delete.")
	}

}

// NewTaskManager is a standalone function, not a method
func NewTaskManager(dbPath string) (*TaskManager, error) {
	DB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT,
			done BOOLEAN,
			timestamp INTEGER
		)
	`)
	if err != nil {
		return nil, err
	}

	return &TaskManager{db: DB}, nil
}

func (m *TaskManager) AddTask(description string) error {
	_, err := m.db.Exec(`
		INSERT INTO tasks (description, done, timestamp) VALUES (?, ?, ?)
	`, description, false, time.Now().UnixNano())
	return err
}

func (m *TaskManager) CompleteTask(id int) error {
	_, err := m.db.Exec(`
		UPDATE tasks SET done = ? WHERE id = ?
	`, 1, id)
	return err
}

func (m *TaskManager) DeleteTask(id int) error {
	_, err := m.db.Exec(`
		DELETE FROM tasks WHERE id = ?
	`, id)
	return err
}

func (m *TaskManager) ListTasks() error {

	fmt.Println("xdd")
	return nil
}







