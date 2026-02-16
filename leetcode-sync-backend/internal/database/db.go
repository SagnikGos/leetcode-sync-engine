package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "leetcode.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	problemsTable := `
	CREATE TABLE IF NOT EXISTS problems (
		slug TEXT PRIMARY KEY,
		title TEXT,
		difficulty TEXT,
		times_solved INTEGER,
		first_solved TEXT,
		last_solved TEXT
	);
	`

	submissionsTable := `
	CREATE TABLE IF NOT EXISTS submissions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		slug TEXT,
		code_hash TEXT,
		language TEXT,
		created_at TEXT,
		UNIQUE(slug, code_hash)
	);
	`

	_, err := DB.Exec(problemsTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(submissionsTable)
	if err != nil {
		log.Fatal(err)
	}
}
