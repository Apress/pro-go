package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ { // loop with a counter
		PrintHello()   // print out a message
		PrintNumber(i) // print out the counter
	}
}

func PrintHello() {
	fmt.Println("Hello, Go")
}

func PrintNumber(number int) {
	fmt.Println(number)
}
