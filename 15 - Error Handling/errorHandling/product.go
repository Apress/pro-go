package main

import "strconv"

type Product struct {
    Name, Category string
    Price float64
}

type ProductSlice []*Product

var Products = ProductSlice {
    { "Kayak", "Watersports", 279 },
    { "Lifejacket", "Watersports", 49.95 },
    { "Soccer Ball", "Soccer", 19.50 },
    { "Corner Flags", "Soccer", 34.95 },
    { "Stadium", "Soccer", 79500 },        
    { "Thinking Cap", "Chess", 16 },
    { "Unsteady Chair", "Chess", 75 },
    { "Bling-Bling King", "Chess", 1200 },    
}

func ToCurrency(val float64) string {
    return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}
