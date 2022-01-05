package store

import (
    "sportsstore/models"
    "platform/http/actionresults"
    "net/http"
)

type StatusCodeResult struct {
    code int 
}

func (action *StatusCodeResult) Execute(ctx *actionresults.ActionContext) error {
    ctx.ResponseWriter.WriteHeader(action.code)
    return nil
}

type RestHandler struct {
    Repository models.Repository
}

func (h RestHandler) GetProduct(id int) actionresults.ActionResult {
    return actionresults.NewJsonAction(h.Repository.GetProduct(id))
}

func (h RestHandler) GetProducts() actionresults.ActionResult {
    return actionresults.NewJsonAction(h.Repository.GetProducts())
}

type ProductReference struct {
    models.Product
    CategoryID int
}

func (h RestHandler) PostProduct(p ProductReference) actionresults.ActionResult {
    if p.ID == 0 {
        return actionresults.NewJsonAction(h.processData(p))
    } else {
        return &StatusCodeResult{ http.StatusBadRequest }
    }
}

func (h RestHandler) PutProduct(p ProductReference) actionresults.ActionResult {
    if p.ID > 0 {
        return actionresults.NewJsonAction(h.processData(p))
    } else {
        return &StatusCodeResult{ http.StatusBadRequest }
    }
}

func (h RestHandler) processData(p ProductReference) models.Product {
    product := p.Product
    product.Category = &models.Category {
        ID: p.CategoryID,
    }
    h.Repository.SaveProduct(&product)  
    return h.Repository.GetProduct(product.ID)
}
