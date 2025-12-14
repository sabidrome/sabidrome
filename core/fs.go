package core

import (
	"fmt"
	"os"

	"github.com/h2non/filetype"
	"github.com/taylorskalyo/goreader/epub"
)

type Book struct {
    Id int
    Name string
    Author string
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
        Path: path,
    }

    fmt.Println("[Debug] File information extracted.")

    fmt.Printf("[Debug] Name   -> %s \n", bookObject.Name)
    fmt.Printf("[Debug] Author -> %s \n", bookObject.Author)
    fmt.Printf("[Debug] Path   -> %s \n", bookObject.Path)

    return bookObject

}

func CheckValidFileType(path string) bool {

    buf, _ := os.ReadFile(path)

    kind, _ := filetype.Match(buf)
    if kind == filetype.Unknown {
        fmt.Println("[Debug] Unkown filetype")
        return false
    }

    if kind.Extension != "zip" {
        fmt.Printf("[Debug] Filetype %s is not supported.\n", kind.Extension)
        return false
    }

    fmt.Println("[Debug] File is a valid epub container")

    return true

}
