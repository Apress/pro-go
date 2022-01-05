package main

import (
    "sync"
    "platform/http"
    "platform/http/handling"
    "platform/services"
    "platform/pipeline"
    "platform/pipeline/basic"
    "sportsstore/store"
    "sportsstore/models/repo"
    "platform/sessions"
    "sportsstore/store/cart"
    "sportsstore/admin"
    "platform/authorization"
    "sportsstore/admin/auth"
)

func registerServices() {
    services.RegisterDefaultServices()
    //repo.RegisterMemoryRepoService()
    repo.RegisterSqlRepositoryService()
    sessions.RegisterSessionService()
    cart.RegisterCartService()
    authorization.RegisterDefaultSignInService()
    authorization.RegisterDefaultUserService()
    auth.RegisterUserStoreService()
}

func createPipeline() pipeline.RequestPipeline {
    return pipeline.CreatePipeline(
        &basic.ServicesComponent{},
        &basic.LoggingComponent{},
        &basic.ErrorComponent{},
        &basic.StaticFileComponent{},
        &sessions.SessionComponent{},


        authorization.NewAuthComponent(
            "admin", 
            authorization.NewRoleCondition("Administrator"),
            admin.AdminHandler{},            
            admin.ProductsHandler{},
            admin.CategoriesHandler{},           
            admin.OrdersHandler{},            
            admin.DatabaseHandler{},     
            admin.SignOutHandler{},
        ).AddFallback("/admin/section/", "^/admin[/]?$"),
        
        handling.NewRouter(
            handling.HandlerEntry{ "",  store.ProductHandler{}},
            handling.HandlerEntry{ "",  store.CategoryHandler{}},
            handling.HandlerEntry{ "", store.CartHandler{}},            
            handling.HandlerEntry{ "", store.OrderHandler{}},            
            // handling.HandlerEntry{ "admin", admin.AdminHandler{}},            
            // handling.HandlerEntry{ "admin", admin.ProductsHandler{}},   
            // handling.HandlerEntry{ "admin", admin.CategoriesHandler{}},           
            // handling.HandlerEntry{ "admin", admin.OrdersHandler{}},            
            // handling.HandlerEntry{ "admin", admin.DatabaseHandler{}},                  
            handling.HandlerEntry{ "", admin.AuthenticationHandler{}},
            handling.HandlerEntry{ "api", store.RestHandler{}},
        ).AddMethodAlias("/", store.ProductHandler.GetProducts, 0, 1).
            AddMethodAlias("/products[/]?[A-z0-9]*?", 
                store.ProductHandler.GetProducts, 0, 1),    )
}

func main() {
    registerServices()
    results, err := services.Call(http.Serve, createPipeline())
    if (err == nil) {
        (results[0].(*sync.WaitGroup)).Wait()
    } else {
        panic(err)
    }    
}
