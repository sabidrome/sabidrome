package core

import (
    "os"
    "strconv"
    "log/slog"
    "database/sql"

    _ "github.com/mattn/go-sqlite3"
)

func InitDatabase(dialect string, path string) (db *sql.DB) {

    db, err1 := sql.Open(dialect, path)
    if err1 != nil {
        slog.Error("Database dialect or path is invalid")
        os.Exit(1)
    }

    query := `CREATE TABLE IF NOT EXISTS bookshelf (
        bookid INTEGER PRIMARY KEY AUTOINCREMENT,
        name   TEXT NOT NULL,
        author TEXT NOT NULL,
        path   TEXT NOT NULL

        );`

    result, err2 := db.Exec(query)
    if err2 != nil {
        slog.Debug("Something went wrong", "result", result, "err", err2)
        slog.Error("Failed to connect to database")
        os.Exit(1)
    }

    slog.Info("Database initialized at", "path", path)

    return db

}

func AddBookToDatabase(db *sql.DB, b *Book) {

    query := `INSERT INTO bookshelf (name, author, path) VALUES (?, ?, ?);`

    _, err := db.Exec(query, b.Name, b.Author, b.Path)
    if err != nil {
        return
    }

    slog.Debug("New book in database", "path", b.Path)

}

func AddBook(db *sql.DB, path string) {

    exists := FindPathInDatabase(db, path)
    if exists {
        slog.Debug("Book exists in database", "path", path)
        return
    }

    book_struct := GetBookMetadataFromPath(path)
    AddBookToDatabase(db, &book_struct)

}

func RemoveBookFromDatabase(db *sql.DB, path string) {

    query := `DELETE FROM bookshelf WHERE path = ?`

    _, err := db.Exec(query, path)
    if err != nil {
        slog.Error("Could not add book to database", "path", path)
        os.Exit(1)
    }

    slog.Debug("Book removed from database", "path", path)

}

func RemoveBook(db *sql.DB, path string) {

    exists_in_db := FindPathInDatabase(db, path)
    if !exists_in_db {
        slog.Debug("File does not exist in database", "path", path)
        slog.Debug("Ignoring request")
        return
    }

    _, err := os.Stat(path)
    if err == nil {
        slog.Debug("File exists in filesystem", "path", path)
        slog.Debug("Ignoring request")
        return
    }

    RemoveBookFromDatabase(db, path)

}

func ListBookshelf(db *sql.DB) {

    query := `SELECT * FROM bookshelf`

    rows, err := db.Query(query)
    if err != nil {
        slog.Error("Could not query bookshelf table")
        os.Exit(1)
    }

    for rows.Next() {
        var id int
        var name string
        var author string
        var path string
        err := rows.Scan(&id, &name, &author, &path)
        if err != nil {
            slog.Error("Could not scan bookshelft row")
            os.Exit(1)
        }
        str_id := strconv.Itoa(id)
        slog.Info("Current library", "id", str_id, "name", name, "author", author, "path", path)
    }

}

func FindBookByName(db *sql.DB, name string) (*Book, error) {

    slog.Debug("Looking book by name in database", "name", name)

    query := `SELECT * FROM bookshelf WHERE name LIKE '%' || ? || '%'`

    row   := db.QueryRow(query, name)
    b     := &Book{}

    err := row.Scan(&b.Id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return nil, err
    }

    return b, nil

}

func FindBookByAuthor(db *sql.DB, author string) (*Book, error) {

    slog.Debug("Looking book by author in database", "author", author)

    query := `SELECT * FROM bookshelf WHERE author LIKE '%' || ? || '%'`

    row   := db.QueryRow(query, author)
    b     := &Book{}

    err := row.Scan(&b.Id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return nil, err
    }

    return b, nil

}

func FindBook(db *sql.DB, query string) (int) {

    book, err := FindBookByName(db, query)
    if err != nil {
        book, err = FindBookByAuthor(db, query)
    }

    if book != nil {

        slog.Debug("Found a match", "query", query)
        slog.Debug("Metadata", "id", book.Id)
        slog.Debug("Metadata", "name", book.Name)
        slog.Debug("Metadata", "author", book.Author)
        slog.Debug("Metadata", "path", book.Path)

	return book.Id
    } else {
        slog.Debug("There are no matches", "query", query)
	return -1
    }

}

func FindPathInDatabase(db *sql.DB, path string) bool {

    slog.Debug("Looking exact match for", "path", path)

    query := `SELECT * FROM bookshelf WHERE path = ?`

    row   := db.QueryRow(query, path)
    b     := &Book{}

    err := row.Scan(&b.Id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return false
    }

    return true

}
