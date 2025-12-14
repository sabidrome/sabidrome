package core

import (
    "fmt"
    "os"
    "log/slog"

    "github.com/h2non/filetype"
    "github.com/taylorskalyo/goreader/epub"
)

type Book struct {
    Id int
    Name string
    Author string
    Path string
    Identifier string
}

func GetBookMetadataFromPath(path string) ( bookObject Book ) {

    rc, err := epub.OpenReader(path)
    if err != nil {
        panic(err)
    }
    defer rc.Close()

    book := rc.Rootfiles[0]

    fmt.Println(book.Identifier)

    bookObject = Book {
        Name: book.Title,
        Author: book.Creator,
        Path: path,
        Identifier: book.Identifier,
    }

    slog.Debug("File information extracted", "name", bookObject.Name, "author", bookObject.Author, "path", bookObject.Path, "identifier", bookObject.Identifier)

    return bookObject

}

func CheckValidFileType(path string) bool {

    buf, _ := os.ReadFile(path)

    kind, _ := filetype.Match(buf)
    if kind == filetype.Unknown {
    slog.Debug("Unkown filetype")
        return false
    }

    if kind.Extension != "zip" {
    slog.Debug("Filetype not supported", "kind", kind.Extension)
        return false
    }

    slog.Debug("File is a valid epub container", "kind", kind.Extension)

    return true

}
