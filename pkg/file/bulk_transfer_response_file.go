package file

import (
	"fmt"
)

type BulkTransferResponseFileModel struct {
	CustomerId       string `json:"customerId,omitempty"`       // length 10
	CustomerTag      string `json:"customerTag,omitempty"`      // length 50
	TransferTag      string `json:"transferTag,omitempty"`      // length 50
	TransferKind     string `json:"transferKind,omitempty"`     // length 3
	TransferAmount   string `json:"transferAmount,omitempty"`   // length 10
	ToAccountId      string `json:"toAccountId,omitempty"`      // length 10
	FromAccountId    string `json:"fromAccountId,omitempty"`    // length 10
	ToAccountTag     string `json:"toAccountTag,omitempty"`     // length 50
	FromAccountTag   string `json:"fromAccountTag,omitempty"`   // length 50
	NachaDescription string `json:"nachaDescription,omitempty"` // length 255
	ErrorNumber      string `json:"errorNumber,omitempty"`      // length 10
	ErrorMessage     string `json:"errorMessage,omitempty"`     // length 255

	Typed_CustomerId     int64   `json:"-"`
	Typed_TransferAmount float64 `json:"-"`
	Typed_ToAccountId    int64   `json:"-"`
	Typed_FromAccountId  int64   `json:"-"`
	Typed_ErrorNumber    int64   `json:"-"`
}

type TypedBulkTransferResponseFileModel struct {
	CustomerId       int64   `json:"customerId,omitempty"`       // length 10
	CustomerTag      string  `json:"customerTag,omitempty"`      // length 50
	TransferTag      string  `json:"transferTag,omitempty"`      // length 50
	TransferKind     string  `json:"transferKind,omitempty"`     // length 3
	TransferAmount   float64 `json:"transferAmount,omitempty"`   // length 10
	ToAccountId      int64   `json:"toAccountId,omitempty"`      // length 10
	FromAccountId    int64   `json:"fromAccountId,omitempty"`    // length 10
	ToAccountTag     string  `json:"toAccountTag,omitempty"`     // length 50
	FromAccountTag   string  `json:"fromAccountTag,omitempty"`   // length 50
	NachaDescription string  `json:"nachaDescription,omitempty"` // length 255
	ErrorNumber      int64   `json:"errorNumber,omitempty"`      // length 10
	ErrorMessage     string  `json:"errorMessage,omitempty"`     // length 255
}

func (mdl *BulkTransferResponseFileModel) IsTabDelimitedFile() bool {
	return false
}

func (mdl *BulkTransferResponseFileModel) HeaderFieldCount() int {
	return 9
}

func (mdl *BulkTransferResponseFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedBulkTransferResponseFileModel{
			CustomerId:       mdl.Typed_CustomerId,
			CustomerTag:      mdl.CustomerTag,
			TransferTag:      mdl.TransferTag,
			TransferKind:     mdl.TransferKind,
			TransferAmount:   mdl.Typed_TransferAmount,
			ToAccountId:      mdl.Typed_ToAccountId,
			FromAccountId:    mdl.Typed_FromAccountId,
			ToAccountTag:     mdl.ToAccountTag,
			FromAccountTag:   mdl.FromAccountTag,
			NachaDescription: mdl.NachaDescription,
			ErrorNumber:      mdl.Typed_ErrorNumber,
			ErrorMessage:     mdl.ErrorMessage,
		}
	}
}

func (mdl *BulkTransferResponseFileModel) GetValues() []string {
	return []string{
		mdl.CustomerId,
		mdl.CustomerTag,
		mdl.TransferTag,
		mdl.TransferKind,
		mdl.TransferAmount,
		mdl.ToAccountId,
		mdl.FromAccountId,
		mdl.ToAccountTag,
		mdl.FromAccountTag,
		mdl.NachaDescription,
		mdl.ErrorNumber,
		mdl.ErrorMessage,
	}
}

func (mdl *BulkTransferResponseFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	pos, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.CustomerTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.TransferTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.TransferKind, line, pos, 3, preserve, err)
	pos, err = ConsumeFloatField(&mdl.TransferAmount, &mdl.Typed_TransferAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeIntField(&mdl.ToAccountId, &mdl.Typed_ToAccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.FromAccountId, &mdl.Typed_FromAccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.ToAccountTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.FromAccountTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.NachaDescription, line, pos, 255, preserve, err)
	pos, err = ConsumeIntField(&mdl.ErrorNumber, &mdl.Typed_ErrorNumber, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.ErrorMessage, line, pos, 255, preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
