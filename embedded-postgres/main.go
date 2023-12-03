package main

import (
	"database/sql"
	"fmt"
	"io"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
)

// CREATE TABLE
const sqlCreateTables string = `
CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    age INT NOT NULL
);
`

// WRITE TO TABLE
// $1, $2 is PostgreSQL syntax for palceholders
const sqlWriteData string = `
INSERT INTO 
people (name, age) 
VALUES ($1, $2);
`

// READ FROM TABLE
const sqlReadData string = `
SELECT name, age 
FROM people
`

// DELETE FROM TABLE
const sqlDeleteData string = `
DELETE FROM people
`

func main() {
	var noLogger io.Writer = nil // to avoid printing lots of postgre logs
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().Logger(noLogger))
	panicOnErr(postgres.Start())
	defer postgres.Stop()

	db := openConn("host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	createTables(db)
	writeData(db)
	readData(db)
	deleteData(db)
}

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
	_, err := db.Exec(sqlWriteData, "Andrzej", 32) // supply arguments for placeholders $1, $2
	panicOnErr(err)

	_, err = db.Exec(sqlWriteData, "Jola", 24) // supply arguments for placeholders $1, $2
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

func deleteData(db *sql.DB) {
	_, err := db.Exec(sqlDeleteData)
	panicOnErr(err)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
