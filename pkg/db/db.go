package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/DBHackathon")
	if err != nil {
		fmt.Println("aqa")
		panic(err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("aqa ping")

		panic(err)
	}
	log.Println("Database Configured")
}
