package main

import (
    "fmt"
    "sort"
)

func main() {

    names := [3]string {"Alice", "Charlie", "Bob"}

    secondPosition := &names[1]

    fmt.Println(*secondPosition)

    sort.Strings(names[:])

    fmt.Println(*secondPosition)
}
