package main

import (
    "net/http"
    "io"
    "strings"
)

type StringHandler struct {
    message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, 
        request *http.Request) {
    Printfln("Request for %v", request.URL.Path)
    io.WriteString(writer, sh.message)    
}

func HTTPSRedirect(writer http.ResponseWriter, 
        request *http.Request) {
    host := strings.Split(request.Host, ":")[0] 
    target := "https://" + host + ":5500" + request.URL.Path 
    if len(request.URL.RawQuery) > 0 {
        target += "?" + request.URL.RawQuery
    }
    http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}

func main() {
    http.Handle("/message", StringHandler{ "Hello, World"})
    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

    fsHandler := http.FileServer(http.Dir("./static"))
    http.Handle("/files/", http.StripPrefix("/files", fsHandler))

    go func () {
        err := http.ListenAndServeTLS(":5500", "certificate.cer", 
            "certificate.key", nil)
        if (err != nil) {
            Printfln("HTTPS Error: %v", err.Error())
        }
    }()

    err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
    if (err != nil) {
        Printfln("Error: %v", err.Error())
    }
}
