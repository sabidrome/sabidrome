package main

import (
    "os"
    "fmt"

    "github.com/sabidrome/sabidrome/core"
)

func main() {

	var (
        dialect = "sqlite3"
        path    = "./sabidrome.db"
    )

    db := core.InitDatabase(dialect, path)

    command := os.Args[1]

    switch command {
        case "add":
            path    := os.Args[2]
            core.AddBook(db, path)

        case "search":
            query   := os.Args[2]
            bookid := core.FindBook(db, query)
            if bookid == -1 {
                fmt.Printf("[Debug] Search yielded no results")
            } else {
                fmt.Printf("[Debug] Book id is '%d'\n", bookid)
            }


        case "watch":
            core.FilesystemWatcher()

        default:
            fmt.Println("Unknown command")
    }

}
