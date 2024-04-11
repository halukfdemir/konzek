package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    var err error
    db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/", handler)
    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", nil)
}
