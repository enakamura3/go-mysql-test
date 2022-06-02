package main

import (
	"fmt"

	"github.com/enakamura3/go-mysql-test/adapters"
	"github.com/enakamura3/go-mysql-test/domains"
)

func main() {
	dbConnection := adapters.MysqlConnect()
	defer dbConnection.Close()

	query := "select id, first_name, last_name, birthday, created_at from test.user;"

	rows := adapters.ExecQuery(dbConnection, query)
	defer rows.Close()

	users := []domains.User{}

	for rows.Next() {
		u := domains.User{}
		err := rows.Scan(&u.Id, &u.First_name, &u.Last_name, &u.Birthday, &u.Created_at)
		if err != nil {
			panic(err)
		}
		users = append(users, u)
	}
	fmt.Printf("%#v\n", users)
}
