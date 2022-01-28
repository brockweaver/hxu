package file

import (
	"fmt"
)

type PostedTransactionFileModel struct {
	CustomerId             string `json:"customerId,omitempty"`             // length 10
	CustomerTag            string `json:"customerTag,omitempty"`            // length 50
	AccountId              string `json:"accountId,omitempty"`              // length 10
	AccountTag             string `json:"accountTag,omitempty"`             // length 50
	AccountName            string `json:"accountName,omitempty"`            // length 50
	TransactionId          string `json:"transactionId,omitempty"`          // length 19
	TransactionTag         string `json:"transactionTag,omitempty"`         // length 50
	TransactionTypeCode    string `json:"transactionTypeCode,omitempty"`    // length 6
	TransactionAmount      string `json:"transactionAmount,omitempty"`      // length 10
	Action                 string `json:"action,omitempty"`                 // length 1
	TransactionDescription string `json:"transactionDescription,omitempty"` // length 255
	NachaDescription       string `json:"nachaDescription,omitempty"`       // length 255
	CreatedDate            string `json:"createdDate,omitempty"`            // length 34
	SettledDate            string `json:"settledDate,omitempty"`            // length 34
	AvailableDate          string `json:"availableDate,omitempty"`          // length 34
	MasterId               string `json:"masterId,omitempty"`               // length 19
	ReturnCode             string `json:"returnCode,omitempty"`             // length 3
	FeeCode                string `json:"feeCode,omitempty"`                // length 3
	ExternalAccountId      string `json:"externalAccountid,omitempty"`      // length 10
	ReturnedTransactionId  string `json:"returnedTransactionId,omitempty"`  // length 19
	DebitAccountId         string `json:"debitAccountId,omitempty"`         // length 10
	CreditAccountId        string `json:"creditAccountId,omitempty"`        // length 10
	ProductId              string `json:"productId,omitempty"`              // length 10

	Typed_CustomerId            int64   `json:"-"`
	Typed_AccountId             int64   `json:"-"`
	Typed_TransactionId         int64   `json:"-"`
	Typed_TransactionAmount     float64 `json:"-"`
	Typed_MasterId              int64   `json:"-"`
	Typed_ExternalAccountId     int64   `json:"-"`
	Typed_ReturnedTransactionId int64   `json:"-"`
	Typed_DebitAccountId        int64   `json:"-"`
	Typed_CreditAccountId       int64   `json:"-"`
	Typed_ProductId             int64   `json:"-"`
}

type TypedPostedTransactionFileModel struct {
	CustomerId             int64   `json:"customerId,omitempty"`             // length 10
	CustomerTag            string  `json:"customerTag,omitempty"`            // length 50
	AccountId              int64   `json:"accountId,omitempty"`              // length 10
	AccountTag             string  `json:"accountTag,omitempty"`             // length 50
	AccountName            string  `json:"accountName,omitempty"`            // length 50
	TransactionId          int64   `json:"transactionId,omitempty"`          // length 19
	TransactionTag         string  `json:"transactionTag,omitempty"`         // length 50
	TransactionTypeCode    string  `json:"transactionTypeCode,omitempty"`    // length 6
	TransactionAmount      float64 `json:"transactionAmount,omitempty"`      // length 10
	Action                 string  `json:"action,omitempty"`                 // length 1
	TransactionDescription string  `json:"transactionDescription,omitempty"` // length 255
	NachaDescription       string  `json:"nachaDescription,omitempty"`       // length 255
	CreatedDate            string  `json:"createdDate,omitempty"`            // length 34
	SettledDate            string  `json:"settledDate,omitempty"`            // length 34
	AvailableDate          string  `json:"availableDate,omitempty"`          // length 34
	MasterId               int64   `json:"masterId,omitempty"`               // length 19
	ReturnCode             string  `json:"returnCode,omitempty"`             // length 3
	FeeCode                string  `json:"feeCode,omitempty"`                // length 3
	ExternalAccountId      string  `json:"externalAccountid,omitempty"`      // length 10
	ReturnedTransactionId  string  `json:"returnedTransactionId,omitempty"`  // length 19
	DebitAccountId         int64   `json:"debitAccountId,omitempty"`         // length 10
	CreditAccountId        int64   `json:"creditAccountId,omitempty"`        // length 10
	ProductId              int64   `json:"productId,omitempty"`              // length 10
}

