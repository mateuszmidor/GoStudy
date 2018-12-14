package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "gopkg.in/goracle.v2"
)

func exitOnError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func connectDbOrExit() *sql.DB {
	// the below oracle db instance lives in AWS Cloud
	ConnString := "mateusz/SecretPass@(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=testdb.cyjughgpwadc.eu-central-1.rds.amazonaws.com)(PORT=1521))(CONNECT_DATA=(SERVER=DEDICATED)(SERVICE_NAME=orcl)))"
	db, err := sql.Open("goracle", ConnString)
	exitOnError(err)
	return db
}

func checkConnectionAliveOrExit(db *sql.DB) {
	err := db.Ping()
	exitOnError(err)
}

func querryDbOrExit(db *sql.DB, q string) *sql.Rows {
	rows, err := db.Query(q)
	exitOnError(err)
	return rows
}

func listColorsOrExit(db *sql.DB) {
	rows := querryDbOrExit(db, "select rownum, name from root.Colors")
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		exitOnError(err)
		fmt.Println(id, name)
	}
}

// First install goracle package:
// go get gopkg.in/goracle.v2,
// and drivers for OracleDB
// https://oracle.github.io/odpi/doc/installation.html#linux
func main() {
	db := connectDbOrExit()
	defer db.Close()

	checkConnectionAliveOrExit(db)

	listColorsOrExit(db)
}
