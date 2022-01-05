DROP TABLE IF EXISTS Categories;
DROP TABLE IF EXISTS Products;

CREATE TABLE IF NOT EXISTS Categories (
    Id INTEGER NOT NULL PRIMARY KEY,			
    Name TEXT
);

CREATE TABLE IF NOT EXISTS Products (
    Id INTEGER NOT NULL PRIMARY KEY,
    Name TEXT,
    Category INTEGER,
    Price decimal(8, 2),
    CONSTRAINT CatRef FOREIGN KEY(Category) REFERENCES Categories (Id)
);

INSERT INTO Categories (Id, Name) VALUES
    (1, "Watersports"),
    (2, "Soccer");

INSERT INTO Products (Id, Name, Category, Price) VALUES
    (1, "Kayak", 1, 279),
    (2, "Lifejacket", 1, 48.95),
    (3, "Soccer Ball", 2, 19.50),
    (4, "Corner Flags", 2, 34.95);
