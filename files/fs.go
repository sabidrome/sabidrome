package files

import (
    "archive/zip"
    "io"
    "io/fs"
    "log/slog"
    "os"
    "path/filepath"
    "strings"
    "errors"
)

func ListDir(dir string) {

    aux_func := func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            slog.Error("Failed to access path", "path", path, "error", err)
            return err
        }
        slog.Debug("Visited file or dir", "path", path)
        return nil

    }

    err := filepath.Walk(dir, aux_func)
    if err != nil {
        slog.Error("Error walking the path", "path", dir)
    }

}

func ListBasePath(dir string) {

    isDirectory := func(path string) (bool, error) {
        fileInfo, err := os.Stat(path)
        if err != nil {
            return false, err
        }

        return fileInfo.IsDir(), err
    }

    aux_func := func(path string, info fs.FileInfo, err1 error) error {
        // Handle potential error
        if err1 != nil {
            slog.Error("Failed to access a path", "path", path, "error", err1)
            return err1
        }

        is_directory, _ := isDirectory(path)
        if !is_directory {
            // slog.Debug("Found element", "dir", filepath.Dir(path), "file", filepath.Base(path))
            slog.Debug("Found element", "file", filepath.Base(path))
        }
        return nil
    }

    err := filepath.Walk(dir, aux_func)
    if err != nil {
        slog.Error("Error walking the path", "path", dir)
    }

}

func ListEpubFileContent(path string) {

    zipListing, err := zip.OpenReader(path)
    if err != nil {
        slog.Error("Could not open zip file", "path", path)
    }
    defer zipListing.Close()
    for _, file := range zipListing.File {
        slog.Debug("File inside zip", "file", file.Name)
    }

}

func EpubContainerContent(zipPath string) (string, error) {

    // Epub spec guarantees this file
    containerPath := "META-INF/container.xml"

    zipFile, err := zip.OpenReader(zipPath)
    if err != nil {
        slog.Error("Failed to open archive", "path", zipPath)
        os.Exit(1)
    }
    defer zipFile.Close()

    var b []byte

    for _, file := range zipFile.File {
        if strings.EqualFold(file.Name, containerPath) {
            v, err := file.Open()
            if err != nil {
                slog.Error("Failed to open archived file", "name", file.Name)
                os.Exit(1)
            }
            defer v.Close()

            b, err = io.ReadAll(v)
            if err != nil {
                slog.Error("Failed to read content of archived file", "name", file.Name)
                os.Exit(1)
            }
        }
    }
    if b != nil {
        return string(b), nil
    }
    return "", errors.New("File not found")
}
