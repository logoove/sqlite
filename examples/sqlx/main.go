package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/logoove/sqlite"
)

var db *sqlx.DB

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

func main() {
	db, _ = sqlx.Open("sqlite3", "./db.db")
	db.Exec(`CREATE TABLE users (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name" char(15) NOT NULL
);`)
	db.Exec(`insert into users(name) values(?)`, "李白")
	db.Exec(`insert into users(name) values(?)`, "白居易")
	var rows []User
	db.Select(&rows, "SELECT id,name FROM users ORDER BY id")
	fmt.Println(rows)
}
