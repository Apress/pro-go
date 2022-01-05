package main

import (
    "reflect"
    //"strings"
    // "fmt"
)

func swap(first interface{}, second interface{}) {
    firstValue, secondValue := reflect.ValueOf(first), reflect.ValueOf(second)
    if firstValue.Type() == secondValue.Type() && 
            firstValue.Kind() == reflect.Ptr && 
            firstValue.Elem().CanSet() && secondValue.Elem().CanSet() {

        temp :=  reflect.New(firstValue.Elem().Type())
        temp.Elem().Set(firstValue.Elem())
        firstValue.Elem().Set(secondValue.Elem())
        secondValue.Elem().Set(temp.Elem())
    }
}

func main() {

    name := "Alice"
    price := 279
    city := "London"

    swap(&name, &city)
    for _, val := range []interface{} { name, price, city }  {
        Printfln("Value: %v", val)
    }
}
