package main 

import (
    "fmt"
    "strconv"
)

func main() {

    val := 49.95

    Fstring := strconv.FormatFloat(val, 'f', 2, 64)
    Estring := strconv.FormatFloat(val, 'e', -1, 64)
    
    fmt.Println("Format F: " + Fstring)        
    fmt.Println("Format E: " + Estring)        
}
