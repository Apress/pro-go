package main

import (
    "net/http"
    "fmt"
    "io"
)

func init() {

    http.HandleFunc("/form", 
        func (writer http.ResponseWriter, request *http.Request) {
            err := request.ParseMultipartForm(10000000)
            if (err == nil) {
                for name, vals := range request.MultipartForm.Value {
                    fmt.Fprintf(writer, "Field %v: %v\n", name, vals)    
                }
                for name, files := range request.MultipartForm.File {
                    for _, file := range files {
                        fmt.Fprintf(writer, "File %v: %v\n", name, file.Filename)    
                        if f, err := file.Open(); err == nil {
                            defer f.Close()
                            io.Copy(writer, f)
                        }
                    }
                }                
            } else {
                fmt.Fprintf(writer, "Cannot parse form %v", err.Error())           
            }
        })
}
