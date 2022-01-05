package repo

import "sportsstore/models"

func (repo *SqlRepository) SaveProduct(p *models.Product) {

    if (p.ID == 0) {
        result, err := repo.Commands.SaveProduct.ExecContext(repo.Context, p.Name, 
            p.Description, p.Category.ID, p.Price)
        if err == nil {
            id, err := result.LastInsertId()
            if err == nil {
                p.ID = int(id)
                return
            } else {
                repo.Logger.Panicf("Cannot get inserted ID: %v", err.Error())       
            }
        } else {
            repo.Logger.Panicf("Cannot exec SaveProduct command: %v", err.Error())            
        }
    } else {
        result, err := repo.Commands.UpdateProduct.ExecContext(repo.Context, p.Name, 
            p.Description, p.Category.ID, p.Price, p.ID)
        if err == nil {
            affected, err := result.RowsAffected()
            if err == nil && affected != 1 {
                repo.Logger.Panicf("Got unexpected rows affected: %v", affected)       
            } else if err != nil {
                repo.Logger.Panicf("Cannot get rows affected: %v", err)       
            }
        } else {
            repo.Logger.Panicf("Cannot exec Update command: %v", err.Error())                   
        }
    }
}
