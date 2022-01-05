package main

import (
    "net/http"
    "encoding/json"
)

func HandleJsonRequest(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(Products)
}

func init() {
    http.HandleFunc("/json", HandleJsonRequest)
}
