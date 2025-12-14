package core

import (
    "fmt"
    "os"
    "database/sql"

    _ "github.com/mattn/go-sqlite3"
)

func InitDatabase(dialect string, path string) (db *sql.DB) {

    db, err1 := sql.Open(dialect, path)
    if err1 != nil {
        fmt.Println(err1)
        return nil
    }

    query := `CREATE TABLE IF NOT EXISTS bookshelf (
        bookid INTEGER PRIMARY KEY AUTOINCREMENT,
        name   TEXT NOT NULL,
        author TEXT NOT NULL,
        path   TEXT NOT NULL

        );`

    _, err2 := db.Exec(query)
    if err2 != nil {
        fmt.Println(err2)
        return nil
    }

    fmt.Println("[Debug] Succesfully initialized database.")

    return db

}

func AddBookToDatabase(db *sql.DB, b *Book) {

    query := `INSERT INTO bookshelf (name, author, path) VALUES (?, ?, ?);`

    _, err := db.Exec(query, b.Name, b.Author, b.Path)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("[Debug] '%s' added to database succesfully.\n", b.Name)

}

func AddBook(db *sql.DB, path string) {

    exists := FindPathInDatabase(db, path)
    if exists {
        fmt.Println("[Debug] File exists on database, skipping.")
        return
    }

    book_struct := GetBookMetadataFromPath(path)
    AddBookToDatabase(db, &book_struct)

}

func RemoveBookFromDatabase(db *sql.DB, path string) {

    query := `DELETE FROM bookshelf WHERE path = ?`

    _, err := db.Exec(query, path)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("[Debug] '%s' removed from database succesfully\n", path)

}

func RemoveBook(db *sql.DB, path string) {

    exists_in_db := FindPathInDatabase(db, path)
    if !exists_in_db {
        fmt.Println("[Debug] File does not exist on database, ignoring request.")
        return
    }

    _, err := os.Stat(path)
    if err == nil {
        fmt.Println("[Debug] File exists on filesystem, ignoring request.")
        return
    }

    RemoveBookFromDatabase(db, path)

}

func ListBookshelf(db *sql.DB) {

    query := `SELECT * FROM bookshelf`

    rows, err := db.Query(query)
    if err != nil {
        fmt.Println(err)
        return
    }

    for rows.Next() {
        var id int
        var name string
        var author string
        var path string
        err := rows.Scan(&id, &name, &author, &path)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("%d || %s || %s || %s\n", id, name, author, path)
    }

}

func FindBookByName(db *sql.DB, name string) (*Book, error) {

    fmt.Printf("[Debug] Looking '%s' as book name in database.\n", name)

    query := `SELECT * FROM bookshelf WHERE name LIKE '%' || ? || '%'`

    row   := db.QueryRow(query, name)
    b     := &Book{}

    err := row.Scan(&b.Id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return nil, err
    }

    fmt.Printf("[Debug] Found match for '%s' in %s\n", name, b.Path)

    return b, nil

}

func FindBookByAuthor(db *sql.DB, author string) (*Book, error) {

    fmt.Printf("[Debug] Looking '%s' as book author in database.\n", author)

    query := `SELECT * FROM bookshelf WHERE author LIKE '%' || ? || '%'`

    row   := db.QueryRow(query, author)
    b     := &Book{}

    err := row.Scan(&b.Id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return nil, err
    }

    fmt.Printf("[Debug] Found match for '%s' in %s\n", author, b.Path)

    return b, nil

}

func FindBook(db *sql.DB, query string) (int) {

    book, err := FindBookByName(db, query)
    if err != nil {
        book, err = FindBookByAuthor(db, query)
    }

    if book != nil {
        fmt.Printf("[Debug] Name   -> %s\n", book.Name)
        fmt.Printf("[Debug] Author -> %s\n", book.Author)
        fmt.Printf("[Debug] Path   -> %s\n", book.Path)
	return book.Id
    } else {
        fmt.Printf("[Debug] There are no matches for '%s'\n", query)
	return -1
    }

}

func FindPathInDatabase(db *sql.DB, path string) bool {

    fmt.Printf("[Debug] Looking exact match for %s\n", path)

    query := `SELECT * FROM bookshelf WHERE path = ?`

    row   := db.QueryRow(query, path)
    b     := &Book{}

    err := row.Scan(&b.Id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return false
    }

    return true

}
