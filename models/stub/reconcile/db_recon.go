package reconcile

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"txn/structs"
	structPg "txn/structs/db"
)

//DBRecon - stub implementation of IDBRecon
type DBRecon struct {
	TxnLogs    []structPg.TypeTxnDB
	TxnAccount []structPg.TransactionAccount
	TxnHistory []structPg.TransactionHistoryDB
}

func (b *DBRecon) setUpData() {
	b.TxnLogs = []structPg.TypeTxnDB{
		{
			RelatedBankCode: "506",
			ToAccountBranch: 2,
			DrCr:            "Dr",
			TxnVal:          25.0,
			Cur:             "THB",
			AccountType:     "GL",
			TxnTime:         time.Date(2018, time.Month(5), 17, 0, 0, 0, 0, &time.Location{}),
		},
		{
			RelatedBankCode:   "506",
			FromAccountBranch: 1,
			DrCr:              "Dr",
			TxnVal:            25.0,
			Cur:               "THB",
			AccountType:       "GL",
			TxnTime:           time.Date(2018, time.Month(5), 18, 0, 0, 0, 1, &time.Location{}),
		},
		{
			RelatedBankCode: "506",
			DrCr:            "Dr",
			TxnVal:          25.0,
			Cur:             "THB",
			AccountType:     "GL",
			TxnTime:         time.Date(2018, time.Month(5), 19, 0, 0, 0, 2, &time.Location{}),
		},
	}
	b.TxnAccount = []structPg.TransactionAccount{
		{
			AccountNumber: "50000000018",
			CustomerType:  "1010",
			AccountBranch: 1,
			Currency:      "THB",
			ProductCode:   "SA003",
			AccountStatus: 1,
		},
	}
	b.TxnHistory = []structPg.TransactionHistoryDB{
		{
			AccountNumber: "50000000018",
			CurrentLedger: 100.0,
			TxnTime:       time.Date(2018, time.Month(5), 16, 0, 0, 0, 0, &time.Location{}),
		},
	}
}

//GetReportInterClearing - ...
func (b *DBRecon) GetReportInterClearing(interClearingDate string) ([]structs.InterclearingData, error) {
	b.setUpData()
	inputDate := strings.Split(interClearingDate, "-")
	var resInterClearing []structs.InterclearingData
	for _, v := range b.TxnLogs {
		y, m, d := v.TxnTime.Date()
		if strconv.Itoa(y) == inputDate[0] && fmt.Sprintf("%02d", int(m)) == inputDate[1] && fmt.Sprintf("%02d", d) == inputDate[2] {
			resInterClearing = append(resInterClearing, structs.InterclearingData{
				BankCode:               v.RelatedBankCode,
				DrCr:                   v.DrCr,
				TransactionAmount:      fmt.Sprintf("%.2f", v.TxnVal),
				ToAccountBranch:        v.ToAccountBranch,
				FromAccountBranch:      v.FromAccountBranch,
				Currency:               v.Cur,
				TransactionAccountType: v.AccountType,
			})
		}
	}

	return resInterClearing, nil

}

//GetReportOutstanding - ...
func (b *DBRecon) GetReportOutstanding(outstandingDate string) ([]structs.OutstandingData, error) {
	b.setUpData()
	inputTime, _ := time.Parse("2006-01-02", outstandingDate)
	var outstandingData []structs.OutstandingData
	for _, hist := range b.TxnHistory {
		for _, ac := range b.TxnAccount {
			if hist.AccountNumber == ac.AccountNumber && (hist.TxnTime.Before(inputTime) || hist.TxnTime.Equal(inputTime)) {
				outstandingData = append(outstandingData, structs.OutstandingData{
					AccountNumber:   ac.AccountNumber,
					AccountStatus:   ac.AccountStatus,
					ProductCode:     ac.ProductCode,
					Currency:        ac.Currency,
					AccountBranch:   int64(ac.AccountBranch),
					CifCustomerType: ac.CustomerType,
					LedgerBalance:   hist.CurrentLedger,
				})
			}
		}
	}

	return outstandingData, nil
}
