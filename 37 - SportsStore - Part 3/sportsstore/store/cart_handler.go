package store

import (
    "platform/http/actionresults"
    "platform/http/handling"
    "sportsstore/models"
    "sportsstore/store/cart"
)

type CartHandler struct {
    models.Repository
    cart.Cart
    handling.URLGenerator
}

type CartTemplateContext struct {
    cart.Cart
    ProductListUrl string
    CartUrl string
    CheckoutUrl string
    RemoveUrl string
}

func (handler CartHandler) GetCart() actionresults.ActionResult {
    return actionresults.NewTemplateAction("cart.html", CartTemplateContext {
        Cart: handler.Cart,
        ProductListUrl: handler.mustGenerateUrl(ProductHandler.GetProducts, 0, 1),
        RemoveUrl: handler.mustGenerateUrl(CartHandler.PostRemoveFromCart),
        CheckoutUrl: handler.mustGenerateUrl(OrderHandler.GetCheckout),                
    })
}

func (handler CartHandler) GetWidget() actionresults.ActionResult {
    return actionresults.NewTemplateAction("cart_widget.html", CartTemplateContext {
        Cart: handler.Cart,
        CartUrl: handler.mustGenerateUrl(CartHandler.GetCart),
    })
}

type CartProductReference struct {
    ID int
}

func (handler CartHandler) PostAddToCart(ref CartProductReference) actionresults.ActionResult {
    p := handler.Repository.GetProduct(ref.ID)
    handler.Cart.AddProduct(p)
    return actionresults.NewRedirectAction(
        handler.mustGenerateUrl(CartHandler.GetCart))
}

func (handler CartHandler) PostRemoveFromCart(ref CartProductReference) actionresults.ActionResult {
    handler.Cart.RemoveLineForProduct(ref.ID)
    return actionresults.NewRedirectAction(
        handler.mustGenerateUrl(CartHandler.GetCart))
}

func (handler CartHandler) mustGenerateUrl(method interface{}, data ...interface{}) string {
    url, err := handler.URLGenerator.GenerateUrl(method, data...)
    if (err != nil) {
        panic(err)
    }
    return url
}
