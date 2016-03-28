package main

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "strings"
    "path/filepath"
    "github.com/koyachi/go-term-ansicolor/ansicolor"
    "github.com/fsnotify/fsnotify"
)

func watcher(path, command string, args[]string) {
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
                if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write {
                    runCommand(command, args)
                }
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

func runCommand(command string, args[]string) {
    cmd := exec.Command(command, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    fmt.Println(ansicolor.Cyan("Xavier 0.0.1 is watching your files."))

    args := os.Args[1:]

    if len(args) == 0 {
        log.Fatal("Not enough arguments")
    }

    currentPath, _ := filepath.Abs(".")
    watcher(currentPath, args[0], args[1:])
}
