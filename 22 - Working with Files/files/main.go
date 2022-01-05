package main

import (
    // "fmt"
    // "time"
    "os"
    //"encoding/json"
    "path/filepath"
)

func callback(path string, dir os.DirEntry, dirErr error) (err error) {
    info, _ := dir.Info()
    Printfln("Path %v, Size: %v", path, info.Size())
    return
}

func main() {

    path, err := os.Getwd()
    if (err == nil) {
        err = filepath.WalkDir(path, callback)
    } else {
        Printfln("Error %v", err.Error())
    }
}
