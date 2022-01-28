package file

import (
	"fmt"
)

type AchTransactionFileModel struct {
	CustomerId               string `json:"customerId,omitempty"`               // length 10
	CustomerTag              string `json:"customerTag,omitempty"`              // length 50
	AccountId                string `json:"accountId,omitempty"`                // length 10
	AccountTag               string `json:"accountTag,omitempty"`               // length 50
	AccountName              string `json:"accountName,omitempty"`              // length 50
	TransactionId            string `json:"transactionId,omitempty"`            // length 19
	TransactionTag           string `json:"transactionTag,omitempty"`           // length 50
	TransactionTypeCode      string `json:"transactionTypeCode,omitempty"`      // length 6
	TraceNumber              string `json:"traceNumber,omitempty"`              // length 15
	StandardEntryClassCode   string `json:"standardEntryClassCode,omitempty"`   // length 3
	CompanyName              string `json:"companyName,omitempty"`              // length 16
	CompanyDiscretionaryData string `json:"companyDiscretionaryData,omitempty"` // length 20
	CompanyEntryDescription  string `json:"companyEntryDescription,omitempty"`  // length 10
	ReceivingCompanyName     string `json:"receivingCompanyName,omitempty"`     // length 22
	IdentificationNumber     string `json:"identificationNumber,omitempty"`     // length 15

	Typed_CustomerId    int64 `json:"-"`
	Typed_AccountId     int64 `json:"-"`
	Typed_TransactionId int64 `json:"-"`
}

type TypedAchTransactionFileModel struct {
	CustomerId               int64  `json:"customerId,omitempty"`               // length 10
	CustomerTag              string `json:"customerTag,omitempty"`              // length 50
	AccountId                int64  `json:"accountId,omitempty"`                // length 10
	AccountTag               string `json:"accountTag,omitempty"`               // length 50
	AccountName              string `json:"accountName,omitempty"`              // length 50
	TransactionId            int64  `json:"transactionId,omitempty"`            // length 19
	TransactionTag           string `json:"transactionTag,omitempty"`           // length 50
	TransactionTypeCode      string `json:"transactionTypeCode,omitempty"`      // length 6
	TraceNumber              string `json:"traceNumber,omitempty"`              // length 15
	StandardEntryClassCode   string `json:"standardEntryClassCode,omitempty"`   // length 3
	CompanyName              string `json:"companyName,omitempty"`              // length 16
	CompanyDiscretionaryData string `json:"companyDiscretionaryData,omitempty"` // length 20
	CompanyEntryDescription  string `json:"companyEntryDescription,omitempty"`  // length 10
	ReceivingCompanyName     string `json:"receivingCompanyName,omitempty"`     // length 22
	IdentificationNumber     string `json:"identificationNumber,omitempty"`     // length 15
}

func (mdl *AchTransactionFileModel) IsTabDelimitedFile() bool {
	return false
}

func (mdl *AchTransactionFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *AchTransactionFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return &TypedAchTransactionFileModel{
			CustomerId:               mdl.Typed_CustomerId,
			CustomerTag:              mdl.CustomerTag,
			AccountId:                mdl.Typed_AccountId,
			AccountTag:               mdl.AccountTag,
			AccountName:              mdl.AccountName,
			TransactionId:            mdl.Typed_TransactionId,
			TransactionTag:           mdl.TransactionTag,
			TransactionTypeCode:      mdl.TransactionTypeCode,
			TraceNumber:              mdl.TraceNumber,
			StandardEntryClassCode:   mdl.StandardEntryClassCode,
			CompanyName:              mdl.CompanyName,
			CompanyDiscretionaryData: mdl.CompanyDiscretionaryData,
			CompanyEntryDescription:  mdl.CompanyEntryDescription,
			ReceivingCompanyName:     mdl.ReceivingCompanyName,
			IdentificationNumber:     mdl.IdentificationNumber,
		}
	}
}

func (mdl *AchTransactionFileModel) GetValues() []string {
	return []string{
		mdl.CustomerId,
		mdl.CustomerTag,
		mdl.AccountId,
		mdl.AccountTag,
		mdl.AccountName,
		mdl.TransactionId,
		mdl.TransactionTag,
		mdl.TransactionTypeCode,
		mdl.TraceNumber,
		mdl.StandardEntryClassCode,
		mdl.CompanyName,
		mdl.CompanyDiscretionaryData,
		mdl.CompanyEntryDescription,
		mdl.ReceivingCompanyName,
		mdl.IdentificationNumber,
	}
}

func (mdl *AchTransactionFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	pos, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.CustomerTag, line, pos, 50, preserve, err)
	pos, err = ConsumeIntField(&mdl.AccountId, &mdl.Typed_AccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.AccountTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.AccountName, line, pos, 50, preserve, err)
	pos, err = ConsumeIntField(&mdl.TransactionId, &mdl.Typed_TransactionId, line, pos, 19, preserve, err)
	pos, err = ConsumeField(&mdl.TransactionTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.TransactionTypeCode, line, pos, 6, preserve, err)
	pos, err = ConsumeField(&mdl.TraceNumber, line, pos, 15, preserve, err)
	pos, err = ConsumeField(&mdl.StandardEntryClassCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.CompanyName, line, pos, 16, preserve, err)
	pos, err = ConsumeField(&mdl.CompanyDiscretionaryData, line, pos, 20, preserve, err)
	pos, err = ConsumeField(&mdl.CompanyEntryDescription, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.ReceivingCompanyName, line, pos, 22, preserve, err)
	pos, err = ConsumeField(&mdl.IdentificationNumber, line, pos, 15, preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
