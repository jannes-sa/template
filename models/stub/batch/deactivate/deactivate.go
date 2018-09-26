package deactivate

import (
	"errors"
	"strings"
	"time"
	"txn/helper/constant"
	structDB "txn/structs/db"

	"github.com/astaxie/beego/orm"
)

// Deactivate ...
type Deactivate struct{}

// FindBatchInfo Method to find batch info
func (b *Deactivate) FindBatchInfo(
	fullFilename string,
	batchDate string,
) (
	res []structDB.BatchEarningStatus,
	err error,
) {

	return
}

// FindAccountByAccountNumber ...
func (b *Deactivate) FindAccountByAccountNumber(o orm.Ormer, accountNumber string) (
	txnAccount structDB.TransactionAccount, err error,
) {
	if accountNumber == "1234AlreadyClosed" {
		txnAccount = structDB.TransactionAccount{
			AccountNumber: accountNumber,
			AccountBranch: 1,
			LedgerBalance: 1000000,
			AccountStatus: constant.AccountStatusClosed,
			ProductCode:   constant.LoyaltyProductCode,
		}
		return
	}
	txnAccount = structDB.TransactionAccount{
		AccountNumber: accountNumber,
		AccountBranch: 1,
		LedgerBalance: 1000000,
		AccountStatus: constant.AccountStatusActive,
		ProductCode:   constant.LoyaltyProductCode,
	}
	return
}

// CloseAccount Method to close account status
func (b *Deactivate) CloseAccount(o orm.Ormer, accountNumber string,
	closedDatetime time.Time) (err error) {
	return nil
}

// RemoveHistory Method to Clear Earning Batch History based on Filename
func (b *Deactivate) RemoveHistory(filename string) error {

	return nil
}

// BulkInsert Bulking Insert
func (b *Deactivate) BulkInsert(theRows []string, condBulk string, filename string) error {

	return nil
}

// UpdateEarningHistory ...
func (b *Deactivate) UpdateEarningHistory(
	o orm.Ormer,
	status int,
	val string,
	jobID string,
) error {
	var e error
	if status == 1 {
		e = errors.New("|1|err_code|err_desc")
	} else {
		e = errors.New("|0")
	}
	return e
}

// GetAllEarningHistoryAndSortIt ...
func (b *Deactivate) GetAllEarningHistoryAndSortIt(filename string) (
	res []structDB.EarningBatchHistory,
	err error) {
	//res = b.earningBatchHistory
	responseErr := "|0||"
	if strings.Contains(filename, "Fail") {
		responseErr = "|1|error_code|err_desc"
	}
	if strings.Contains(filename, "LS_ADJ_EARN") {
		res = append(res,
			structDB.EarningBatchHistory{
				JobID:      "123",
				Filename:   "",
				ID:         1,
				Status:     2,
				TextString: "D|180209C8A0D47006659884|180209C8A0D47061538193|MyMoGSB|1031|11000000000009|2400|SFE02" + responseErr,
			},
			structDB.EarningBatchHistory{
				JobID:      "456",
				Filename:   "",
				ID:         2,
				Status:     2,
				TextString: "D|180209C8A0D47009281037|180209C8A0D47001736292|MyMoGSB|1348|11000000000010|2400|SFE02" + responseErr,
			},
			structDB.EarningBatchHistory{
				JobID:      "789",
				Filename:   "",
				ID:         3,
				Status:     2,
				TextString: "D|180209C8A0D47091730173|180209C8A0D47345678283|MyMoGSB|1053|11000000000011|2400|SFE02" + responseErr,
			},
		)
	} else if strings.Contains(filename, "LS_DEACTIVATE") {
		res = append(res,
			structDB.EarningBatchHistory{
				JobID:      "123",
				Filename:   "",
				ID:         1,
				Status:     2,
				TextString: "D|180209C8A0D47006659456|1031|11000000000009|C|2.00|0.00" + responseErr,
			},
			structDB.EarningBatchHistory{
				JobID:      "456",
				Filename:   "",
				ID:         2,
				Status:     2,
				TextString: "D|180209C8A0D47002222222|1348|11000000000010|C|3.00|0.00" + responseErr,
			},
			structDB.EarningBatchHistory{
				JobID:      "789",
				Filename:   "",
				ID:         3,
				Status:     2,
				TextString: "D|180209C8A0D47093333333|1053|11000000000011|C|4.00|0.00" + responseErr,
			},
		)
	} else {
		res = append(res,
			structDB.EarningBatchHistory{
				JobID:      "123",
				Filename:   "",
				ID:         1,
				Status:     2,
				TextString: "D|180209C8A0D47006659884||MyMoGSB|1031|11000000000009|2400|SFE01" + responseErr,
			},
			structDB.EarningBatchHistory{
				JobID:      "456",
				Filename:   "",
				ID:         2,
				Status:     2,
				TextString: "D|180209C8A0D47009281037||MyMoGSB|1348|11000000000010|2400|SFE01" + responseErr,
			},
			structDB.EarningBatchHistory{
				JobID:      "789",
				Filename:   "",
				ID:         3,
				Status:     2,
				TextString: "D|180209C8A0D47091730173||MyMoGSB|1053|11000000000011|2400|SFE01" + responseErr,
			},
		)
	}

	err = nil

	return
}
