package reverseredemption

import (
	"errors"
	"strings"
	"time"
	"txn/helper/constant"
	structDB "txn/structs/db"

	"github.com/astaxie/beego/orm"
)

//Reverse
type Reverse struct{}

// GetAllEarningHistoryAndSortIt ...
func (b *Reverse) GetAllEarningHistoryAndSortIt(filename string) (
	res []structDB.ReverseRedemptionBatchHistory,
	err error) {

	responseErr := "|0||"
	if strings.Contains(filename, "Fail") {
		responseErr = "|1|error_code|err_desc"
	}
	res = append(res,
		structDB.ReverseRedemptionBatchHistory{
			JobID:      "123",
			Filename:   "",
			ID:         1,
			Status:     2,
			TextString: "D|180209C734D47006045677|180208C7AAD47002067492|50000000005|30.00|EX1" + responseErr,
		},
		structDB.ReverseRedemptionBatchHistory{
			JobID:      "456",
			Filename:   "",
			ID:         2,
			Status:     2,
			TextString: "D|180209C567D47006098756|180208C787D47002089450|50000000008|10.00|EX1" + responseErr,
		},
	)

	return
}

// RemoveHistory Method to Clear Reverse Redemption History based on Filename
func (b *Reverse) RemoveHistory(filename string) error {

	return nil
}

// BulkInsert Bulking Insert
func (b *Reverse) BulkInsert(theRows []string, filename string) error {

	return nil
}

// FindTxnlogByJobID method to get original transaction log
func (b *Reverse) FindTxnlogByJobID(o orm.Ormer, jobID string) (res []structDB.TypeTxnDB, err error) {
	if jobID == "originalJobIDNotFound" {

		return
	}

	if jobID == "jobIDforFailInsertTxnLog" {
		res = []structDB.TypeTxnDB{
			{
				TxnDesc: "desc1",
				ReqID:   "",
			},
			{
				TxnDesc: "desc2",
				ReqID:   "",
			},
		}
		return
	}

	res = []structDB.TypeTxnDB{
		{
			TxnDesc: "desc1",
			ReqID:   jobID,
		},
		{
			TxnDesc: "desc2",
			ReqID:   jobID,
		},
	}

	return
}

// SetIsReverseTxnlog method to update is_reverse in original transaction of trx_log and trx_history
func (b *Reverse) SetIsReverseTxnlog(o orm.Ormer, jobID string, isReverse bool) error {
	if jobID == "jobIDfailReverse" {
		return errors.New("fail update is_reverse")
	}
	return nil
}

// UpdateTransactionAccount ...
func (b *Reverse) UpdateTransactionAccount(o orm.Ormer, accountNumber string,
	transactionAmount float64, txnTime time.Time,
	returningUpdateAccount *structDB.TypeReturningUpdateAccount) error {

	if accountNumber == "accForFailUpdateAccount" {
		return errors.New("fail update account")
	}
	(*returningUpdateAccount) = structDB.TypeReturningUpdateAccount{
		LedgerBalance:           transactionAmount,
		LedgerBalanceLastUpdate: txnTime,
		AccountNumber:           accountNumber,
		TotalCreditMonthToDate:  transactionAmount,
	}
	return nil
}

// FindAccountByAccountNumber ...
func (b *Reverse) FindAccountByAccountNumber(o orm.Ormer, accountNumber string) (
	txnAccount []structDB.TransactionAccount, err error,
) {
	if accountNumber == "" {
		err = errors.New("failed to get account")
		return
	}

	if accountNumber == "accNumberNotFound" {
		return
	}

	if accountNumber == "1234AlreadyClosed" {
		txnAccount = append(txnAccount, structDB.TransactionAccount{
			AccountNumber: accountNumber,
			AccountBranch: 1,
			LedgerBalance: 1000000,
			AccountStatus: constant.AccountStatusClosed,
			ProductCode:   constant.LoyaltyProductCode,
		})
		return
	}

	txnAccount = append(txnAccount, structDB.TransactionAccount{
		AccountNumber: accountNumber,
		AccountBranch: 1,
		LedgerBalance: 1000000,
		AccountStatus: constant.AccountStatusActive,
		ProductCode:   constant.LoyaltyProductCode,
	})

	return
}

// UpdateReverseHistory Update reverse_redemption_batch_history
func (b *Reverse) UpdateReverseHistory(
	o orm.Ormer,
	status int,
	val string,
	jobID string,
) error {

	return nil
}
