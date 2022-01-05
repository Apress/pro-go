INSERT INTO Categories(Id, Name) VALUES 
	(1, "Watersports"), (2, "Soccer"), (3, "Chess");
	
INSERT INTO Products(Id, Name, Description, Category, Price) VALUES
	(1, "Kayak", "A boat for one person", 1, 275),
	(2, "Lifejacket", "Protective and fashionable", 1, 48.95),
	(3, "Soccer Ball", "FIFA-approved size and weight", 2, 19.50),
	(4, "Corner Flags", "Give your playing field a professional touch", 2, 34.95),	
	(5, "Stadium", "Flat-packed 35,000-seat stadium", 2, 79500),	
	(6, "Thinking Cap", "Improve brain efficiency by 75%", 3, 16),	
	(7, "Unsteady Chair", "Secretly give your opponent a disadvantage", 3, 29.95),	
	(8, "Human Chess Board", "A fun game for the family", 3, 75),	
	(9, "Bling-Bling King", "Gold-plated, diamond-studded King", 3, 1200);

INSERT INTO Orders(Id, Name, StreetAddr, City, Zip, Country, Shipped) VALUES
	(1, "Alice", "123 Main St", "New Town", "12345", "USA", false),
	(2, "Bob", "The Grange", "Upton", "UP12 6YT", "UK", false);

INSERT INTO OrderLines(Id, OrderId, ProductId, Quantity) VALUES
	(1, 1, 1, 1), (2, 1, 2, 2), (3, 1, 8, 1), (4, 2, 5, 2);
