# 📝 Task Tracker CLI in Go

A simple command-line application to manage your daily tasks using Go. You can add, list, mark done, and delete tasks. All data is stored locally in a `data.json` file.

---

## 📁 Project Structure

```
tasktracker/
│
├── main.go              # Entry point of the CLI
├── data.json            # JSON file storing task data (auto-created)
└── task/
    └── task.go          # Task struct and data persistence logic
```

---

## 🚀 Getting Started

### 1. Clone & navigate into the project

```bash
git clone <your-repo-url>
cd tasktracker
```

### 2. Initialize Go module

```bash
go mod init tasktracker
```

### 3. Run the program

Use the `go run` command with flags:

```bash
go run main.go -add="Learn Go CLI"
go run main.go -list
go run main.go -done=1
go run main.go -delete=1
```

---

## 🔧 CLI Options

| Flag         | Description                         | Example                                 |
|--------------|-------------------------------------|-----------------------------------------|
| `-add`       | Add a new task                      | `-add="Buy groceries"`                  |
| `-list`      | List all tasks                      | `-list`                                 |
| `-done`      | Mark a task as done by ID           | `-done=2`                               |
| `-delete`    | Delete a task by ID                 | `-delete=1`                             |

---

## 📦 Build

You can build the program as a standalone executable:

```bash
go build -o task
./task -list
```

---

## 🧠 Example Usage

```bash
# Add tasks
go run main.go -add="Write README"
go run main.go -add="Push to GitHub"

# List tasks
go run main.go -list

# Mark task ID 1 as done
go run main.go -done=1

# Delete task ID 2
go run main.go -delete=2
```

---

## 🗃️ Data Storage

- All tasks are stored in a local `data.json` file.
- If the file doesn’t exist, it will be created automatically.
- Each task includes:
  - ID
  - Title
  - Done status
  - CreatedAt timestamp

---

## ✨ Future Improvements

- [ ] Search tasks by keyword
- [ ] Categorize tasks (Work, Personal, etc.)
- [ ] Store data using SQLite or BoltDB
- [ ] Export tasks to Markdown or PDF

---

## 🧑‍💻 Author

- [Your Name Here]
- Go CLI made simple ✨

---