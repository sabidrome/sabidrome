package core

import (
    "log/slog"
    "strings"
    "path/filepath"
    "database/sql"

    "github.com/andreaskoch/go-fswatch"
)

func FilesystemWatcher(db *sql.DB) {

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
                // fmt.Println(" -> New or modified items detected")

            case <-folderWatcher.Moved():
                // fmt.Println(" -> Items have been moved")

            case changes := <-folderWatcher.ChangeDetails():
                // fmt.Printf("    -> '%s'\n", changes.String())
                // fmt.Printf("    -> Modified: '%#v'\n", changes.Modified())

                // Remove files are catched here
                for i:=0; i<len(changes.Moved()); i++ {
                    book_path := changes.Moved()[i]
                    slog.Info("Missing file", "path", book_path)
                    RemoveBook(db, book_path)
                    ListBookshelf(db)
                }

                for i:=0; i<len(changes.New()); i++ {
                    book_path := changes.New()[i]
                    slog.Info("New file", "path", book_path)
                    AddBook(db, book_path)
                    ListBookshelf(db)
                }
        }
    }
}
