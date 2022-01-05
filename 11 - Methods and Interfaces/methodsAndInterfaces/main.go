package main

import "fmt"

type Expense interface {
    getName() string
    getCost(annual bool) float64
}

type Person struct {
    name, city string
}

func processItems(items ...interface{}) {
    for _, item := range items {
        switch value := item.(type) {
            case Product:
                fmt.Println("Product:", value.name, "Price:", value.price)           
            case *Product:
                fmt.Println("Product Pointer:", value.name, "Price:", value.price)                           
            case Service:
                fmt.Println("Service:", value.description, "Price:", 
                    value.monthlyFee * float64(value.durationMonths))
            case Person:                
                fmt.Println("Person:", value.name, "City:", value.city)
            case *Person:
                fmt.Println("Person Pointer:", value.name, "City:", value.city)                
            case string, bool, int:
                fmt.Println("Built-in type:", value)
            default:
                fmt.Println("Default:", value)       
        }
    }
}

func main() {

    var expense Expense = &Product { "Kayak", "Watersports", 275 }

    data := []interface{} {
        expense,
        Product { "Lifejacket", "Watersports", 48.95 },
        Service {"Boat Cover", 12, 89.50, []string{} },
        Person { "Alice", "London"},
        &Person { "Bob", "New York"},
        "This is a string",
        100,
        true,
    }

    processItems(data...)
}
