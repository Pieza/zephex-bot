package db

import (
	"database/sql"
	// SQLite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// Meme type
type Meme struct {
	ID      int
	Command string
	URL     string
}

// Connect creates a database instance
func Connect() (*sql.DB, error) {
	return sql.Open("sqlite3", "./zephex.db")
}

// SetUpDB configures the database
func SetUpDB() {
	database, _ := Connect()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS memes( id INTEGER PRIMARY KEY,  command TEXT, url TEXT)")
	statement.Exec()
	database.Close()
}

// InsertCommand inserts a new meme command
func InsertCommand(command string, url string) {
	database, _ := Connect()
	statement, _ := database.Prepare("INSERT INTO memes (command, url) VALUES (?, ?)")
	statement.Exec(command, url)
	database.Close()
}

// SelectAllCommands selects all meme commands
func SelectAllCommands() []Meme {
	database, _ := Connect()
	rows, _ := database.Query("SELECT id, command, url FROM memes")
	var items = []Meme{}
	for rows.Next() {
		item := Meme{}
		rows.Scan(&item.ID, &item.Command, &item.URL)
		items = append(items, item)
	}
	database.Close()
	return items
}
