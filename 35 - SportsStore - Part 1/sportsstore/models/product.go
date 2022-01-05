package models

type Product struct {
    ID int
    Name string
    Description string 
    Price float64       
    *Category
}
