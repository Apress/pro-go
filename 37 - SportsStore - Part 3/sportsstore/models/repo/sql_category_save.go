package repo

import "sportsstore/models"

func (repo *SqlRepository) SaveCategory(c *models.Category) {
    if (c.ID == 0) {
        result, err := repo.Commands.SaveCategory.ExecContext(repo.Context, 
            c.CategoryName)
        if err == nil {
            id, err := result.LastInsertId()
            if err == nil {
                c.ID = int(id)
                return
            } else {
                repo.Logger.Panicf("Cannot get inserted ID: %v", err.Error())       
            }
        } else {
            repo.Logger.Panicf("Cannot exec SaveCategory command: %v", err.Error())            
        }        
    } else {
        result, err := repo.Commands.UpdateCategory.ExecContext(repo.Context, 
            c.CategoryName, c.ID)
        if err == nil {
            affected, err := result.RowsAffected()
            if err == nil && affected != 1 {
                repo.Logger.Panicf("Got unexpected rows affected: %v", affected)       
            } else if err != nil {
                repo.Logger.Panicf("Cannot get rows affected: %v", err)       
            }
        } else {
            repo.Logger.Panicf("Cannot exec UpdateCategory command: %v", err.Error())                   
        }       
    }
}
