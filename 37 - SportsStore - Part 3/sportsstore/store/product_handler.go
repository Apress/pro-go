package store

import (
    "sportsstore/models"
    "platform/http/actionresults"
    "platform/http/handling"
    "math"
)

const pageSize = 4

type ProductHandler struct {
    Repository models.Repository
    URLGenerator handling.URLGenerator
}

type ProductTemplateContext struct {
    Products []models.Product
    Page int
    PageCount int
    PageNumbers []int
    PageUrlFunc func(int) string
    SelectedCategory int
    AddToCartUrl string
}

func (handler ProductHandler) GetProducts(category, 
        page int) actionresults.ActionResult {
    prods, total := handler.Repository.GetProductPageCategory(category, 
        page, pageSize)
    pageCount := int(math.Ceil(float64(total) / float64(pageSize)))
    return actionresults.NewTemplateAction("product_list.html", 
        ProductTemplateContext {
            Products: prods,
            Page: page,
            PageCount: pageCount,
            PageNumbers: handler.generatePageNumbers(pageCount),
            PageUrlFunc: handler.createPageUrlFunction(category),
            SelectedCategory: category,
            AddToCartUrl: mustGenerateUrl(handler.URLGenerator, 
                 CartHandler.PostAddToCart),
        })     
}

func (handler ProductHandler) createPageUrlFunction(category int) func(int) string {
    return func(page int) string {
        url, _ := handler.URLGenerator.GenerateUrl(ProductHandler.GetProducts, 
            category, page)
        return url
    }
}

func (handler ProductHandler) generatePageNumbers(pageCount int) (pages []int) {
    pages = make([]int, pageCount)
    for i := 0; i < pageCount; i++ {
        pages[i] = i + 1
    }
    return
}

func mustGenerateUrl(generator handling.URLGenerator, target interface{}) string {
    url, err := generator.GenerateUrl(target)    
    if (err != nil) {
        panic(err)
    }
    return url;
}
