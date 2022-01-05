package repo

import (
    //"platform/services"
    "sportsstore/models"
    "math"
)

// func RegisterMemoryRepoService() {
//     services.AddSingleton(func() models.Repository {
//         repo := &MemoryRepo{}
//         repo.Seed()
//         return repo
//     })
// }

type MemoryRepo struct {
    products []models.Product
    categories []models.Category
}

func (repo *MemoryRepo) GetProduct(id int) (product models.Product) {
    for _, p := range repo.products {
        if (p.ID == id) {
            product = p
            return
        }
    }
    return
} 

func (repo *MemoryRepo) GetProducts() (results []models.Product) {
    return repo.products
}

func (repo *MemoryRepo) GetCategories() (results []models.Category) {
    return repo.categories
}

func (repo *MemoryRepo) GetProductPage(page, pageSize int) ([]models.Product, int) {
    return getPage(repo.products, page, pageSize), len(repo.products)
}

func (repo *MemoryRepo) GetProductPageCategory(category int, page, 
        pageSize int) (products []models.Product, totalAvailable int) {
    if category == 0 {
        return repo.GetProductPage(page, pageSize)
    } else {
        filteredProducts := make([]models.Product, 0, len(repo.products))
        for _, p := range repo.products {
            if p.Category.ID == category {
                filteredProducts = append(filteredProducts, p)
            }
        }
        return getPage(filteredProducts, page, pageSize), len(filteredProducts)
    }
}

func getPage(src []models.Product, page, pageSize int) []models.Product {
    start := (page -1) * pageSize
    if page > 0 && len(src) > start {
        end := (int)(math.Min((float64)(len(src)), (float64)(start + pageSize)))
        return src[start : end]
    } 
    return []models.Product{}
}
