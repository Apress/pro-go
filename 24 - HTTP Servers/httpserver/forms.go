package main

import (
	"net/http"
	"strconv"
)

func ProcessFormData(writer http.ResponseWriter, request *http.Request) {
    if (request.Method == http.MethodPost) {
        index, _ := strconv.Atoi(request.PostFormValue("index"))                
        p := Product {}
        p.Name = request.PostFormValue("name")
        p.Category = request.PostFormValue("category")
        p.Price, _ = strconv.ParseFloat(request.PostFormValue("price"), 64)
        Products[index] = p
    } 
    http.Redirect(writer, request, "/templates", http.StatusTemporaryRedirect)
}

func init() {
    http.HandleFunc("/forms/edit", ProcessFormData)
}
