package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go Mysql Tutorial")

	// Connect to database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golangtest")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully connected to mysql database")

	// Insert
	insert, err := db.Query("INSERT INTO table_test VALUES ('','Test 1')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	// Select Query
	results, err := db.Query("SELECT * FROM table_test")
	for results.Next() {
		var id int
		var name string
		results.Scan(&id, &name)
		fmt.Println(id, name)
	}
	defer results.Close()

	// Query Single Row
	var id int
	var name string
	// Execute Query
	db.QueryRow("SELECT id, name FROM table_test where id = ?", 3).Scan(&id, &name)
	fmt.Println(id,name)

}
