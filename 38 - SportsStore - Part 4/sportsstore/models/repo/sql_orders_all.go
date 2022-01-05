package repo

import "sportsstore/models"

func (repo *SqlRepository) GetOrders() []models.Order {
    orderMap := make(map[int]*models.Order, 10)
    orderRows, err := repo.Commands.GetOrders.QueryContext(repo.Context)
    if err != nil {
        repo.Logger.Panicf("Cannot exec GetOrders command: %v", err.Error())
    }
    for orderRows.Next() {
        order := models.Order { Products: []models.ProductSelection {}}
        err := orderRows.Scan(&order.ID, &order.Name, &order.StreetAddr, &order.City, 
            &order.Zip, &order.Country, &order.Shipped)
        if (err != nil) {
            repo.Logger.Panicf("Cannot scan order data: %v", err.Error())
            return  []models.Order {}
        }   
        orderMap[order.ID] = &order
    }

    lineRows, err := repo.Commands.GetOrdersLines.QueryContext(repo.Context)
    if (err != nil) {
        repo.Logger.Panicf("Cannot exec GetOrdersLines command: %v", err.Error())
    }
    for lineRows.Next() {
        var order_id int
        ps := models.ProductSelection { 
            Product: models.Product{ Category: &models.Category{} },
        }
        err = lineRows.Scan(&order_id, &ps.Quantity, &ps.Product.ID,          
            &ps.Product.Name, &ps.Product.Description, &ps.Product.Price, 
            &ps.Product.Category.ID, &ps.Product.Category.CategoryName)
        if err == nil {
            orderMap[order_id].Products = append(orderMap[order_id].Products, ps)
        } else {
            repo.Logger.Panicf("Cannot scan order line data: %v", err.Error())
        }        
    }
    orders := make([]models.Order, 0, len(orderMap))
    for _, o := range orderMap {
        orders = append(orders, *o)
    }
    return orders
}
