package constant

const (
	TypeDebit          = "withdrawal"
	ChanName           = "MBL"
	TypeCredit         = "deposit"
	Dr                 = "Dr"
	Cr                 = "Cr"
	LoyaltyProductCode = "PN001"
	TypeIncoming       = "incoming"
	TypeOutgoing       = "outgoing"
	AC                 = "AC"

	StatusTxnSuccess               = 1
	StatusTxnPending               = 2
	StatusTxnFailPrepostValidation = 3
	StatusTxnFailPostValidation    = 4
	StatusTxnFailTXNDomain         = 5
	StatusTxnTrxFence              = 6

	ExceptionLastInsertID = "no LastInsertId available"
)
