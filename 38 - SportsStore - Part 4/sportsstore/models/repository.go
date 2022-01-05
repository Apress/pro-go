package models

type Repository interface {

    GetProduct(id int) Product
    GetProducts() []Product
    SaveProduct(*Product)

    GetProductPage(page, pageSize int) (products []Product, totalAvailable int)

    GetProductPageCategory(categoryId int, page, pageSize int) (products []Product, 
        totalAvailable int)
    
    GetCategories() []Category
    SaveCategory(*Category)

    GetOrder(id int) Order
    GetOrders() []Order
    SaveOrder(*Order) 
    SetOrderShipped(*Order) 
 
    Seed()
    Init()
}
