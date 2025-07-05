package main

import (
	"expensetracker/expense"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Define command-line flags
	add := flag.String("add", "", "Add a new expense in the format 'amount:description'")
	list := flag.Bool("list", false, "List all expenses")
	del := flag.Int("delete", 0, "Delete an expense by ID")
	sum := flag.Bool("sumAll", false, "Calculate the total amount of expenses")
	sumPerMonth := flag.String("sum", "", "Calculate the total amount of expenses per month 'month:year'")

	flag.Parse()

	expenses, err := expense.LoadTasks()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		parts := strings.SplitN(*add, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format for -add. Use 'amount:description'.")
			os.Exit(1)
		}

		amount, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Invalid amount:", parts[0])
			os.Exit(1)
		}

		newExpense := expense.Expense{
			ID:          getNextID(expenses),
			Date:        time.Now(),
			Amount:      amount,
			Description: parts[1],
		}
		expenses = append(expenses, newExpense)
		if err := expense.SaveTasks(expenses); err == nil {
			fmt.Println("Expense added successfully!", newExpense)
		} else {
			fmt.Println("Error saving new expense:", err)
		}
	case *list:
		if len(expenses) == 0 {
			fmt.Println("No expenses found.")
		} else {
			fmt.Println("Expenses:")
			for _, exp := range expenses {
				fmt.Printf("ID: %d, Date: %s, Amount: %.2f, Description: %s\n",
					exp.ID, exp.Date.Format("2006-01-02"), exp.Amount, exp.Description)
			}
		}
	case *sum:
		total := 0.0
		for _, exp := range expenses {
			total += exp.Amount
		}
		fmt.Printf("Total expenses: %.2f\n", total)
	case *sumPerMonth != "":
		monthYear := strings.Split(*sumPerMonth, ":")
		if len(monthYear) != 2 {
			fmt.Println("Invalid format for -sum. Use 'month:year'.")
			os.Exit(1)
		}
		month, err := strconv.Atoi(monthYear[0])
		if err != nil || month < 1 || month > 12 {
			fmt.Println("Invalid month. Please provide a month between 1 and 12.")
			os.Exit(1)
		}
		year, err := strconv.Atoi(monthYear[1])
		if err != nil || year < 0 {
			fmt.Println("Invalid year. Please provide a valid year.")
			os.Exit(1)
		}
		totalPerMonth := 0.0
		for _, exp := range expenses {
			if exp.Date.Month() == time.Month(month) && exp.Date.Year() == year {
				totalPerMonth += exp.Amount
			}
		}
		fmt.Printf("Total expenses for %02d/%d: %.2f\n", month, year, totalPerMonth)
	case *del > 0:
		for i, exp := range expenses {
			if exp.ID == *del {
				expenses = append(expenses[:i], expenses[i+1:]...)
				err := expense.SaveTasks(expenses)
				if err == nil {
					fmt.Println("Expense deleted successfully!")
				} else {
					fmt.Println("Error saving changes after deletion:", err)
				}
				return
			}
		}
		fmt.Printf("Expense with ID %d not found.\n", *del)
	default:
		fmt.Println("No valid command provided. Use -add, -list, or -delete.")
	}
}

func getNextID(expenses []expense.Expense) int {
	if len(expenses) == 0 {
		return 1 // Start with ID 1 if no expenses exist
	}
	maxID := expenses[0].ID
	for _, exp := range expenses {
		if exp.ID > maxID {
			maxID = exp.ID
		}
	}
	return maxID + 1 // Return the next ID
}
