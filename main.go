package main

import (
	"os"
	"fmt"

	"github.com/sabidrome/sabidrome/core"
)

func main() {

    db := core.Database {
        Type: "sqlite3",
        Path: "./sabidrome.db",
    }

    command := os.Args[1]

    switch command {
        case "add":
            path    := os.Args[2]
            core.AddBook(&db, path)

        case "rm":
            fmt.Println("Oh no, rm not implemented")

        case "watch":
            core.FilesystemWatcher()

        default:
            fmt.Println("Unknown command")
    }

}
