package main

import (
    "database/sql"
    "log"
    "sync"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var mu sync.Mutex

func init() {
    var err error
    db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
}
