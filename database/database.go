package database

import (
	"Flaming_Sword_Ministry/models"
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

	log.Println("✅ Database connected successfully.")
}

func createUsersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fullname TEXT NOT NULL,
		phone TEXT NOT NULL UNIQUE,
		gender TEXT NOT NULL,
		password TEXT NOT NULL,
		role TEXT DEFAULT 'member',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Users table is ready.")
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
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Sermons table is ready.")
}
func GetAllUsers() ([]models.User, error) {

	rows, err := DB.Query(`
		SELECT id, fullname, phone, gender, role
		FROM users
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Phone,
			&user.Gender,
			&user.Role,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
