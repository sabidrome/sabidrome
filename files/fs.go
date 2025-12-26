package files

import (
    "io/fs"
    "log/slog"
    "os"
    "path/filepath"
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

