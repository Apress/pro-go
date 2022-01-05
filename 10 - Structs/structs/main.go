package main

import "fmt"

type Product struct {
    name, category string
    price float64
    *Supplier
}

type Supplier struct {
    name, city string
}

func main() {

    var prod Product = Product{ Supplier: &Supplier{}}
    var prodPtr *Product

    fmt.Println("Value:", prod.name, prod.category, prod.price, prod.Supplier.name)
    fmt.Println("Pointer:", prodPtr)
}
