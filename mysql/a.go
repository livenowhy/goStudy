package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// mysql://root:root123@index.boxlinker.com:3306/registry
	db, err := sql.Open("mysql", "root:root123@index.boxlinker.com:3306/registry")
	if err != nil {
		fmt.Println("ssssss")
		fmt.Println(err)
	}
	stmt, err := db.Prepare("create table if not exists dev(id int UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,uid varchar(64),did varchar(64),name varchar(64),qid varchar(64),status char DEFAULT 'u')")
	if stmt != nil {
		stmt.Exec()
		stmt.Close()
	}
	fmt.Println("sssssss")

}
