package adapters

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// using hostname and schema from wallet
func MysqlConnectHS(hostname, dbname string) *sql.DB {
	return MysqlConnect(USERNAME, PASSWORD, hostname, dbname)
}

func MysqlConnect(username, password, hostname, dbname string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname))
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	fmt.Println("Connected")
	return db
}

func ExecQuery(db *sql.DB, query string) *sql.Rows {
	rows, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Query executed")
	return rows
}
