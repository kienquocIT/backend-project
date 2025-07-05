# 📊 Expense Tracker CLI (Golang)

A simple command-line application to track your expenses using Go. You can add, list, delete, and calculate expenses (all or by specific month/year).

---

## ✅ Features

- ➕ Add a new expense with description
- 📃 List all expenses
- ❌ Delete expense by ID
- 💰 Calculate total expenses
- 📅 Calculate expenses by month/year (e.g., July 2025)

---

## 🚀 Run Instructions

### 1. Clone the repo
```bash
git clone https://github.com/your-username/expensetracker.git
cd expensetracker
```

### 2. Build or Run
```bash
go run main.go [flags]
```

---

## 🧭 Command Usage

### ➕ Add an expense
```bash
go run main.go -add="100:Dinner with friends"
```

### 📋 List all expenses
```bash
go run main.go -list
```

### ❌ Delete an expense by ID
```bash
go run main.go -delete=3
```

### 💰 Total of all expenses
```bash
go run main.go -sumAll
```

### 📅 Total expenses for a specific month/year
```bash
go run main.go -sum="7:2025"
```

---

## 🗂 Example Output

```
Expense added successfully! {ID: 1, Date: 2025-07-05, Amount: 100, Description: Dinner with friends}

Expenses:
ID: 1, Date: 2025-07-05, Amount: 100.00, Description: Dinner with friends
ID: 2, Date: 2025-07-05, Amount: 50.00, Description: Taxi

Total expenses: 150.00
Total expenses for 07/2025: 150.00
```

---

## 📁 Project Structure

```
expensetracker/
│
├── main.go                # CLI logic
├── expense/
│   └── expense.go         # Expense model + load/save functions
└── expenses.json          # Stored data (auto-created)
```

---

## 💾 Data Format

Expenses are saved in `expenses.json` in the following format:

```json
[
  {
    "ID": 1,
    "Date": "2025-07-05T14:23:00Z",
    "Amount": 100.0,
    "Description": "Dinner with friends"
  }
]
```

---

## 📦 Dependencies

- Golang `1.18+`
- Standard packages only (no external dependencies)

---

## 🧑‍💻 Author

Created by **Your Name** – feel free to fork and modify!

---

## 📜 License

MIT License