package DBPost

import (
	"errors"
	"time"
	"txn/structs"
	structDB "txn/structs/db"

	"txn/helper"
	"txn/helper/constant"

	"github.com/astaxie/beego/orm"
)

// DBPost ...
type DBPost struct{}

// InsertEarningHistory ...
func (b *DBPost) InsertEarningHistory(
	o orm.Ormer,
	status int,
	id string,
	val string,
) error {
	return nil
}

// GetAllEarningHistory ...
func (b *DBPost) GetAllEarningHistory() (
	res []structDB.EarningBatchHistory,
	err error) {
	res = append(res, structDB.EarningBatchHistory{
		ID:         1,
		Status:     1,
		TextString: "",
	}, structDB.EarningBatchHistory{
		ID:         2,
		Status:     1,
		TextString: "",
	})

	err = nil
	return
}

// CreateTransactionLog ...
func (b *DBPost) CreateTransactionLog(
	txLogs []structDB.TypeTxnDB,
	o orm.Ormer,
) (err error) {
	if txLogs[0].ReqID == "" {
		err = errors.New("error ReqID")
	}
	return
}

// CreateHistoryTrx ...
func (b *DBPost) CreateHistoryTrx(
	txHistory structDB.TransactionHistoryDB,
	o orm.Ormer,
) (err error) {
	var timeCompare time.Time
	if txHistory.TxnTime == timeCompare {
		err = errors.New("error txnTime")
	}
	return
}

// CheckErrorForInsertWithoutID ...
func (b *DBPost) CheckErrorForInsertWithoutID(
	err error,
	msg string,
	stateTx *bool,
	o orm.Ormer,
	errCode *[]structs.TypeError,
) {
	(*stateTx) = true
}

// GetAllEarningHistoryStatusPending ...
func (b *DBPost) GetAllEarningHistoryStatusPending() (
	res []structDB.EarningBatchHistory,
	err error,
) {
	err = nil
	return
}

// UpdateEarningHistory ...
func (b *DBPost) UpdateEarningHistory(
	o orm.Ormer,
	status int,
	val string,
	jobID string,
) error {
	return nil
}

func (b *DBPost) FindAccountByAccountNumber(accountNumber string) (
	account structDB.TransactionAccount, err error) {

	account = structDB.TransactionAccount{
		AccountNumber: accountNumber,
		AccountBranch: 1,
		LedgerBalance: 1000000,
	}
	return
}

func (b *DBPost) CreateOrUpdateEODBalance(date time.Time,
	accountNumber string,
	amount float64,
	o orm.Ormer,
) error {
	var err error
	if accountNumber == "" {
		err = errors.New("account number cannot be empty")
	}
	return err

}

// UpdateTransactionLogError ..
func (b *DBPost) UpdateTransactionLogError(jobID string, txnStatus int,
	errCode string, errDesc string) (err error) {
	if errCode == "" || errDesc == "" {
		err = errors.New("errCode or errDesc should not empty @UpdateTransactionLogError")
	}
	return
}

//FindTxnLogByRequestID ...
func (b *DBPost) FindTxnLogByRequestID(jobID string, status int) (res []structDB.TypeGetTxnLogByJobIDAndStatus, err error) {

	if jobID == "1234" {
		res = append(res, structDB.TypeGetTxnLogByJobIDAndStatus{
			JobID:             "1234",
			TransactionStatus: status,
			DrCr:              "DR",
			TransactionAmount: 10000,
			AccountNumber:     "1001",
			ServiceBranch:     1,
			FromAccountBranch: 1,
			ToAccountBranch:   1,
			OriginationDomain: constant.DOMAINID,
			DestinationDomain: constant.DOMAINID,
		})

		res = append(res, structDB.TypeGetTxnLogByJobIDAndStatus{
			JobID:             "1234",
			TransactionStatus: status,
			DrCr:              "CR",
			TransactionAmount: 10000,
			AccountNumber:     "1001",
			ServiceBranch:     1,
			FromAccountBranch: 1,
			ToAccountBranch:   1,
			OriginationDomain: constant.DOMAINID,
			DestinationDomain: constant.DOMAINID,
		})
	} else {
		err = errors.New("TnxLog " + jobID + " not found ")
		helper.CheckErr("Error@FindTxnLogByRequestID", err)
	}

	return
}

func (b *DBPost) UpdateTransactionLog(
	o orm.Ormer,
	log structDB.TypeGetTxnLogByJobIDAndStatus,
	postHistory structs.TypeResponsePostTxnProfile,
	txnStatus int64,
	returningUpdate *structDB.TypeReturningUpdate,
	transactionTime time.Time,
) (err error) {

	return
}

func (b *DBPost) InsertInTFench(o orm.Ormer, v interface{}) (err error) {
	return nil
}

func (b *DBPost) UpdateTransactionAccount(postingAmount float64, totalCreditMonthToDate float64, accountNumber string,
	txnTime time.Time, o orm.Ormer) (returningUpdateAccount structDB.TypeReturningUpdateAccountWithLedger, err error) {

	return
}
