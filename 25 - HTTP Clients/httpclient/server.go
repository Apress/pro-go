package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {

    http.HandleFunc("/html", 
        func (writer http.ResponseWriter, request *http.Request) {
            http.ServeFile(writer, request, "./index.html")
        })
    http.HandleFunc("/json", 
        func (writer http.ResponseWriter, request *http.Request) {
            writer.Header().Set("Content-Type", "application/json")
            json.NewEncoder(writer).Encode(Products)
        })
    http.HandleFunc("/echo", 
        func (writer http.ResponseWriter, request *http.Request) {
            writer.Header().Set("Content-Type", "text/plain")
            fmt.Fprintf(writer, "Method: %v\n", request.Method)
            for header, vals := range request.Header {
                fmt.Fprintf(writer, "Header: %v: %v\n", header, vals)
            }
            fmt.Fprintln(writer, "----")
            data, err := io.ReadAll(request.Body)
            if (err == nil) {
                if len(data) == 0 {
                    fmt.Fprintln(writer, "No body")                
                } else {
                    writer.Write(data)
                }
            } else {
                fmt.Fprintf(os.Stdout,"Error reading body: %v\n", err.Error())
            }
        })        
}
