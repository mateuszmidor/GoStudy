package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// CREATE TABLE
const sqlCreateTables string = `
CREATE TABLE IF NOT EXISTS people(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(128) NOT NULL,
	age INT NOT NULL
);
`

// WRITE TABLE
// :name is sqlite3 syntax for palceholders
const sqlWriteData string = `
INSERT INTO 
people( name,  age) 
VALUES(:name, :age)
`

// READ TABLE
const sqlReadData string = `
SELECT name, age 
FROM people
`

func openConn(dbFileName string) *sql.DB {
	db, err := sql.Open("sqlite3", dbFileName)
	panicOnErr(err)

	return db
}

func createTables(db *sql.DB) {
	_, err := db.Exec(sqlCreateTables)
	panicOnErr(err)
}

func writeData(db *sql.DB) {
	_, err := db.Exec(sqlWriteData, sql.Named("name", "Andrzej"), sql.Named("age", 32))
	panicOnErr(err)

	_, err = db.Exec(sqlWriteData, sql.Named("name", "Jola"), sql.Named("age", 24))
	panicOnErr(err)
}

func readData(db *sql.DB) {
	rows, err := db.Query(sqlReadData)
	panicOnErr(err)

	for rows.Next() {
		var name string
		var age int
		rows.Scan(&name, &age)
		fmt.Println(name, age)
	}
	panicOnErr(rows.Close())
}

func panicOnErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	db := openConn("db.sqlite3")
	defer db.Close()

	createTables(db)
	writeData(db)
	readData(db)
}
