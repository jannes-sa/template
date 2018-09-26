package loyalty

import (
	"time"
	"txn/helper/timetn"
	"txn/structs"
	structDB "txn/structs/db"

	"fmt"
	"strconv"
	"strings"
	"txn/helper/constant"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// DBLoyalty ...
type DBLoyalty struct {
	txnHistRows []structDB.TransactionHistoryDB
	txnLogRows  []structDB.TypeTxnDB
}

func (b *DBLoyalty) setUpData() {
	b.txnHistRows = []structDB.TransactionHistoryDB{
		{
			ChannelPostDate:  time.Date(2001, time.Month(1), 1, 1, 1, 1, 1, &time.Location{}),
			TxnTime:          time.Date(2001, time.Month(1), 1, 1, 1, 1, 1, &time.Location{}),
			JobID:            "jobid1",
			AccountNumber:    "500000000012",
			TxnAmount:        20.0,
			TransactionType:  "earning",
			TxnCode:          constant.EarnPointEventCode,
			AdjustmentAction: "11",
			Comment:          "comment field in earning",
			Currency:         "RWP",
			DrCr:             "Cr",
			CifNumber:        3,
		},
		{
			ChannelPostDate:      time.Date(2001, time.Month(1), 1, 1, 1, 1, 1, &time.Location{}),
			TxnTime:              time.Date(2001, time.Month(1), 1, 1, 1, 1, 1, &time.Location{}),
			JobID:                "jobid2",
			AccountNumber:        "500000000013",
			TxnAmount:            30.0,
			TransactionType:      "redemption",
			TxnCode:              constant.RedeemPointEventCode,
			OriginalReverseJobId: "jobid3",
			AdjustmentAction:     "22",
			Currency:             "RWP",
			DrCr:                 "Cr",
			CifNumber:            3,
			ValueOfRedemtion:     66,
		},
		{
			ChannelPostDate:  time.Date(2001, time.Month(1), 1, 1, 1, 1, 1, &time.Location{}),
			TxnTime:          time.Date(2001, time.Month(1), 1, 1, 1, 1, 1, &time.Location{}),
			JobID:            "jobid3",
			AccountNumber:    "500000000013",
			TxnAmount:        30.0,
			TransactionType:  "redemption",
			TxnCode:          constant.RedeemPointEventCode,
			Comment:          "Test Original Reverse Comment",
			Currency:         "RWP",
			DrCr:             "Dr",
			CifNumber:        3,
			ValueOfRedemtion: 66,
		},
	}

	b.txnLogRows = []structDB.TypeTxnDB{
		{
			TxnTime: timetn.Now(),
			ReqID:   "123",
			Comment: "comment",
			DrCr:    "Dr",
		},
		{
			TxnTime: timetn.Now(),
			ReqID:   "123",
			Comment: "comment",
			DrCr:    "Cr",
		},
		{
			TxnTime: timetn.Now(),
			ReqID:   "jobid3",
			Comment: "YFZF-348P-5LZG-P1YC~ABC-XYZ-123-456",
			DrCr:    "Cr",
		},
	}
}

// GetAccount ...
func (b *DBLoyalty) GetAccount(
	accountNumber string,
	o orm.Ormer,
	errCode *[]structs.TypeError,
) structDB.TransactionAccount {

	if accountNumber == "11000000000995" {
		var accs = structDB.TransactionAccount{
			AccountNumber: accountNumber,
			AccountStatus: 0,
			// AccruedInterestUnposted:  0,
			AccountBranch: 10,
			CifNumber:     101019920,
			CreationDate:  timetn.Now(),
			Currency:      "IDR",
			CustomerType:  "SA001",
			Index:         11,
			//InterestID:               "IN001",
			//InterestPosting:          3,
			//InterestPostingFrequency: "D",
			LedgerBalance:  1000010,
			MaximumBalance: 0,
			ProductCode:    "PN001",
		}

		return accs
	}

	var ProductCode = "PN001"
	if accountNumber != "11000000000005" {
		ProductCode = "PNKKI"
	}

	accForEarning := []string{
		"10000000004001",
		"10000000004002",
		"10000000004003",
		"11000000000009",
		"11000000000010",
		"11000000000011",
	}
	for _, v := range accForEarning {
		if v == accountNumber {
			ProductCode = "PN001"
		}
	}

	var accs = structDB.TransactionAccount{
		AccountNumber: accountNumber,
		AccountStatus: 1,
		// AccruedInterestUnposted:  0,
		AccountBranch: 10,
		CifNumber:     101019920,
		CreationDate:  timetn.Now(),
		Currency:      "IDR",
		CustomerType:  "SA001",
		Index:         11,
		//InterestID:               "IN001",
		//InterestPosting:          3,
		//InterestPostingFrequency: "D",
		LedgerBalance:  1000010,
		MaximumBalance: 0,
		ProductCode:    ProductCode,
	}

	return accs
}

// UpdateTransactionAccount ...
func (b *DBLoyalty) UpdateTransactionAccount(
	postingAmount float64,
	accountNumber string,
	stateTx *bool,
	o orm.Ormer,
	returningUpdateAccount *structDB.TypeReturningUpdateAccount,
	txnTime time.Time,
	isEarning bool,
	isRedemption bool,
) {
	(*stateTx) = true

	if accountNumber == "1234" {
		(*returningUpdateAccount) = structDB.TypeReturningUpdateAccount{
			LedgerBalance:           0,
			LedgerBalanceLastUpdate: timetn.Now(),
			AccountNumber:           accountNumber,
			TotalCreditMonthToDate:  postingAmount,
		}
		return
	}

	(*returningUpdateAccount) = structDB.TypeReturningUpdateAccount{
		LedgerBalance:           postingAmount,
		LedgerBalanceLastUpdate: timetn.Now(),
		AccountNumber:           accountNumber,
		TotalCreditMonthToDate:  postingAmount,
	}
}

func (b *DBLoyalty) buildParamAccount(
	postingAmount float64,
	accountNumber string,
	txnTime time.Time,
	isEarning bool,
	isRedemption bool,
) []interface{} {
	param := make([]interface{}, 0)

	return param
}

// UpdateTransactionLogs ...
func (b *DBLoyalty) UpdateTransactionLogs(
	logs []structDB.TypeTxnDB,
	stateTx *bool,
	o orm.Ormer,
	errCode *[]structs.TypeError,
) {
	if logs[0].TxnVal == 0 {
		structs.ErrorCode.UnexpectedError.String(errCode)
		(*stateTx) = false
		return
	}
	(*stateTx) = true
}

// buildTransactionLogItemStmt ...
func (b *DBLoyalty) buildTransactionLogItemStmt(
	transactionAmount float64,
	logJobID string,
	logAccountNumber string,
	TableName string,
	logDrCr string,
) (
	param []interface{},
	sql string,
) {

	return
}

// GetLoyaltyHistoryNonReverse ...
func (b *DBLoyalty) GetLoyaltyHistoryNonReverse(
	channelPostDate string,
) (
	txnHistRows []structDB.TransactionHistoryDB,
	err error,
) {
	beego.Warning("test", channelPostDate)
	inputDateYMD := strings.Split(channelPostDate, "-")
	d := DBLoyalty{}
	d.setUpData()
	for _, v := range d.txnHistRows {
		y, m, d := v.ChannelPostDate.Date()
		if strconv.Itoa(y) == inputDateYMD[0] && fmt.Sprintf("%02d", int(m)) == inputDateYMD[1] && fmt.Sprintf("%02d", d) == inputDateYMD[2] {
			txnHistRows = append(txnHistRows, v)
		}
	}

	return
}

func (b *DBLoyalty) GetTxnLogByJobId(jobId string) (result []structDB.TypeTxnDB, err error) {
	b.setUpData()
	for _, v := range b.txnLogRows {
		if v.ReqID == jobId {
			result = append(result, v)
		}
	}
	return
}

// GetLoyaltyHistoryByDate ...
func (b *DBLoyalty) GetLoyaltyHistoryByDate(
	channelPostDate string,
) (
	txnHistRows []structDB.TransactionHistoryDB,
	err error,
) {
	beego.Warning("test", channelPostDate)
	inputDateYMD := strings.Split(channelPostDate, "-")
	d := DBLoyalty{}
	d.setUpData()
	for _, v := range d.txnHistRows {
		y, m, d := v.ChannelPostDate.Date()
		if strconv.Itoa(y) == inputDateYMD[0] && fmt.Sprintf("%02d", int(m)) == inputDateYMD[1] && fmt.Sprintf("%02d", d) == inputDateYMD[2] {
			txnHistRows = append(txnHistRows, v)
		}
	}

	return
}

// InsertAccountingFailedMessage Used for insert log if MQ fail
func (b *DBLoyalty) InsertAccountingFailedMessage(logMQ structDB.AccountingFailedMessage) error {

	return nil
}
