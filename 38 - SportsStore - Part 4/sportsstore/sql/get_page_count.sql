SELECT COUNT (Products.Id)
FROM Products, Categories 
WHERE Products.Category = Categories.Id;
