package store

type ItemForSale interface {
    Price(taxRate float64) float64
}
