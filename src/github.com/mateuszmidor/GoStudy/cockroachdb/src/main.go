package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// CREATE TABLE
const sqlCreateTables string = `
CREATE TABLE IF NOT EXISTS people(
	id SERIAL PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	age INT NOT NULL
);
`

// WRITE TABLE
// $1, $2 is postgresql syntax for palceholders
const sqlWriteData string = `
INSERT INTO 
people( name,  age) 
VALUES($1, $2)
`

// READ TABLE
const sqlReadData string = `
SELECT name, age 
FROM people
`

func openConn(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	panicOnErr(err)

	return db
}

func createTables(db *sql.DB) {
	_, err := db.Exec(sqlCreateTables)
	panicOnErr(err)
}

func writeData(db *sql.DB) {
	_, err := db.Exec(sqlWriteData, "Andrzej", 32)
	panicOnErr(err)

	_, err = db.Exec(sqlWriteData, "Witek", 44)
	panicOnErr(err)

	_, err = db.Exec(sqlWriteData, "Jola", 24)
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
	db := openConn("postgresql://root@localhost:26257/defaultdb?sslmode=disable") // defaultdb exists in cluster once it has beed initialized
	defer db.Close()

	fmt.Println("Creating tables")
	createTables(db)
	fmt.Print("Done\n\n")

	fmt.Println("Adding records")
	writeData(db)
	fmt.Print("Done\n\n")

	fmt.Println("Reading records")
	readData(db)
	fmt.Print("Done\n\n")
}
