package main

import (
    "database/sql"
    "fmt"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

const (
    hostname = "localhost"
    hostport = 5432
    username = "postgres"
    password = "mypw"
    databasename = "mydb"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, hostport, username, password, databasename)
	db, err := sql.Open("postgres", psqlconn)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Connected!")
}