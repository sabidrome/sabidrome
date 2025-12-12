package core

import (
	"fmt"
	"database/sql"

	"github.com/taylorskalyo/goreader/epub"
	_ "github.com/mattn/go-sqlite3"
)


type Book struct {
	Name string
	Author string
}

type Database struct {
    Type string
    Path string
}

func GetBookMetadataFromPath(path string) ( bookObject Book ) {

    rc, err := epub.OpenReader(path)
    if err != nil {
        panic(err)
    }
    defer rc.Close()

    book := rc.Rootfiles[0]

    bookObject = Book {
        Name: book.Title,
        Author: book.Creator,
    }

    fmt.Println(bookObject.Name)
    fmt.Println(bookObject.Author)

    return bookObject

}

// func CreateDatabase() (db *sql.DB) {

func CreateDatabase(d *Database) {

    // If db doesn't exist, it will create it
    db, err := sql.Open(d.Type, d.Path)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Connected to the SQLite database succesfully.")

    var sqliteVersion string
    err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer db.Close()

    fmt.Println(sqliteVersion)

}

func CreateDatabaseBookshelfTable(d *Database) {

    query := `CREATE TABLE IF NOT EXISTS bookshelf (
            bookid INTEGER PRIMARY KEY AUTOINCREMENT,
            name   TEXT NOT NULL,
            author TEXT NOT NULL

            );`

    db, err := sql.Open(d.Type, d.Path)
    if err != nil {
        fmt.Println(err)
        return
    }

    result, err := db.Exec(query)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(result)
    fmt.Println("Table 'bookshelf' was create succesfully.")

}

func AddBookToDatabase(d *Database, b *Book) {

    query := `INSERT INTO bookshelf (name, author) VALUES (?, ?);`

    db, err := sql.Open(d.Type, d.Path)
    if err != nil {
        fmt.Println(err)
        return
    }

    result, err := db.Exec(query, b.Name, b.Author)
    if err != nil {
        return
    }

    fmt.Printf("Book %s was create succesfully.\n", b.Name)

    fmt.Println(result.LastInsertId())

}
