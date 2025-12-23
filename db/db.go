package db

import (
    "os"
    "log/slog"
    "database/sql"

    "github.com/sabidrome/sabidrome/core"
)

func ConnectOrCreateDatabase() *sql.DB {

    db, err := sql.Open("sqlite3", "./sabidrome.db?_pragma=foreign_keys(1)")
    if err != nil {
        slog.Error("Database connection failed", "error", err)
        os.Exit(1)
    }

    query := `CREATE TABLE IF NOT EXISTS bookshelf (
        id         INTEGER PRIMARY KEY AUTOINCREMENT,
        title      TEXT NOT NULL,
        creator    TEXT NOT NULL,
        publisher  TEXT NOT NULL,
        isbn       INTEGRER NOT NULL,
        path       TEXT NOT NULL
        );`

        db.Exec(query)

        slog.Debug("Connected to database")

        return db
}

func AddBook(db *sql.DB, b *core.Book) int64 {

    query := `INSERT INTO bookshelf (title, creator, publisher, isbn, path) VALUES (?, ?, ?, ?, ?);`

    result, err := db.Exec(query, &b.Title, &b.Creator, &b.Publisher, &b.ISBN, &b.Path)
    if err != nil {
        slog.Error("Could not add book to database", "title", b.Title, "path", b.Path, "result", result)
        os.Exit(1)
    }

    b.Id, _ = result.LastInsertId()
    slog.Debug("Book added to database", "id", b.Id, "title", b.Title, "path", b.Path)

    return b.Id
}

func RemoveBook(db *sql.DB, id int64) {

    query := `DELETE FROM bookshelf WHERE id = ?`

    result, err := db.Exec(query, id)
    if err != nil {
        slog.Error("Could not remove book from database", "id", id, "result", result)
    }

    slog.Debug("Book removed from database", "id", id)
}

func UpdateBookPath(db *sql.DB, id int64, new_path string) {

    query :=`UPDATE bookshelf SET path = ? WHERE id = ?;`

    result, err := db.Exec(query, new_path, id)
    if err != nil {
        slog.Error("Could not update book in database", "id", id, "new_path", new_path, "result", result)
        os.Exit(1)
    }

    slog.Debug("Book updated in database", "id", id, "new_path", new_path)
}

func BooksList(db *sql.DB) {

    query   := `SELECT * FROM bookshelf;`

    rows, err := db.Query(query)
    if err != nil {
        slog.Error("Could not list books from database")
        os.Exit(1)
    }

    // fmt.Printf("Title || Creator || Path\n")

    var books []core.Book
    for rows.Next() {
        b := &core.Book{}
        err := rows.Scan(&b.Id, &b.Title, &b.Creator, &b.Publisher, &b.ISBN, &b.Path)
        if err != nil {
            slog.Error("Error while fetching book list from database")
            os.Exit(1)
        }
        // fmt.Printf("%s || %s || %s\n", b.Title, b.Creator, b.Path)
        books = append(books, *b)
    }
}
