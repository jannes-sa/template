package txnInternal

import (
	"errors"
	structPg "txn/structs/db"
)

// Internal ...
type Internal struct{}

// FindAccountByAccountNumber ...
func (i *Internal) FindAccountByAccountNumber(accountNumber string) (
	account structPg.TransactionAccount, err error) {
	if accountNumber == "" {
		err = errors.New("Failed get account number because accNumber empty")
		return
	}
	account.AccountNumber = accountNumber
	account.AccountBranch = 1
	account.LedgerBalance = 1000000
	err = nil
	return
}

// InsertTransactionLog ...
func (i *Internal) InsertTransactionLog(
	txnDBRows []structPg.TypeTxnDB) error {
	if len(txnDBRows) == 0 {
		err := errors.New("error transaction log")
		return err
	}
	return nil
}
