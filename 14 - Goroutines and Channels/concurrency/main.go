package main

import (
    "fmt"
    "time"
)

func enumerateProducts(channel1, channel2 chan<- *Product) {
    for _, p := range ProductList {
        select {
            case channel1 <- p:
                fmt.Println("Send via channel 1")
            case channel2 <- p:
                fmt.Println("Send via channel 2")
        }
    }
    close(channel1)
    close(channel2)
}

func main() {
    c1 := make(chan *Product, 2)
    c2 := make(chan *Product, 2)
    go enumerateProducts(c1, c2)
    
    time.Sleep(time.Second)

    for p := range c1 {
        fmt.Println("Channel 1 received product:", p.Name)
    }
    for p := range c2 {
        fmt.Println("Channel 2 received product:", p.Name)
    }    
}
