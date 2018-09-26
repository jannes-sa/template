package earning

import (
	"time"
	"txn/structs"
	structDB "txn/structs/db"

	"errors"
	"strings"
	"txn/helper/constant"

	"github.com/astaxie/beego/orm"
)

func init() {
	var (
		txAccount           structDB.TransactionAccount
		txLog               structDB.TypeTxnDB
		txHistory           structDB.TransactionHistoryDB
		batchEarning        structDB.BatchEarningStatus
		earningBatchHistory structDB.EarningBatchHistory
	)

	// TxAccount  ...
	TxAccount = txAccount.TableName()
	// TxLog ...
	TxLog = txLog.TableName()
	// TxHistory ...
	TxHistory = txHistory.TableName()
	// BatchEarning
	BatchEarning = batchEarning.TableName()
	// EarningBatchHistory
	EarningBatchHistory = earningBatchHistory.TableName()
}

// Earning ...
type Earning struct {
	batchEarningStatus  []structDB.BatchEarningStatus
	earningBatchHistory []structDB.EarningBatchHistory
}

// FindBatchEarningStatus Method to check earning status
func (b *Earning) FindBatchEarningStatus(
	filename string,
	batchDate string,
) (
	res []structDB.BatchEarningStatus,
	err error,
) {
	//beego.Debug("@FindBatchEarningStatus", filename)
	if filename == "LW_EARN_ALREADY_EXECUTE" {
		res = append(res, structDB.BatchEarningStatus{
			BatchDate: batchDate,
			Filename:  filename,
			Status:    0,
		})
	}

	return
}

// FindBatchEarningStatusPending Method to check earning status pending
func (b *Earning) FindBatchEarningStatusPending(batcheDate string) (
	res []structDB.BatchEarningStatus, err error) {

	return
}

// InsertBatchEarning Method to Insert Batch Earning Status
func (b *Earning) InsertBatchEarning(v interface{}) error {

	return nil
}

// RemoveHistory Method to Clear Earning Batch History
func (b *Earning) RemoveHistory(filename string) error {
	b.TruncateTable(EarningBatchHistory)
	return nil
}

// UpdateBatchEarning Method to Update Batch Earning Status ...
func (b *Earning) UpdateBatchEarning(filenaem string, status int, currentDate string, n1 time.Time) error {

	return nil
}

// MigrateTransactionAccount Method to Migrate Transaction Account
// for Scenario Testing
func (b *Earning) MigrateTransactionAccount(errCode *[]structs.TypeError) {

	return
}

// MigrateFromFileTransactionAccount Method To Create Account From File
func (b *Earning) MigrateFromFileTransactionAccount(
	fileInput []string,
	errCode *[]structs.TypeError,
) {

}

// BulkInsert Bulking Insert
func (b *Earning) BulkInsert(theRows []string, condBulk string, filename string) error {
	for i, v := range theRows {
		splitInputFile := strings.Split(v, constant.ReportSeparator)
		b.earningBatchHistory = append(b.earningBatchHistory, structDB.EarningBatchHistory{
			ID:         i,
			TextString: v,
			Status:     2,
			JobID:      splitInputFile[1],
		})
	}
	return nil
}

func bulkFilterEarningBatchHistory(o orm.Ormer, vv []string,
	valueStrings []string, valueArgs []interface{}) error {

	return nil
}

func bulkFilterAccount(o orm.Ormer, vv []string,
	valueStrings []string, valueArgs []interface{}) error {

	return nil
}

// ClearTransactionAccountHistoryAndLog Method To Clear Data
func (b *Earning) ClearTransactionAccountHistoryAndLog(
	errCode *[]structs.TypeError) {

	b.TruncateTable(TxAccount)
	b.TruncateTable(TxLog)
	b.TruncateTable(TxHistory)
	b.TruncateTable(EarningBatchHistory)
}

// ClearBatchEarningStatus Method to Clear Batch Earning Status
func (b *Earning) ClearBatchEarningStatus(errCode *[]structs.TypeError) {
	b.TruncateTable(BatchEarning)
}

// TruncateTable ...
func (b *Earning) TruncateTable(tableName string) {

}

// MigrateToEarningStatus ...
func (b *Earning) MigrateToEarningStatus(batchDate string,
	filename string, status int) {

}

// UpdateEarningHistory ...
func (b *Earning) UpdateEarningHistory(
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
func (b *Earning) GetAllEarningHistoryAndSortIt(filename string) (
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
