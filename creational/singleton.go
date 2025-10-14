package main

import (
    "fmt"
    "sync"
)

type Database struct {
    connection string
}

var instance *Database
var once sync.Once

func GetInstance() *Database {
    once.Do(func() {
        instance = &Database{connection: "Connected to database"}
    })
    return instance
}

func main() {
    db1 := GetInstance()
    db2 := GetInstance()

    fmt.Println(db1.connection)
    fmt.Println(db1 == db2) // true
}
