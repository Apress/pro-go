package main

import (
    "net/http"
    "os"
    "io"
    "time"
    //"encoding/json"
    //"strings"
    //"net/url"
    //"net/http/cookiejar"
    //"fmt"
    "mime/multipart"
    "bytes"
)

func main() {
    go http.ListenAndServe(":5000", nil)
    time.Sleep(time.Second)

    var buffer bytes.Buffer
    formWriter := multipart.NewWriter(&buffer)
    fieldWriter, err := formWriter.CreateFormField("name")
    if (err == nil) {
        io.WriteString(fieldWriter, "Alice")
    }
    fieldWriter, err = formWriter.CreateFormField("city")
    if (err == nil) {
        io.WriteString(fieldWriter, "New York")
    }
    fileWriter, err := formWriter.CreateFormFile("codeFile", "printer.go")
    if (err == nil) {
        fileData, err := os.ReadFile("./printer.go")
        if (err == nil) {
            fileWriter.Write(fileData)
        }
    }

    formWriter.Close()

    req, err := http.NewRequest(http.MethodPost, 
        "http://localhost:5000/form", &buffer)

    req.Header["Content-Type"] = []string{ formWriter.FormDataContentType()}

    if (err == nil) {
        var response *http.Response
        response, err = http.DefaultClient.Do(req)                
        if (err == nil) {        
            io.Copy(os.Stdout, response.Body)
        } else {
            Printfln("Request Error: %v", err.Error())            
        }
    } else {
        Printfln("Error: %v", err.Error())
    } 
}
