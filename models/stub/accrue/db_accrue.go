package accrue

import (
	"txn/structs"
	structPg "txn/structs/db"
	"txn/helper/timetn"
)

// DBAccrue ...
type DBAccrue struct{}

// This stub return [] if the input accrualDate != now
// GetAccrueData ...
func (b *DBAccrue) GetAccrueData(accrualDate string,errCode *[]structs.TypeError)([]structs.ItemAccruAccount,error){
	res := make([]structs.ItemAccruAccount, 0)

	if accrualDate==timetn.Now().Format("2006-01-02") {
		res = append(res, structs.ItemAccruAccount{
			AccountNumber:   "52345678901",
			CifCustomerType: "1010",
			ProductCode:     "SA001",
			Currency:        "THB",
			AccountBranch:   1,
			AccountStatus:   1,
			LedgerBalance:   100.0,
		})
	}
	return res,nil
}

func (b *DBAccrue) MigrateAccountForAccru(name string,account []structPg.TransactionAccount) {

	return

}

func (b *DBAccrue) ClearData(){
	return
}