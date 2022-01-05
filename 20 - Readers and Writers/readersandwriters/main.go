package main

import (
    "io"
    "strings"
    //"bufio"
    "fmt"
)

func writeReplaced(writer io.Writer, str string, subs ...string) {
    replacer := strings.NewReplacer(subs...)
    replacer.WriteString(writer, str)
}

func main() {

    text := "It was a boat. A small boat."
    subs := []string { "boat", "kayak", "small", "huge" }

    var writer strings.Builder
    writeReplaced(&writer, text, subs...)
    fmt.Println(writer.String())
}
