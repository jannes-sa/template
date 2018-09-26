package common

import (
	//"errors"
	structPg "txn/structs/db"
)

// Internal ...
type DBAccount struct{}

 //FindAccountByAccountNumber ...
func (i *DBAccount) FindAccountByAccountNumber(accountNumber string) (
	account structPg.TransactionAccount, err error,
){
	account.AccountNumber = accountNumber
	account.AccountBranch = 1
	account.LedgerBalance = 1000000
	return
}