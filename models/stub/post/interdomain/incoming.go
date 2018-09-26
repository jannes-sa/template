package interdomain

import (
	"time"
	"txn/helper/constant"
	"txn/structs"
	structsPg "txn/structs/db"

	"github.com/astaxie/beego"

	"errors"
	"github.com/astaxie/beego/orm"
)

// Outgoing ...
type Incoming struct{}

// FindTxnLogByRequestID ...
func (b *Incoming) FindTxnLogByRequestID(
	o orm.Ormer,
	jobID string,
	status int,
) (
	res []structsPg.TypeGetTxnLogByJobIDAndStatus,
	err error,
) {
	txnTime := "2017-08-24T13:43:56.41906615Z"
	srcTime, err := time.Parse(time.RFC3339Nano, txnTime)
	if err != nil {
		beego.Warning("err parse time")
	}
	if jobID == "1150" {
		res = append(res, structsPg.TypeGetTxnLogByJobIDAndStatus{
			JobID:             "1150",
			TransactionTime:   srcTime,
			AccountNumber:     "100000000001",
			AccountName:       "Hello Kitty",
			BankCode:          constant.SMITHBANKCODE,
			TransactionStatus: 2,
			TransactionAmount: 1000,
			JournalNumber:     1,
			CifNumber:         1001,
			Currency:          "IDR",
		})
	}

	return
}

// InsertInterDomainIndicator ...
func (b *Incoming) InsertInterDomainIndicator(o orm.Ormer, reqID string, headerAll string,
	reqBody string, url string, reason string) (err error) {
	if reqID == "1150" {
		err = nil
	} else {
		err = errors.New("reqBody not found")
	}
	return
}

// UpdateTransactionAccount ...
func (b *Incoming) UpdateTransactionAccount(
	o orm.Ormer,
	log structsPg.TypeGetTxnLogByJobIDAndStatus,
	txnStatus int64,
	returningUpdateAccount *structsPg.TypeReturningUpdateAccount,
) (err error) {
	err = nil
	return
}

// UpdateTransactionLog ...
func (b *Incoming) UpdateTransactionLog(
	o orm.Ormer,
	log structsPg.TypeGetTxnLogByJobIDAndStatus,
	postHistory structs.TypeResponsePostTxnProfile,
	txnStatus int64,
	returningUpdate *structsPg.TypeReturningUpdate,
	transactionTime time.Time,
	relAccName string,
) (err error) {
	err = nil
	return
}

// InsertTxnHistory ...
func (b *Incoming) InsertTxnHistory(
	o orm.Ormer, hist structsPg.TransactionHistoryDB,
) (err error) {
	err = nil
	return
}
