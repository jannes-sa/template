package interdomain

import (
	structsApi "txn/structs/api"
	structsDB "txn/structs/db"

	"github.com/astaxie/beego/orm"
)

// Outgoing ...
type Outgoing struct {
	TXNLog []structsDB.TypeTxnDB
}

// InsertInterDomain ...
func (b *Outgoing) InsertInterDomain(o orm.Ormer, txnDBRows []structsDB.TypeTxnDB, stTx bool) error {
	//errCode := make([]structs.TypeError,0)
	//structs.ErrorCode.BeneficiaryExceedChar.String(&errCode)
	//_,_,_ := structs.ErrorCode.BeneficiaryExceedChar.GetString()

	// for _, v := range b.TXNLog {
	// for _, r := range txnDBRows {
	// if v.ReqID == r.ReqID {
	// return errors.New("error stub")
	// }
	// }
	// }

	b.TXNLog = append(b.TXNLog, txnDBRows...)
	return nil
}

// UpdateTrxLogFail ...
func (b *Outgoing) UpdateTrxLogFail(o orm.Ormer, jobID string,
	allErr string) (err error) {
	for i, v := range b.TXNLog {
		if v.ReqID == jobID {
			b.TXNLog[i].TxnStatus = 3
		}
	}
	return nil
}

// UpdateTrxLogSuccess ...
func (b *Outgoing) UpdateTrxLogSuccess(o orm.Ormer, jobID string, resBody structsApi.ResInterDomainPrePost) (err error) {

	return nil
}
