package main

import (
    "fmt"
    "regexp"
)

func main() {
    
    pattern := regexp.MustCompile(
        "A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")


    description := "Kayak. A boat for one person."

    replaced := pattern.ReplaceAllStringFunc(description, func(s string) string {
        return "This is the replacement content"
    })
    fmt.Println(replaced)
}
