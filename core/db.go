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

    query := `INSERT INTO bookshelf (name, author, path) VALUES (?, ?, ?);`

    db, err := sql.Open(d.Type, d.Path)
    if err != nil {
        fmt.Println(err)
        return
    }

    _, err = db.Exec(query, b.Name, b.Author, b.Path)
    if err != nil {
        return
    }

    fmt.Printf(" -> Book '%s' added to database.\n", b.Name)

}

func AddBook(db *Database, path string) {

	fmt.Println("[Debug] Begin add a new book")
	book_struct := GetBookMetadataFromPath(path)
	AddBookToDatabase(db, &book_struct)
	fmt.Println("[Debug] End add a new book")

}

func FindBookByName(d *Database, name string) (*Book, error) {

    query := `SELECT * FROM bookshelf WHERE name LIKE '%' || ? || '%'`

    // fmt.Println(query)

    db, err := sql.Open(d.Type, d.Path)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    row   := db.QueryRow(query, name)
    // fmt.Println(row)
    b     := &Book{}

    var id int
    err = row.Scan(&id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        // fmt.Println(err)
        return nil, err
    }

    return b, nil

}

func FindBookByAuthor(d *Database, author string) (*Book, error) {

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

    return b, nil

}

func FindBook(db *Database, query string) {

    fmt.Println("[Debug] Being search a book")
    var book *Book
    var err error

    book, err = FindBookByName(db, query)
    if err != nil {
        fmt.Printf(" -> '%s' is not a book name in the db\n", query)
    }

    book, err = FindBookByAuthor(db, query)
    if err != nil {
        fmt.Printf(" -> '%s' is not a book author in the db\n", query)
    }

    if book != nil {
        fmt.Println(" -> Found a match in the db")
        fmt.Printf("    -> Name: %s\n", book.Name)
        fmt.Printf("    -> Author: %s\n", book.Author)
        fmt.Printf("    -> Path: %s\n", book.Path)
    } else {
        fmt.Println(" -> Did not find a match in the db")
    }

    fmt.Println("[Debug] End search a book")
}
