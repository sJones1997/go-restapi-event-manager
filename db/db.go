package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var CONN *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")
	CONN = DB

	if err != nil {
		panic(err)
	}

	CONN.SetMaxOpenConns(10)
	CONN.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		name TEXT NOT NULL,
		    description TEXT NOT NULL,
		    location TEXT NOT NULL,
		    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		  	user_id INTEGER NOT NULL
		)`

	_, err := CONN.Exec(createEventsTable)

	if err != nil {
		panic(err)
	}

}
