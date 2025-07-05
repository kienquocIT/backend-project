package main

import (
	"flag"
	"fmt"
	"os"
	"tasktracker/task"
	"time"
)

func main() {
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", 0, "Mark a task as done")
	del := flag.Int("delete", 0, "Delete a task")

	flag.Parse()

	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		newTask := task.Task{
			ID:        getNextID(tasks),
			Title:     *add,
			Done:      false,
			CreatedAt: time.Now(),
		}
		tasks = append(tasks, newTask)
		err := task.SaveTasks(tasks)
		if err == nil {
			fmt.Println("Task added:", newTask.Title)
		} else {
			fmt.Println("Error saving task:", err)
		}

	case *list:
		if len(tasks) == 0 {
			fmt.Println("No tasks available.")
		} else {
			for _, t := range tasks {
				status := " "
				if t.Done {
					status = "✔"
				}
				fmt.Printf("[%s] %d: %s (Created at: %s)\n",
					status, t.ID, t.Title, t.CreatedAt.Format("2006-01-02 15:04"))
			}
		}

	case *done > 0:
		for i := range tasks {
			if tasks[i].ID == *done {
				tasks[i].Done = true
				err := task.SaveTasks(tasks)
				if err == nil {
					fmt.Printf("Task %d marked as done.\n", *done)
				} else {
					fmt.Println("Error saving tasks:", err)
				}
				return
			}
		}
		fmt.Printf("Task with ID %d not found.\n", *done)

	case *del > 0:
		for i, t := range tasks {
			if t.ID == *del {
				tasks = append(tasks[:i], tasks[i+1:]...) // Remove task
				err := task.SaveTasks(tasks)
				if err == nil {
					fmt.Printf("Task %d deleted.\n", *del)
				} else {
					fmt.Println("Error deleting task:", err)
				}
				return
			}
		}
		fmt.Printf("Task with ID %d not found.\n", *del)

	default:
		fmt.Println("No valid command provided. Use -add, -list, -done, or -delete.")
	}
}

func getNextID(tasks []task.Task) int {
	maxId := 0
	for _, t := range tasks {
		if t.ID > maxId {
			maxId = t.ID
		}
	}
	return maxId + 1
}
