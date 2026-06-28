package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {

	var err error

	DB, err = sql.Open("sqlite3", "./church.db")
	if err != nil {
		log.Fatal(err)
	}

	createUsersTable()
	createSermonsTable()
	createAnnouncementsTable()

	log.Println("✅ Database connected successfully.")
}
func createUsersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fullname TEXT,
		phone TEXT UNIQUE,
		gender TEXT,
		password TEXT,
		role TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func createSermonsTable() {

	query := `
	CREATE TABLE IF NOT EXISTS sermons (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		bible_verse TEXT NOT NULL,
		scripture_references TEXT,
		content TEXT NOT NULL,
		category TEXT,
		date TEXT,
		created_by TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
func createAnnouncementsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS announcements (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		message TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
