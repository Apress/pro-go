package repo

import "sportsstore/models"

func (repo *SqlRepository) SetOrderShipped(o *models.Order) {
    result, err := repo.Commands.UpdateOrder.ExecContext(repo.Context, 
        o.Shipped, o.ID)
    if err == nil {
        rows, err :=result.RowsAffected()
        if err != nil {
            repo.Logger.Panicf("Cannot get updated ID: %v", err.Error())     
        } else if rows != 1 {
            repo.Logger.Panicf("Got unexpected rows affected: %v", rows)       
        }
    } else {
        repo.Logger.Panicf("Cannot exec UpdateOrder command: %v", err.Error())               
    }
}
