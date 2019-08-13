package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/tutorial")
	if err != nil {
		panic(err)
	}

	fmt.Printf("db = %+v\n", db.Stats())
}
