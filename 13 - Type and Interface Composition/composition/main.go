package main

import (
    "fmt"
    "composition/store"
)

func main() {

    products := map[string]store.ItemForSale {
        "Kayak": store.NewBoat("Kayak", 279, 1, false),
        "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
    }

    for key, p := range products {

        switch item := p.(type) {
            case store.Describable:
                fmt.Println("Name:", item.GetName(), "Category:", item.GetCategory(), 
                    "Price:", item.Price(0.2))                                                  
            default:
                fmt.Println("Key:", key, "Price:", p.Price(0.2))
        }
    }
}
