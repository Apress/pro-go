package main 

import (
    "fmt"
    "packages/store"
    . "packages/fmt"
    "packages/store/cart"
    _ "packages/data"
    //"github.com/fatih/color"
)

func main() {

    product := store.NewProduct("Kayak", "Watersports", 279) 

    cart := cart.Cart {
        CustomerName: "Alice",
        Products: []store.Product{ *product },
    }

    // color.Green("Name: " + cart.CustomerName)
    // color.Cyan("Total: " + ToCurrency(cart.GetTotal()))
    fmt.Println("Name:", cart.CustomerName)
    fmt.Println("Total:",  ToCurrency(cart.GetTotal()))
}
