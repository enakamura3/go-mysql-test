# go-mysql-test

- [go-mysql-test](#go-mysql-test)
  - [test](#test)
  - [read](#read)
  - [adapter](#adapter)
  - [mysql-docker](#mysql-docker)
    - [Create container](#create-container)
    - [Create database](#create-database)
    - [Create table](#create-table)

## test

Execute program:

```sh
go run main.go
```

Output:

```sh
Connected
Query executed
[]domains.User{domains.User{Id:1, First_name:"John", Last_name:"Doe", Birthday:"1990-01-22", Created_at:"2022-06-01 14:30:45", Updated_at:""}, domains.User{Id:2, First_name:"John", Last_name:"Smith", Birthday:"1980-02-23", Created_at:"2022-06-01 14:30:45", Updated_at:""}, domains.User{Id:3, First_name:"Jane", Last_name:"Smith", Birthday:"1981-03-24", Created_at:"2022-06-01 14:30:45", Updated_at:""}}
```

## read

To use golang mysql adapter to connect to a mysql database:

https://golangdocs.com/mysql-golang-crud-example

## adapter

```sh
go get -u https://github.com/go-sql-driver/mysql
```

## mysql-docker

### Create container

```sh
docker run --name mysql-test -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:latest
```

user: root
pass: my-secret-pw

### Create database

```sql
CREATE DATABASE `test`
```

### Create table

```sql
CREATE TABLE test.user (
	id int auto_increment NOT NULL,
	first_name varchar(100) NOT NULL,
	last_name varchar(100) NOT NULL,
	birthday DATE NULL,
	created_at DATETIME NULL,
	updated_at DATETIME NULL,
	PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
```

Insert data

```sql
INSERT INTO test.user (first_name, last_name, birthday, created_at, updated_at) VALUES('John', 'Doe', '1990-01-22', '2022-06-01 14:30:45', '2022-06-01 14:30:45');
INSERT INTO test.user (first_name, last_name, birthday, created_at, updated_at) VALUES('John', 'Smith', '1980-02-23', '2022-06-01 14:30:45', '2022-06-01 14:30:45');
INSERT INTO test.user (first_name, last_name, birthday, created_at, updated_at) VALUES('Jane', 'Smith', '1981-03-24', '2022-06-01 14:30:45', '2022-06-01 14:30:45');
```
