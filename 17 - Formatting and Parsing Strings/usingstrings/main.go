package main

import "fmt"

func Printfln(template string, values ...interface{}) {
    fmt.Printf(template + "\n", values...)
}

func main() {

    var name string
    var category string
    var price float64
    
    source := "Product Lifejacket Watersports 48.95"
    template := "Product %s %s %f"
    n, err := fmt.Sscanf(source, template, &name, &category, &price)

    if (err == nil) {
        Printfln("Scanned %v values", n)
        Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
    } else {
        Printfln("Error: %v", err.Error())
    }
}
