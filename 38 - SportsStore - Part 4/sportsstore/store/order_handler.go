package store

import (
	"encoding/json"
	"platform/http/actionresults"
	"platform/http/handling"
	"platform/sessions"
	"platform/validation"
	"sportsstore/models"
	"sportsstore/store/cart"
	"strings"
)

type OrderHandler struct {
    cart.Cart
    sessions.Session
    Repository models.Repository
    URLGenerator handling.URLGenerator 
    validation.Validator
}

type OrderTemplateContext struct {
    models.ShippingDetails
    ValidationErrors [][]string
    CancelUrl string   
}

func (handler OrderHandler) GetCheckout() actionresults.ActionResult {
    context := OrderTemplateContext {}
    jsonData := handler.Session.GetValueDefault("checkout_details", "")
    if jsonData != nil {
        json.NewDecoder(strings.NewReader(jsonData.(string))).Decode(&context)
    }
    context.CancelUrl = mustGenerateUrl(handler.URLGenerator, CartHandler.GetCart)
    return actionresults.NewTemplateAction("checkout.html", context)
}

func (handler OrderHandler) PostCheckout(details models.ShippingDetails) actionresults.ActionResult {
    valid, errors := handler.Validator.Validate(details)
    if (!valid) {
        ctx := OrderTemplateContext {
            ShippingDetails: details,
            ValidationErrors: [][]string {},
        }   
        for _, err := range errors {
            ctx.ValidationErrors = append(ctx.ValidationErrors, 
                []string { err.FieldName, err.Error.Error()})
        }

        builder := strings.Builder{}
        json.NewEncoder(&builder).Encode(ctx)
        handler.Session.SetValue("checkout_details", builder.String())
        redirectUrl := mustGenerateUrl(handler.URLGenerator, 
            OrderHandler.GetCheckout)
        return actionresults.NewRedirectAction(redirectUrl)
    } else {
        handler.Session.SetValue("checkout_details", "")
    }
    order := models.Order { 
        ShippingDetails: details, 
        Products: []models.ProductSelection {},
    }
    for _, cl := range handler.Cart.GetLines() {
        order.Products = append(order.Products, models.ProductSelection {
            Quantity: cl.Quantity,
            Product: cl.Product,
        })
    }
    handler.Repository.SaveOrder(&order)
    handler.Cart.Reset()
    targetUrl, _ := handler.URLGenerator.GenerateUrl(OrderHandler.GetSummary, 
        order.ID)
    return actionresults.NewRedirectAction(targetUrl)
}

func (handler OrderHandler) GetSummary(id int) actionresults.ActionResult {
    targetUrl, _ := handler.URLGenerator.GenerateUrl(ProductHandler.GetProducts, 
        0, 1)
    return actionresults.NewTemplateAction("checkout_summary.html", struct {
        ID int
        TargetUrl string
    }{ ID: id, TargetUrl: targetUrl})
}
