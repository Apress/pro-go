package main

import (
    "net/http"
    "fmt"
    "strconv"
)

func GetAndSetCookie(writer http.ResponseWriter, request *http.Request) {

    counterVal := 1
    counterCookie, err := request.Cookie("counter")
    if (err == nil) {
        counterVal, _ = strconv.Atoi(counterCookie.Value)
        counterVal++
    } 
    http.SetCookie(writer, &http.Cookie{
        Name: "counter", Value: strconv.Itoa(counterVal),
    })

    if (len(request.Cookies()) > 0) {
        for _, c := range request.Cookies() {
            fmt.Fprintf(writer, "Cookie Name: %v, Value: %v", c.Name, c.Value)
        }
    } else {
        fmt.Fprintln(writer, "Request contains no cookies")
    }
}

func init() {
    http.HandleFunc("/cookies", GetAndSetCookie)
}
