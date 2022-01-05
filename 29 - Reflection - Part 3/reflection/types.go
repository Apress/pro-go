package main

type Product struct {
    Name, Category string
    Price float64
}

type Customer struct {
    Name, City string
}

type Purchase struct {
    Customer
    Product
    Total float64
    taxRate float64
}
