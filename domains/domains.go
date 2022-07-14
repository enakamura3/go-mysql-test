package domains

import "database/sql"

type User struct {
	Id         int
	First_name string
	Last_name  string
	Birthday   string
	Created_at string
	Updated_at string
}

type UserSql struct {
	Id         sql.NullInt64
	First_name sql.NullString
	Last_name  sql.NullString
	Birthday   sql.NullString
	Created_at sql.NullString
	Updated_at sql.NullString
}
