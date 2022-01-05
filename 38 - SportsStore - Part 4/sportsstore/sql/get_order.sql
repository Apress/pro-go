SELECT Orders.Id, Orders.Name, Orders.StreetAddr, Orders.City, Orders.Zip, 
    Orders.Country, Orders.Shipped
FROM Orders
WHERE Orders.Id = ?
