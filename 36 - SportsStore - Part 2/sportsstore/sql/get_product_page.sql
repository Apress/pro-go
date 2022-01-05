SELECT Products.Id, Products.Name, Products.Description, Products.Price, 
    Categories.Id, Categories.Name 
FROM Products, Categories 
WHERE Products.Category = Categories.Id	
ORDER BY Products.Id
LIMIT ? OFFSET ?
