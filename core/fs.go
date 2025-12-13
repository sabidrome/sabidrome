package core

import (
    "fmt"
    "path/filepath"
    "strings"

    "github.com/andreaskoch/go-fswatch"
    "github.com/taylorskalyo/goreader/epub"
)

type Book struct {
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

    fmt.Println(" -> File information extracted.")

    fmt.Printf("    -> %s \n", bookObject.Name)
    fmt.Printf("    -> %s \n", bookObject.Author)
    fmt.Printf("    -> %s \n", bookObject.Path)

    return bookObject

}

func FilesystemWatcher() {

    recurse := true

    skipDotFilesAndFolders := func(path string) bool {
        return strings.HasPrefix(filepath.Base(path), ".")
    }

    checkIntervalInSeconds := 2

    folderWatcher := fswatch.NewFolderWatcher(
        "/tmq/test",
        recurse,
        skipDotFilesAndFolders,
        checkIntervalInSeconds,
    )

    folderWatcher.Start()

    for folderWatcher.IsRunning() {
        select {

            case <-folderWatcher.Modified():
                fmt.Println(" -> New or modified items detected")

            case <-folderWatcher.Moved():
                fmt.Println(" -> Items have been moved")

            case changes := <-folderWatcher.ChangeDetails():
                fmt.Printf("    -> '%s'\n", changes.String())
                fmt.Printf("    -> New: '%#v'\n", changes.New())
                fmt.Printf("    -> Modified: '%#v'\n", changes.Modified())
                fmt.Printf("    -> Moved: '%v'\n", changes.Moved())
        }
    }

}
