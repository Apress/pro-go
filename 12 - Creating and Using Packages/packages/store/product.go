package store

var standardTax = newTaxRate(0.25, 20)

type Product struct {
    Name, Category string
    price float64 
}

func NewProduct(name, category string, price float64) *Product {
    return &Product{ name, category, price }
}

func (p *Product) Price() float64 {
    return standardTax.calcTax(p)
}

func (p *Product) SetPrice(newPrice float64)  {
    p.price = newPrice
}
