package core

import (
	"fmt"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
    Type string
    Path string
}

func InitDatabase(d *Database) {

    db, err := sql.Open(d.Type, d.Path)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("[Debug] Succesfully connected to database %s\n", d.Path)

    query := `CREATE TABLE IF NOT EXISTS bookshelf (
        bookid INTEGER PRIMARY KEY AUTOINCREMENT,
        name   TEXT NOT NULL,
        author TEXT NOT NULL,
        path   TEXT NOT NULL

        );`

    _, err = db.Exec(query)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("[Debug] Succesfully created table 'bookshelf'")

    defer db.Close()

}

func AddBookToDatabase(d *Database, b *Book) {

    fmt.Printf("[Debug] '%s' requested for addition to database.\n", b.Name)

    query := `INSERT INTO bookshelf (name, author, path) VALUES (?, ?, ?);`

    db, err := sql.Open(d.Type, d.Path)
    if err != nil {
        fmt.Println(err)
        return
    }

    _, err = db.Exec(query, b.Name, b.Author, b.Path)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("[Debug] '%s' added to database succesfully.\n", b.Name)

}

func AddBook(db *Database, path string) {

	book_struct := GetBookMetadataFromPath(path)
	AddBookToDatabase(db, &book_struct)

}

func FindBookByName(d *Database, name string) (*Book, error) {

    fmt.Printf("[Debug] Looking '%s' as book name in database.\n", name)

    query := `SELECT * FROM bookshelf WHERE name LIKE '%' || ? || '%'`

    db, err := sql.Open(d.Type, d.Path)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    row   := db.QueryRow(query, name)
    b     := &Book{}

    var id int
    err = row.Scan(&id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return nil, err
    }

    fmt.Printf("[Debug] Found match for '%s' in %s\n", name, b.Path)

    return b, nil

}

func FindBookByAuthor(d *Database, author string) (*Book, error) {

    fmt.Printf("[Debug] Looking '%s' as book author in database.\n", author)

    query := `SELECT * FROM bookshelf WHERE author LIKE '%' || ? || '%'`

    // fmt.Println(query)

    db, err := sql.Open(d.Type, d.Path)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    row   := db.QueryRow(query, author)
    // fmt.Println(row)
    b     := &Book{}

    var id int
    err = row.Scan(&id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        // fmt.Println(err)
        return nil, err
    }

    fmt.Printf("[Debug] Found match for '%s' in %s\n", author, b.Path)

    return b, nil

}

func FindBook(db *Database, query string) {

    var book *Book
    var err error

    book, err = FindBookByName(db, query)
    if err != nil {
        book, err = FindBookByAuthor(db, query)
    }

    if book != nil {
        fmt.Printf("[Debug] Name   -> %s\n", book.Name)
        fmt.Printf("[Debug] Author -> %s\n", book.Author)
        fmt.Printf("[Debug] Path   -> %s\n", book.Path)
    } else {
        fmt.Printf("[Debug] There are no matches for '%s'\n", query)
    }
}
