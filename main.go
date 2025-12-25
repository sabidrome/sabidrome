package main

import (
    // "fmt"
    "database/sql"
    "log/slog"
    "os"

    _ "github.com/mattn/go-sqlite3"

    "github.com/sabidrome/sabidrome/core"
    "github.com/sabidrome/sabidrome/db"
    "github.com/sabidrome/sabidrome/files"
)

func test_basic_funcs_db(session_db *sql.DB) {

    db.BooksList(session_db)

    // Test Add Book
    b := &core.Book{0, "Dummy Title", "Dummy Creator", "Dummy Publisher", 12345, "/home/dummy/dummy.epub"}
    b.Id = db.AddBook(session_db, b)
    db.BooksList(session_db)

    // Test Update Book
    b.Path = "/home/smart/smart.epub"
    db.UpdateBookPath(session_db, b.Id, b.Path)
    db.BooksList(session_db)

    // Test Remove Book
    db.RemoveBook(session_db, b.Id)
    db.BooksList(session_db)

}

func test_basic_funcs_fs(dir string) {

    files.ListDir(dir)
    files.ListBasePath(dir)

}

func main() {

    slog.SetLogLoggerLevel(slog.LevelDebug)

    session_db := db.ConnectOrCreateDatabase()

    command := os.Args[1]

    switch command  {
        case "basic-test-db":
            test_basic_funcs_db(session_db)

        case "basic-test-fs":
            test_basic_funcs_fs(os.Args[2])

        default:
            os.Exit(255)
    }



}
