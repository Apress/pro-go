package main

import (
    //"fmt"
    "time"
)

func writeToChannel(nameChannel chan <- string) {

    names := []string { "Alice", "Bob", "Charlie", "Dora" }

    ticker := time.NewTicker(time.Second / 10)
    index := 0
    for {
        <- ticker.C
        nameChannel <- names[index]
        index++
        if (index == len(names)) {
            ticker.Stop()
            close(nameChannel)
            break
        }
    }
}

func main() {

    nameChannel := make (chan string)

    go writeToChannel(nameChannel)

    for name := range nameChannel {
        Printfln("Read name: %v", name)
    }
}
