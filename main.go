package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "path/filepath"
    "github.com/koyachi/go-term-ansicolor/ansicolor"
    "github.com/fsnotify/fsnotify"
)

func header() {
    fmt.Println(ansicolor.Cyan("Xavier 0.0.1 is watching your files."))
}

func watcher(path string) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    done := make(chan bool)
    go func() {
        for {
            select {
            case event := <-watcher.Events:
                log.Println("event:", event)
            case err := <-watcher.Errors:
                log.Println("error", err)
            }
        }
    }()

    folders := SubFolders(path)

    for _, folder := range folders {
        err := watcher.Add(folder)
        if err != nil {
            log.Println("Error watching: ", folder, err)
        }
    }
    <-done
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

    currentPath, _ := filepath.Abs(".")
    watcher(currentPath)
}
