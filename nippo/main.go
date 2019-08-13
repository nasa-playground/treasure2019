package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type Report struct {
	id        int
	title     string
	body      string
	createdAt time.Time
	comments  []Comment
}

type Comment struct {
	id        int
	reportID  int
	content   string
	createdAt time.Time
}

func main() {
	db, _ = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/treasure_app")

	command := os.Args[1]

	switch command {
	case "new":
		newReport()
	case "update":
		updateReport()
	case "delete":
		deleteReport()
	case "show":
		showReport()
	case "comment":
		commentReport()
	}
}

func newReport() {
	title := os.Args[2]
	body := os.Args[3]

	_, err := db.Exec("insert into reports(title, body, created_at) value(?, ?, NOW())", title, body)

	if err != nil {
		panic(err)
	}
}

func showReport() {
	title := os.Args[2]

	rows, err := db.Query("select * from reports where title = ?", title)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		report := new(Report)
		rows.Scan(&report.id, &report.title, &report.body, &report.createdAt)

		fmt.Printf("id: %v  title: %v created_at: %v\n%v", report.id, report.title, report.createdAt, report.body)

		commentRows, err := db.Query("select content from comments where report_id = ?", report.id)

		if err != nil {
			panic(err)
		}

		defer commentRows.Close()

		for commentRows.Next() {
			var content string
			rows.Scan(&content)

			fmt.Println(content)
		}
	}
}

func deleteReport() {
	title := os.Args[2]

	_, err := db.Exec("delete from reports where title = ?", title)

	if err != nil {
		panic(err)
	}
}

func updateReport() {
	title := os.Args[2]
	body := os.Args[3]

	// rows, err := db.Query("select id from reports where title = ?", title)
	// TODO 見つからないときにエラーにする

	db.Exec("update reports set body = ? where title = ?", body, title)
}

func commentReport() {
	id := os.Args[2]
	comment := os.Args[3]

	_, err := db.Exec("insert into comments(report_id, content, created_at) value(?, ?, NOW())", id, comment)

	if err != nil {
		panic(err)
	}
}
