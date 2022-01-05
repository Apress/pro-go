package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
    fmt.Println("Product:", product, "Price:", calculator(price))
}

var prizeGiveaway = false

func priceCalcFactory(threshold, rate float64, zeroPrices *bool) calcFunc {
    return func(price float64) float64 {
        if (*zeroPrices) {
            return 0
        } else if (price > threshold) {
            return price + (price * rate)
        }
        return price
    }
}

func main() {

    watersportsProducts := map[string]float64 {
        "Kayak" : 275,
        "Lifejacket": 48.95,
    }

    soccerProducts := map[string] float64 {
        "Soccer Ball": 19.50,
        "Stadium": 79500,
    }

    prizeGiveaway = false
    waterCalc := priceCalcFactory(100, 0.2, &prizeGiveaway); 
    prizeGiveaway = true
    soccerCalc := priceCalcFactory(50, 0.1, &prizeGiveaway)

    for product, price := range watersportsProducts {
        printPrice(product, price, waterCalc)
    }    

    for product, price := range soccerProducts {
        printPrice(product, price, soccerCalc)
    }    
}
