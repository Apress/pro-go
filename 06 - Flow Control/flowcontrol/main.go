package main

import (
    "fmt"
    //"strconv"
)

func main() {
    
    counter := 0
    target: fmt.Println("Counter", counter)
    counter++;
    if (counter < 5) {
        goto target 
    }
}
