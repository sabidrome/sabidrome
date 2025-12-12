package main

import (
	"os"
	"fmt"

	"github.com/sabidrome/sabidrome/core"
)

func main() {
    command := os.Args[1]

    if command == "add" {
        path := os.Args[2]
        fmt.Println("Hello, World!")

	databaseObject := core.Database {
		Type: "sqlite3",
		Path: "./sabidrome.db",
	}

	core.CreateDatabase(&databaseObject)
	core.CreateDatabaseBookshelfTable(&databaseObject)
	book_struct := core.GetBookMetadataFromPath(path)
	core.AddBookToDatabase(&databaseObject, &book_struct)


    } else if command == "rm" {
        fmt.Println("Oh no")

    } else {
        fmt.Println("Ah ha")
    }
}
