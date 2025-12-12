package main

import (
	"os"
	"fmt"

	"github.com/sabidrome/sabidrome/core"
)

func AddNewBook(db *core.Database, path string) {

    fmt.Println("[Debug] Begin add a new book")
    core.CreateDatabase(db)
    core.CreateDatabaseBookshelfTable(db)
    book_struct := core.GetBookMetadataFromPath(path)
    core.AddBookToDatabase(db, &book_struct)
    fmt.Println("[Debug] End add a new book")

}

func main() {

    db := core.Database {
        Type: "sqlite3",
        Path: "./sabidrome.db",
    }

    command := os.Args[1]
    path    := os.Args[2]

    switch command {
        case "add":
            AddNewBook(&db, path)

        case "rm":
            fmt.Println("Oh no, rm not implemented")

        default:
            fmt.Println("Unknown command")
    }

}
