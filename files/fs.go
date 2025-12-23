package files

import (
    "fmt"
    "path/filepath"
    "io/fs"
    "log/slog"
)

func ListDir(dir string) {

    err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
            if err != nil {
                slog.Error("Failed to access path", "path", path, "error", err)
                return err
            }
            slog.Debug("Visited file or dir", "path", path)
            return nil
    }
    if err != nil {
        slog.Error("Error walking the path", "path", dir)
    }

    )

}
