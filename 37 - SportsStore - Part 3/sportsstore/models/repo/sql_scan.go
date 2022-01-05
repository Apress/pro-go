package repo

import (
    "database/sql"
    "sportsstore/models"
)

func scanProducts(rows *sql.Rows) (products []models.Product, err error) {
    products = make([]models.Product, 0, 10)
    for rows.Next() {
        p := models.Product{ Category: &models.Category{}}
        err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, 
            &p.Category.ID, &p.Category.CategoryName)
        if (err == nil) {
            products = append(products, p)
        } else {
            return
        }
    }
    return
}

func scanProduct(row *sql.Row) (p models.Product, err error) {
    p = models.Product{ Category: &models.Category{}}
    err = row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category.ID, 
        &p.Category.CategoryName)
    return p, err
}
