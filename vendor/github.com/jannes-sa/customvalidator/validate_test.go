package customvalidator

import (
	"encoding/json"
	"log"
	"testing"
)

func TestValidate(t *testing.T) {
	type TesInterface struct {
		BeneficiaryAccountNumber interface{} `json:"beneficiary_account_number" type:"string" convert:"removezero=E02000401" validate:"required=E02030009,stringnumericonly=E02030008,lte=20=E02030007,type=E02030001,identicField=="`
		BeneficiaryBankCode      interface{} `json:"beneficiary_bank_code" type:"string" validate:"stringnumericonly=E02030018,lte=3=E02030017,required=E02030019"`
		SourceAccountNumber      interface{} `json:"source_account_number" type:"string" convert:"removezero=E02000401" validate:"stringnumericonly=E02030013,lte=20=E02030014,required=E02030015,identicField=BeneficiaryAccountNumber=E02030016"`
		SourceBankCode           interface{} `json:"source_bank_code" type:"string" validate:"stringnumericonly=E02030021,lte=3=E02030020,required=E02030022"`
		TransactionAmount        interface{} `json:"transaction_amount" type:"float64" validate:",gte=0=E0209898984,lte=999999999999999=E02030010,type=E02030011,required=E02030012"`
	}

	type Tes struct {
		BeneficiaryAccountNumber string  `json:"beneficiary_account_number"`
		BeneficiaryBankCode      string  `json:"beneficiary_bank_code"`
		SourceAccountNumber      string  `json:"source_account_number"`
		SourceBankCode           string  `json:"source_bank_code"`
		TransactionAmount        float64 `json:"transaction_amount"`
	}

	bodyJSON := []byte(`
		{
			"beneficiary_account_number":"123456789123",
			"beneficiary_bank_code":"009",
			"source_account_number":"97837641248321",
			"source_bank_code":"021",
			"transaction_amount":100000
		}
	`)

	var tesInter TesInterface
	err := json.Unmarshal(bodyJSON, &tesInter)
	if err != nil {
		t.Error("Error ", err)
		return
	}
	log.Println(tesInter)

	var tes Tes
	codeError := Validate(tesInter, &tes)

	log.Println(tes, codeError)
}

func TestValidateArray(t *testing.T) {
	type ArrTesInterface struct {
		ID            interface{} `json:"id" type:"string" validate:"required=E02030009"`
		AccountNumber interface{} `json:"account_number" type:"string" validate:"required=E02030009,stringnumericonly=E02030008,lte=20=E02030007"`
		Status        interface{} `json:"status" type:"int" validate:"required=E02030009"`
		Amount        interface{} `json:"amount" type:"float64" validate:"required=E02030009"`
	}

	type TesInterface2 struct {
		TransactionAmount interface{}       `json:"transaction_amount" type:"float64" validate:"gte=0=E0209898984,lte=999999999999999=E02030010,required=E02030012"`
		Data              []ArrTesInterface `json:"data"`
	}

	type ArrTes struct {
		ID            string  `json:"id"`
		AccountNumber string  `json:"account_number"`
		Status        int     `json:"status"`
		Amount        float64 `json:"amount"`
	}

	type Tes2 struct {
		TransactionAmount float64
		Data              []ArrTes
	}

	bodyJSON := []byte(`
		{
			"transaction_amount":100000,
			"data":[
				{
					"id":"a",
					"account_number":"123456789123",
					"status":1,
					"amount":2000
				},
				{
					"id":"b",
					"account_number":"985654637462",
					"status":2,
					"amount":3000
				},
				{
					"id":"c",
					"account_number":"101010637462",
					"status":3,
					"amount":1500	
				}
			]
		}
	`)

	var tesInter TesInterface2
	err := json.Unmarshal(bodyJSON, &tesInter)
	if err != nil {
		t.Error("Error ", err)
		return
	}

	var tes Tes2
	errCode := Validate(tesInter, &tes)

	if len(errCode) > 0 {
		t.Error("validate Fail ", errCode)
		log.Println(errCode, tes)
		return
	}

	log.Println(errCode, tes)

}

func TestTimeValid(t *testing.T) {
	type TesInterface struct {
		DateString interface{} `json:"date" type:"string" validate:"required=E02030009,timevalid=E02031109"`
	}

	type Tes struct {
		DateString string
	}

	bodyJSON := []byte(`
		{
			"date":"1992-01-12"
		}
	`)

	var tesInter TesInterface
	err := json.Unmarshal(bodyJSON, &tesInter)
	if err != nil {
		t.Error("Error ", err)
		return
	}
	log.Println(tesInter)

	var tes Tes
	codeError := Validate(tesInter, &tes)

	log.Println(tes, codeError)
}
