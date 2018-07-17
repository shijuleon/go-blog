package db

import (
	"database/sql"
   //	"fmt"
	_ "github.com/mattn/go-sqlite3"
	//"time"
)

// Read from json and init database

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}

func main(){
	const dbname = "blog.db" 
	db, err := sql.Open("sqlite3", dbname)
	checkErr(err)

	_, err = db.Exec(`CREATE TABLE users (
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) NULL,
        password VARCHAR(64) NULL,
        createdAt DATE NULL
	);`)
	checkErr(err)

	_, err = db.Exec(`CREATE TABLE posts (
        pid INTEGER PRIMARY KEY AUTOINCREMENT,
        title VARCHAR(64) NULL,
        body VARCHAR(64) NULL,
        time DATE NULL
	);`)
	checkErr(err)
}