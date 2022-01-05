package main

import "fmt"

func calcTotalPrice(products map[string]float64) (count int, total float64)  {
    fmt.Println("Function started")
    defer fmt.Println("First defer call")
    count = len(products)
    for _, price := range products {
        total += price
    }
    defer fmt.Println("Second defer call")
    fmt.Println("Function about to return")
    return
}

func main() {

    products := map[string]float64 {
        "Kayak" : 275,
        "Lifejacket": 48.95,
    }

    _, total  := calcTotalPrice(products)
    fmt.Println("Total:", total)
}
