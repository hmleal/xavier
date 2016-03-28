package main

import (
    "fmt"
    "os"
    "strings"
    "path/filepath"
    "github.com/koyachi/go-term-ansicolor/ansicolor"
)

func header() {
    fmt.Println(ansicolor.Cyan("Xavier 0.0.1 is watching your files."))
}

func SubFolders(path string) (paths []string) {
    filepath.Walk(path, func (newPath string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() {
            name := info.Name()
            if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_") {
                return filepath.SkipDir
            }
            paths = append(paths, newPath)
        }
        return nil
    })
    return paths
}

func main() {
    header()
    paths := SubFolders("/Users/hmleal/work/gocode")
    for key, value := range paths {
        fmt.Println(key, value)
    }
}
