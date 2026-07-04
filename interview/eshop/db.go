package main

import (
	"database/sql"
	"log"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteStorage struct {
	conn *sql.DB
}

func NewSqliteStorage() *SqliteStorage {
	os.Remove("./eshop.db")
	db, err := sql.Open("sqlite3", "./eshop.db")
	if err != nil {
		log.Fatal(err)
	}

	const sql = "create table Items(id integer primary key, name varchar(128), count integer, price_cents integer)"
	_, err = db.Exec(sql)
	if err != nil {
		slog.Error(err.Error())
	}
	return &SqliteStorage{conn: db}
}

func (s *SqliteStorage) Close() {
	// s.conn.Close()
}

func (s *SqliteStorage) Add(item Item) Item {
	const statement = `
	insert into Items(name, count, price_cents)
	values(:name, :count, :price_cents)
	returning name, count, price_cents
	`
	ret, err := s.conn.Exec(statement, sql.Named("name", item.Name), sql.Named("count", item.Count), sql.Named("price_cents", item.PriceInCents))
	if err != nil {
		slog.Error(err.Error())
		return Item{}
	}
	id, err := ret.LastInsertId()
	if err != nil {
		slog.Error(err.Error())
		return Item{}
	}
	item.Id = int(id)
	return item
}

func (s *SqliteStorage) Delete(id IdType) {
	const stm = "delete from Items where id = 1"
	_, err := s.conn.Exec(stm, sql.Named("id", id))
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func (s *SqliteStorage) GetAll() []Item {
	const statement = "select id, name, count, price_cents from Items"
	rows, err := s.conn.Query(statement)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}
	defer rows.Close()
	result := []Item{}
	for rows.Next() {
		var item Item
		rows.Scan(&item.Id, &item.Name, &item.Count, &item.PriceInCents)
		result = append(result, item)
	}
	return result
}

func (s *SqliteStorage) Get(id IdType) *Item {
	const statement = "select id, name, count, price_cents from Items where id = :id"
	rows, err := s.conn.Query(statement, sql.Named("id", id))
	if err != nil {
		slog.Error(err.Error())
		return nil
	}
	defer rows.Close()
	if !rows.Next() {
		slog.Error("not found")
		return nil
	}
	var item Item
	err = rows.Scan(&item.Id, &item.Name, &item.Count, &item.PriceInCents)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}
	return &item
}
