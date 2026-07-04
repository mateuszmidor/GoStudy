package main

import (
	"database/sql"
	"log/slog"
	"os"
)

type SqliteDB struct {
	db *sql.DB
}

func NewSqliteDB() *SqliteDB {
	os.Remove("products.db")
	db, err := sql.Open("sqlite3", "products.db")
	if err != nil {
		slog.Error("NewSqliteDB: failure - " + err.Error())
		return nil
	}
	const statement = "create table Products(id integer primary key, name text, count integer);"
	_, err = db.Exec(statement)
	if err != nil {
		slog.Error("NewSqliteDB: failure - " + err.Error())
		return nil
	}
	return &SqliteDB{db: db}

}

func (db *SqliteDB) Create(p Product) (Product, error) {
	const statement = "insert into Products(name, count) values(:name, :count)"
	ret, err := db.db.Exec(statement, sql.Named("name", p.Name), sql.Named("count", p.Count))
	if err != nil {
		return Product{}, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return Product{}, err
	}
	p.ID = int(id)
	return p, nil
}

func (db *SqliteDB) GetAll() ([]Product, error) {
	result := []Product{}
	const statement = "select id, name, count from Products;"
	rows, err := db.db.Query(statement)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Count)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, nil
}

func (db *SqliteDB) Get(id int) (Product, error) {
	const query = "select id,  name, count from Products where id = :id"
	row := db.db.QueryRow(query, sql.Named("id", id))
	if row.Err() != nil {
		return Product{}, row.Err()
	}

	p := Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Count)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}
