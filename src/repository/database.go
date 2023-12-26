package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	connectionString := os.Getenv("DATABASE")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Server Error!!", err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Server Error!!", err.Error())
	}
	fmt.Println("Connected Database MySQL!")
	return db
}