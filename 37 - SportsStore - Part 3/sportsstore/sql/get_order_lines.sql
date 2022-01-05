SELECT OrderLines.Quantity, Products.Id, Products.Name, Products.Description, 
    Products.Price, Categories.Id, Categories.Name
FROM Orders, OrderLines, Products, Categories
WHERE Orders.Id = OrderLines.OrderId 
    AND OrderLines.ProductId = Products.Id 
    AND Products.Category = Categories.Id
    AND Orders.Id = ?
ORDER BY Products.Id
