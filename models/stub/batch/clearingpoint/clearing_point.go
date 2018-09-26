package clearingpoint

import (
	"txn/helper/constant"
	structDB "txn/structs/db"

	"github.com/astaxie/beego/orm"
)

// ClearingPoint ...
type ClearingPoint struct{}

// FindTransactionAccount ...
func (b *ClearingPoint) FindTransactionAccount(o orm.Ormer, accountStatus int) (
	txnAccount []structDB.TransactionAccount, err error,
) {
	txnAccount = []structDB.TransactionAccount{
		{
			AccountNumber: "50000000013",
			AccountBranch: 1,
			LedgerBalance: 0,
			AccountStatus: accountStatus,
			ProductCode:   constant.LoyaltyProductCode,
		},
		{
			AccountNumber: "50000000014",
			AccountBranch: 1,
			LedgerBalance: 1000000.00,
			AccountStatus: accountStatus,
			ProductCode:   constant.LoyaltyProductCode,
		},
		{
			AccountNumber: "50000000015",
			AccountBranch: 1,
			LedgerBalance: 0,
			AccountStatus: accountStatus,
			ProductCode:   constant.LoyaltyProductCode,
		},
		{
			AccountNumber: "50000000016",
			AccountBranch: 1,
			LedgerBalance: 2000000,
			AccountStatus: accountStatus,
			ProductCode:   constant.LoyaltyProductCode,
		},
		{
			AccountNumber: "50000000017",
			AccountBranch: 1,
			LedgerBalance: 3000000,
			AccountStatus: accountStatus,
			ProductCode:   constant.LoyaltyProductCode,
		},
	}
	return
}

// InsertClearingPointBatchHistory ...
func (b *ClearingPoint) InsertClearingPointBatchHistory(o orm.Ormer, textString string, status int, jobID string) error {

	return nil
}

// SetBalance ...
func (b *ClearingPoint) SetBalance(
	o orm.Ormer,
	balanceSet float64,
	accountNumber string,
	returningUpdateAccount *structDB.TypeReturningUpdateAccount,
) (err error) {

	return nil
}

// RemoveHistory ...
func (b *ClearingPoint) RemoveHistory() error {
	return nil
}

// GetClearingPointHistory ...
func (b *ClearingPoint) GetClearingPointHistory(o orm.Ormer) (
	clearingPointBatchHistory []structDB.ResetPointBatchHistory, err error,
) {
	clearingPointBatchHistory = []structDB.ResetPointBatchHistory{
		{
			TextString: "D||18|50000000013|A|0.00|0.00|0||",
			Status:     0,
			JobID:      "",
		},
		{
			TextString: "D|1808086F704BFA02059659|19|50000000014|A|1000000.00|0.00|0||",
			Status:     0,
			JobID:      "1808086F704BFA02059659",
		},
		{
			TextString: "D||20|50000000015|A|0.00|0.00|0||",
			Status:     0,
			JobID:      "",
		},
		{
			TextString: "D|1808086F704BFA02059660|21|50000000016|A|2000000.00|0.00|0||",
			Status:     0,
			JobID:      "1808086F704BFA02059660",
		},
		{
			TextString: "D|1808086F704BFA02059661|22|50000000017|A|3000000.00|0.00|0||",
			Status:     0,
			JobID:      "1808086F704BFA02059661",
		},
	}
	return
}
