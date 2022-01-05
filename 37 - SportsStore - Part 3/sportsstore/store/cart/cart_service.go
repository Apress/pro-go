package cart

import (
    "platform/services"
    "platform/sessions"
    "sportsstore/models"
    "encoding/json"
    "strings"
)

const CART_KEY string = "cart"

func RegisterCartService() {
    services.AddScoped(func(session sessions.Session) Cart {
        lines := []*CartLine {}
        sessionVal := session.GetValue(CART_KEY)
        if strVal, ok := sessionVal.(string); ok {
            json.NewDecoder(strings.NewReader(strVal)).Decode(&lines)
        }
        return &sessionCart{ 
            BasicCart: &BasicCart{ lines: lines},
            Session: session,
        }
    })
}

type sessionCart struct {
    *BasicCart
    sessions.Session
}

func (sc *sessionCart) AddProduct(p models.Product) {
    sc.BasicCart.AddProduct(p)
    sc.SaveToSession()
}

func (sc *sessionCart) RemoveLineForProduct(id int) {
    sc.BasicCart.RemoveLineForProduct(id)
    sc.SaveToSession()
}

func (sc *sessionCart) SaveToSession() {
    builder := strings.Builder{}
    json.NewEncoder(&builder).Encode(sc.lines)
    sc.Session.SetValue(CART_KEY, builder.String())
}

func (sc *sessionCart) Reset() {
    sc.lines = []*CartLine{}
    sc.SaveToSession()
}
