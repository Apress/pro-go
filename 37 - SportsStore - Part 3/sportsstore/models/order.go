package models

type Order struct {
    ID int
    ShippingDetails
    Products []ProductSelection
    Shipped bool
}

type ShippingDetails struct {
    Name string `validation:"required"`
    StreetAddr string `validation:"required"`
    City string `validation:"required"`
    State string `validation:"required"`
    Zip string `validation:"required"`
    Country string `validation:"required"`
}

type ProductSelection struct{
    Quantity int
    Product
}
