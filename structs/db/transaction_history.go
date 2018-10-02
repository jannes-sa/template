package models

import "time"

type TransactionHistory struct {
	AccountNumber          string    `orm:"column(account_number)"`
	SequenceRef            int       `orm:"column(sequence_ref)"`
	TransactionTime        time.Time `orm:"column(transaction_time);type(timestamp with time zone)"`
	SystemReferenceId      string    `orm:"column(system_reference_id)"`
	JournalNumber          int       `orm:"column(journal_number)"`
	CifNumber              int64     `orm:"column(cif_number)"`
	EventCode              string    `orm:"column(event_code)"`
	TransactionDescription string    `orm:"column(transaction_description)"`
	DrCr                   string    `orm:"column(dr_cr)"`
	BankCode               string    `orm:"column(bank_code)"`
	ProductCode            string    `orm:"column(product_code)"`
	IndexNumber            int64     `orm:"column(index_number)"`
	AccountName            string    `orm:"column(account_name)"`
	TransactionAmount      float64   `orm:"column(transaction_amount)"`
	TransactionStatus      int       `orm:"column(transaction_status)"`
	TotalCreditMonthToDate float64   `orm:"column(total_credit_month_to_date)"`
	Remarks                string    `orm:"column(remarks);null"`
	Currency               string    `orm:"column(currency)"`
	AccountingType         string    `orm:"column(accounting_type)"`
	TransactionAccountType string    `orm:"column(transaction_account_type)"`
	AccountingTranCode     string    `orm:"column(accounting_tran_code)"`
	RelatedAccount         string    `orm:"column(related_account)"`
	RelatedAccountName     string    `orm:"column(related_account_name)"`
	TransactionType        string    `orm:"column(transaction_type)"`
	RelatedBankCode        string    `orm:"column(related_bank_code);null"`
	JobId                  string    `orm:"column(job_id);null"`
	LedgerBalance          float64   `orm:"column(ledger_balance)"`
	CreateDate             time.Time `orm:"column(create_date);type(timestamp with time zone);null"`
	UpdateDate             time.Time `orm:"column(update_date);type(timestamp with time zone);null"`
	Comment                string    `orm:"column(comment);null"`
	ActualTime             time.Time `orm:"column(actual_time);type(timestamp with time zone);null"`
	ChannelPostDate        time.Time `orm:"column(channel_post_date);type(timestamp with time zone);null"`
	FromAccountBranch      int       `orm:"column(from_account_branch);null"`
	ServiceBranch          int       `orm:"column(service_branch);null"`
	ToAccountBranch        int       `orm:"column(to_account_branch);null"`
	OriginationDomain      string    `orm:"column(origination_domain);null"`
	DestinationDomain      string    `orm:"column(destination_domain);null"`
	AdjustmentAction       string    `orm:"column(adjustment_action);null"`
	ValueOfRedemption      float64   `orm:"column(value_of_redemption);null"`
	OriginalReverseJobId   string    `orm:"column(original_reverse_job_id)"`
	IsReverse              bool      `orm:"column(is_reverse);null"`
}
