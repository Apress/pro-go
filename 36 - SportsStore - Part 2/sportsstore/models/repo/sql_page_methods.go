package repo

import "sportsstore/models"

func (repo *SqlRepository) GetProductPage(page, 
        pageSize int) (products []models.Product, totalAvailable int) {
    rows, err := repo.Commands.GetPage.QueryContext(repo.Context, 
        pageSize, (pageSize * page) - pageSize)
    if err == nil {
        if products, err = scanProducts(rows); err != nil {
            repo.Logger.Panicf("Cannot scan data: %v", err.Error())    
            return
        }
    } else {
        repo.Logger.Panicf("Cannot exec GetProductPage command: %v", err)
        return 
    }
    row := repo.Commands.GetPageCount.QueryRowContext(repo.Context)
    if row.Err() == nil {
        if err := row.Scan(&totalAvailable); err != nil {
            repo.Logger.Panicf("Cannot scan data: %v", err.Error())    
        }
    } else {
        repo.Logger.Panicf("Cannot exec GetPageCount command: %v", row.Err().Error())
    }
    return
}

func (repo *SqlRepository) GetProductPageCategory(categoryId int, page, 
        pageSize int) (products []models.Product, totalAvailable int) {
    if (categoryId == 0) {
        return repo.GetProductPage(page, pageSize)
    }
    rows, err := repo.Commands.GetCategoryPage.QueryContext(repo.Context, categoryId, 
        pageSize, (pageSize * page) - pageSize)
    if err == nil {
        if products, err = scanProducts(rows); err != nil {
            repo.Logger.Panicf("Cannot scan data: %v", err.Error())    
            return
        }
    } else {
        repo.Logger.Panicf("Cannot exec GetProductPage command: %v", err)
        return         
    }
    row := repo.Commands.GetCategoryPageCount.QueryRowContext(repo.Context, 
        categoryId)
    if row.Err() == nil {
        if err := row.Scan(&totalAvailable); err != nil {
            repo.Logger.Panicf("Cannot scan data: %v", err.Error())    
        }        
    } else {
        repo.Logger.Panicf("Cannot exec GetCategoryPageCount command: %v", 
            row.Err().Error())
    }
    return
}