func (mdl *PostedTransactionFileModel) IsTabDelimitedFile() bool {
	return false
}

func (mdl *PostedTransactionFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *PostedTransactionFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedPostedTransactionFileModel{
			CustomerId:             mdl.Typed_CustomerId,
			CustomerTag:            mdl.CustomerTag,
			AccountId:              mdl.Typed_AccountId,
			AccountTag:             mdl.AccountTag,
			AccountName:            mdl.AccountName,
			TransactionId:          mdl.Typed_TransactionId,
			TransactionTag:         mdl.TransactionTag,
			TransactionTypeCode:    mdl.TransactionTypeCode,
			TransactionAmount:      mdl.Typed_TransactionAmount,
			Action:                 mdl.Action,
			TransactionDescription: mdl.TransactionDescription,
			NachaDescription:       mdl.NachaDescription,
			CreatedDate:            mdl.CreatedDate,
			SettledDate:            mdl.SettledDate,
			AvailableDate:          mdl.AvailableDate,
			MasterId:               mdl.Typed_MasterId,
			ReturnCode:             mdl.ReturnCode,
			FeeCode:                mdl.FeeCode,
			ExternalAccountId:      mdl.ExternalAccountId,
			ReturnedTransactionId:  mdl.ReturnedTransactionId,
			DebitAccountId:         mdl.Typed_DebitAccountId,
			CreditAccountId:        mdl.Typed_CreditAccountId,
			ProductId:              mdl.Typed_ProductId,
		}
	}
}

func (mdl *PostedTransactionFileModel) GetValues() []string {
	return []string{
		mdl.CustomerId,
		mdl.CustomerTag,
		mdl.AccountId,
		mdl.AccountTag,
		mdl.AccountName,
		mdl.TransactionId,
		mdl.TransactionTypeCode,
		mdl.TransactionAmount,
		mdl.Action,
		mdl.TransactionDescription,
		mdl.NachaDescription,
		mdl.CreatedDate,
		mdl.SettledDate,
		mdl.AvailableDate,
		mdl.MasterId,
		mdl.ReturnCode,
		mdl.FeeCode,
		mdl.ExternalAccountId,
		mdl.ReturnedTransactionId,
		mdl.DebitAccountId,
		mdl.CreditAccountId,
		mdl.ProductId,
	}
}

func (mdl *PostedTransactionFileModel) ParseNative(line string, strict bool, preserve bool) error {
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
	pos, err = ConsumeFloatField(&mdl.TransactionAmount, &mdl.Typed_TransactionAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.Action, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.TransactionDescription, line, pos, 255, preserve, err)
	pos, err = ConsumeField(&mdl.NachaDescription, line, pos, 255, preserve, err)
	pos, err = ConsumeField(&mdl.CreatedDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.SettledDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.AvailableDate, line, pos, 34, preserve, err)
	pos, err = ConsumeIntField(&mdl.MasterId, &mdl.Typed_MasterId, line, pos, 19, preserve, err)
	pos, err = ConsumeField(&mdl.ReturnCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.FeeCode, line, pos, 3, preserve, err)
	pos, err = ConsumeIntField(&mdl.ExternalAccountId, &mdl.Typed_ExternalAccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.ReturnedTransactionId, &mdl.Typed_ReturnedTransactionId, line, pos, 19, preserve, err)
	pos, err = ConsumeIntField(&mdl.DebitAccountId, &mdl.Typed_DebitAccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.CreditAccountId, &mdl.Typed_CreditAccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.ProductId, &mdl.Typed_ProductId, line, pos, 10, preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
