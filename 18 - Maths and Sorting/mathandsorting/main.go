package main

import (
    //"sort"
)

func main() {

    products := []Product {
        { "Kayak", 279} ,
        { "Lifejacket", 49.95 },
        { "Soccer Ball",  19.50 },
    }

    SortWith(products, func (p1, p2 Product) bool {
        return p1.Name < p2.Name
    })   

    for _, p := range products {
        Printfln("Name: %v, Price: %.2f",  p.Name, p.Price)
    }
}
