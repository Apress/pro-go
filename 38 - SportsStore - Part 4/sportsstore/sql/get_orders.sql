SELECT Orders.Id, Orders.Name, Orders.StreetAddr, Orders.City, Orders.Zip, Orders.Country, Orders.Shipped
FROM Orders
ORDER BY Orders.Shipped, Orders.Id
