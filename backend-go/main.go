package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "enconding/json"
    "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
    var err error
    // The database will be created as a file named 'help_assistant.db' in the backend directory
    db, err = sql.Open("sqlite3", "./help_assistant.db")
    if err != nil {
        log.Fatal("Error opening SQLite database:", err)
    }

    // Define the SQL to create the table
    createTableSQL := `CREATE TABLE IF NOT EXISTS query_log (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        query_text TEXT NOT NULL,
        response_text TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

    // Execute the SQL command
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal("Error creating table:", err)
    }

    fmt.Println("Successfully connected to SQLite and ensured table exists!")
}

func logQuery(query, response string) error {
    // SQLite uses '?' as the placeholder, which is automatically handled by the driver
    // The columns are identical to the PostgreSQL plan

        _, err := db.Exec("INSERT INTO query_log (query_text, response_text) VALUES (?, ?)", query, response)
        if err != nil {
            return fmt.Errorf("error logging query: %w", err)
        }
        return nil
}