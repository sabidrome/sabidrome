package main

import (
    // "fmt"
    "log/slog"

    _ "github.com/mattn/go-sqlite3"
    "github.com/sabidrome/sabidrome/db"
    "github.com/sabidrome/sabidrome/core"
)




func main() {

    slog.SetLogLoggerLevel(slog.LevelDebug)

    session_db := db.ConnectOrCreateDatabase()

    db.BooksList(session_db)

    // Test Add Book
    b := &Book{0, "Dummy Title", "Dummy Creator", "Dummy Publisher", 12345, "/home/dummy/dummy.epub"}
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
