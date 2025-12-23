package db

import (
    "log/slog"
    "database/sql"

    _ "github.com/mattn/go-sqlite3"
    "github.com/sabidrome/sabidrome/core"
)


func _debugPrintSearchResult(query string, book *core.Book) {

    if book == nil {
        slog.Debug("No matches found", "query", query)
    }

    slog.Debug("Found a match", "query", query)
    slog.Debug("Metadata", "id", book.Id)
    slog.Debug("Metadata", "name", book.Name)
    slog.Debug("Metadata", "author", book.Author)
    slog.Debug("Metadata", "path", book.Path)

}

func SearchBookByName(db *sql.DB, name string, book *core.Book) bool {

    slog.Debug("Searching book by name in db", "name", name)
    query := `SELECT * FROM bookshelf WHERE name LIKE '%' || ? || '%'`

    row   := db.QueryRow(query, name)
    b     := &core.Book{}

    err := row.Scan(&b.Id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return false
    }

    return true

}

func SearchBookByAuthor(db *sql.DB, author string, book *core.Book) bool {

    slog.Debug("Searching book by author in db", "author", author)
    query := `SELECT * FROM bookshelf WHERE author LIKE '%' || ? || '%'`

    row   := db.QueryRow(query, author)
    b     := &core.Book{}

    err := row.Scan(&b.Id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return false
    }

    return true

}

func SearchBookByPatternInBasePath(db *sql.DB, pattern string, book *core.Book) bool {

    slog.Debug("Searching book by pattern in path in db", "pattern", pattern)
    query := `SELECT * FROM bookshelf WHERE path LIKE '%' || ? || '%'`

    row   := db.QueryRow(query, pattern)
    b     := &core.Book{}

    err   := row.Scan(&b.Id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return false
    }

    return true

}

func SearchBookByExactFullPath(db *sql.DB, path string, book *core.Book) bool {

    slog.Debug("Searching book by exact path in db", "path", path)
    query := `SELECT * FROM bookshelf WHERE path = ?`

    row   := db.QueryRow(query, path)
    b     := &core.Book{}

    err := row.Scan(&b.Id, &b.Name, &b.Author, &b.Path)
    if err != nil {
        return false
    }

    return true

}

func SearchBook(db *sql.DB, query string) (*core.Book, bool) {

    book := &core.Book{}

    switch {
        case SearchBookByName(db, query, book):
            _debugPrintSearchResult(query, book)
            return book, true

        case SearchBookByAuthor(db, query, book):
            _debugPrintSearchResult(query, book)
            return book, true

        case SearchBookByPatternInBasePath(db, query, book):
            _debugPrintSearchResult(query, book)
            return book, true

        default:
            return nil, false
    }

}
