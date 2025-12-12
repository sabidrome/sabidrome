package main

import (
	"os"
	"fmt"

	"github.com/sabidrome/sabidrome/core"
)

func NewBookWatcher()

func main() {

    db := core.Database {
        Type: "sqlite3",
        Path: "./sabidrome.db",
    }

    command := os.Args[1]
    path    := os.Args[2]

    switch command {
        case "add":
            core.AddBook(&db, path)

        case "rm":
            fmt.Println("Oh no, rm not implemented")

        default:
            fmt.Println("Unknown command")
    }

}
