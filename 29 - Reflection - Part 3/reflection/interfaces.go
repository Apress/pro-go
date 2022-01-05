package main

import "fmt"

type NamedItem interface {
    GetName() string
    unexportedMethod()
}

type CurrencyItem interface {
    GetAmount() string
    currencyName() string 
}

func (p *Product) GetName() string {
    return p.Name
}

func (c *Customer) GetName() string {
    return c.Name
}

func (p *Product) GetAmount() string {
    return fmt.Sprintf("$%.2f", p.Price)
}

func (p *Product) currencyName() string {
    return "USD"
}

func (p *Product) unexportedMethod() {}
