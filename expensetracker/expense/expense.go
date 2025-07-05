package expense

import (
	"encoding/json"
	"os"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
}

const fileName = "expenses.json"

func LoadTasks() ([]Expense, error) {
	var expenses []Expense

	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return expenses, nil // Return empty slice if file does not exist
		}
		return nil, err // Return error if any other issue occurs
	}

	if len(file) == 0 {
		return expenses, nil // Return empty slice if file is empty
	}

	err = json.Unmarshal(file, &expenses)
	return expenses, err // Unmarshal JSON data into expenses slice
}

func SaveTasks(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err // Return error if JSON marshalling fails
	}
	return os.WriteFile(fileName, data, 0644) // Write JSON data to file
}
