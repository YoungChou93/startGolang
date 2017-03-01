package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:qq329088816@tcp(localhost:3306)/photographypoint?charset=utf8")
	rows, _ := db.Query("select * from user")

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))

	values := make([]interface{}, len(columns))

	for j := range values {

		scanArgs[j] = &values[j]

	}

	record := make(map[string]string)

	for rows.Next() {

		//将行数据保存到record字典

		rows.Scan(scanArgs...)

		for i, col := range values {

			if col != nil {

				record[columns[i]] = string(col.([]byte))

			}

		}

	}

	fmt.Println(record)

}

func checkErr(err error) {

	if err != nil {

		fmt.Println(err)

		panic(err)

	}

}
