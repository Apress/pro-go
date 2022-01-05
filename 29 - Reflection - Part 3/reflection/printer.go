package main

import "fmt"

func Printfln(template string, values ...interface{}) {
    fmt.Printf(template + "\n", values...)
}
