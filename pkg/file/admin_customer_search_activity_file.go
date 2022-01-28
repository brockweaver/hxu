package file

import (
	"fmt"
	"strings"
)

type AdminCustomerSearchActivityFileModel struct {
	UserId                 string `json:"userId,omitempty"`                 // length 10
	FirstName              string `json:"firstName,omitempty"`              // length 255
	LastName               string `json:"lastName,omitempty"`               // length 255
	Tag                    string `json:"tag,omitempty"`                    // length 50
	AccountNumber          string `json:"accountNumber,omitempty"`          // length 50
	EmailAddress           string `json:"emailAddress,omitempty"`           // length 255
	MobilePhone            string `json:"mobilePhone,omitempty"`            // length 50
	TaxId                  string `json:"taxId,omitempty"`                  // length 30
	CustomerId             string `json:"customerId,omitempty"`             // length 10
	AccountTag             string `json:"accountTag,omitempty"`             // length 50
	ExternalAccountTag     string `json:"externalAccountTag,omitempty"`     // length 50
	TransactionTag         string `json:"transactionTag,omitempty"`         // length 50
	ReceiptReferenceNumber string `json:"receiptReferenceNumber,omitempty"` // length 19
	Date                   string `json:"date,omitempty"`                   // length 34
	ProgramId              string `json:"programId,omitempty"`              // length 10
	ProgramName            string `json:"programName,omitempty"`            // length 50

	Typed_UserId                 int64 `json:"-"`
	Typed_CustomerId             int64 `json:"-"`
	Typed_ReceiptReferenceNumber int64 `json:"-"`
	Typed_ProgramId              int64 `json:"-"`
}

type TypedAdminCustomerSearchActivityFileModel struct {
	UserId                 int64  `json:"userId,omitempty"`                 // length 10
	FirstName              string `json:"firstName,omitempty"`              // length 255
	LastName               string `json:"lastName,omitempty"`               // length 255
	Tag                    string `json:"tag,omitempty"`                    // length 50
	AccountNumber          string `json:"accountNumber,omitempty"`          // length 50
	EmailAddress           string `json:"emailAddress,omitempty"`           // length 255
	MobilePhone            string `json:"mobilePhone,omitempty"`            // length 50
	TaxId                  string `json:"taxId,omitempty"`                  // length 30
	CustomerId             int64  `json:"customerId,omitempty"`             // length 10
	AccountTag             string `json:"accountTag,omitempty"`             // length 50
	ExternalAccountTag     string `json:"externalAccountTag,omitempty"`     // length 50
	TransactionTag         string `json:"transactionTag,omitempty"`         // length 50
	ReceiptReferenceNumber int64  `json:"receiptReferenceNumber,omitempty"` // length 19
	Date                   string `json:"date,omitempty"`                   // length 34
	ProgramId              int64  `json:"programId,omitempty"`              // length 10
	ProgramName            string `json:"programName,omitempty"`            // length 50
}

func (mdl *AdminCustomerSearchActivityFileModel) IsTabDelimitedFile() bool {
	return true
}

func (mdl *AdminCustomerSearchActivityFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *AdminCustomerSearchActivityFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return &TypedAdminCustomerSearchActivityFileModel{
			UserId:                 mdl.Typed_UserId,
			FirstName:              mdl.FirstName,
			LastName:               mdl.LastName,
			Tag:                    mdl.Tag,
			AccountNumber:          mdl.AccountNumber,
			EmailAddress:           mdl.EmailAddress,
			MobilePhone:            mdl.MobilePhone,
			TaxId:                  mdl.TaxId,
			CustomerId:             mdl.Typed_CustomerId,
			AccountTag:             mdl.AccountTag,
			ExternalAccountTag:     mdl.ExternalAccountTag,
			TransactionTag:         mdl.TransactionTag,
			ReceiptReferenceNumber: mdl.Typed_ReceiptReferenceNumber,
			Date:                   mdl.Date,
			ProgramId:              mdl.Typed_ProgramId,
			ProgramName:            mdl.ProgramName,
		}
	}
}

func (mdl *AdminCustomerSearchActivityFileModel) GetValues() []string {
	return []string{
		mdl.UserId,
		mdl.FirstName,
		mdl.LastName,
		mdl.Tag,
		mdl.AccountNumber,
		mdl.EmailAddress,
		mdl.MobilePhone,
		mdl.TaxId,
		mdl.CustomerId,
		mdl.AccountTag,
		mdl.ExternalAccountTag,
		mdl.TransactionTag,
		mdl.ReceiptReferenceNumber,
		mdl.Date,
		mdl.ProgramId,
		mdl.ProgramName,
	}
}

func (mdl *AdminCustomerSearchActivityFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	// special: this is a tab-delimited file
	splits := strings.Split(line, "\t")

	const expectedFieldCount = 16
	actualFieldCount := len(splits)

	if strict && actualFieldCount != expectedFieldCount {
		err = fmt.Errorf("expected exactly %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else if !strict && actualFieldCount < expectedFieldCount {
		err = fmt.Errorf("expected at least %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else {

		_, err = ConsumeIntField(&mdl.UserId, &mdl.Typed_UserId, splits[0], pos, len(splits[0]), preserve, err)
		_, err = ConsumeField(&mdl.FirstName, splits[1], pos, len(splits[1]), preserve, err)
		_, err = ConsumeField(&mdl.LastName, splits[2], pos, len(splits[2]), preserve, err)
		_, err = ConsumeField(&mdl.Tag, splits[3], pos, len(splits[3]), preserve, err)
		_, err = ConsumeField(&mdl.AccountNumber, splits[4], pos, len(splits[4]), preserve, err)
		_, err = ConsumeField(&mdl.EmailAddress, splits[5], pos, len(splits[5]), preserve, err)
		_, err = ConsumeField(&mdl.MobilePhone, splits[6], pos, len(splits[6]), preserve, err)
		_, err = ConsumeField(&mdl.TaxId, splits[7], pos, len(splits[7]), preserve, err)
		_, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, splits[8], pos, len(splits[8]), preserve, err)

		_, err = ConsumeField(&mdl.AccountTag, splits[9], pos, len(splits[9]), preserve, err)
		_, err = ConsumeField(&mdl.ExternalAccountTag, splits[10], pos, len(splits[10]), preserve, err)
		_, err = ConsumeField(&mdl.TransactionTag, splits[11], pos, len(splits[11]), preserve, err)
		_, err = ConsumeIntField(&mdl.ReceiptReferenceNumber, &mdl.Typed_ReceiptReferenceNumber, splits[12], pos, len(splits[12]), preserve, err)
		_, err = ConsumeField(&mdl.Date, splits[13], pos, len(splits[13]), preserve, err)

		_, err = ConsumeIntField(&mdl.ProgramId, &mdl.Typed_ProgramId, splits[14], pos, len(splits[14]), preserve, err)
		_, err = ConsumeField(&mdl.ProgramName, splits[15], pos, len(splits[15]), preserve, err)

	}

	return err
}
