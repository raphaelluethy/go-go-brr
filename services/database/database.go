package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func GetDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		panic(err)
	}
	//_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT, password TEXT)")
	//if err != nil {
	//	return nil
	//}
	//_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ('Raphael Luethy', 'some@mail.ch','1234')")
	//if err != nil {
	//	return nil
	//}
	return db
}
