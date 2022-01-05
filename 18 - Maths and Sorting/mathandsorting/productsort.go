package main 

import "sort"

type Product struct {
    Name string
    Price float64
}

type ProductSlice []Product

func ProductSlices(p []Product) {
    sort.Sort(ProductSlice(p))
}

func ProductSlicesAreSorted(p []Product) {
    sort.IsSorted(ProductSlice(p))
}

func (products ProductSlice) Len() int {
    return len(products)
}

func (products ProductSlice) Less(i, j int) bool {
    return products[i].Price < products[j].Price
}

func (products ProductSlice) Swap(i, j int) {
    products[i], products[j] = products[j], products[i]
}

type ProductSliceName struct { ProductSlice }

func ProductSlicesByName(p []Product) {
    sort.Sort(ProductSliceName{ p })
}

func (p ProductSliceName) Less(i, j int) bool {
    return p.ProductSlice[i].Name < p.ProductSlice[j].Name
}

type ProductComparison func(p1, p2 Product) bool

type ProductSliceFlex struct { 
    ProductSlice 
    ProductComparison
}

func (flex ProductSliceFlex) Less(i, j int) bool {
    return flex.ProductComparison(flex.ProductSlice[i], flex.ProductSlice[j])
}

func SortWith(prods []Product, f ProductComparison) {
    sort.Sort(ProductSliceFlex{ prods, f})
}
