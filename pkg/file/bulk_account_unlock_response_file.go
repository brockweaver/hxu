package file

import (
	"fmt"
)

type BulkAccountUnlockResponseFileModel struct {
	CustomerId       string `json:"customerId,omitempty"`       // length 10
	AccountId        string `json:"accountId,omitempty"`        // length 10
	UnlockResultCode string `json:"unlockResultCode,omitempty"` // length 3
	UnlockFailReason string `json:"unlockFailReason,omitempty"` // length 256

	Typed_CustomerId int64 `json:"-"`
	Typed_AccountId  int64 `json:"-"`
}

type TypedBulkAccountUnlockResponseFileModel struct {
	CustomerId       int64  `json:"customerId,omitempty"`       // length 10
	AccountId        int64  `json:"accountId,omitempty"`        // length 10
	UnlockResultCode string `json:"unlockResultCode,omitempty"` // length 3
	UnlockFailReason string `json:"unlockFailReason,omitempty"` // length 256
}

func (mdl *BulkAccountUnlockResponseFileModel) IsTabDelimitedFile() bool {
	return false
}

func (mdl *BulkAccountUnlockResponseFileModel) HeaderFieldCount() int {
	return 9
}

func (mdl *BulkAccountUnlockResponseFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedBulkAccountUnlockResponseFileModel{
			CustomerId:       mdl.Typed_CustomerId,
			AccountId:        mdl.Typed_AccountId,
			UnlockResultCode: mdl.UnlockResultCode,
			UnlockFailReason: mdl.UnlockFailReason,
		}
	}
}

func (mdl *BulkAccountUnlockResponseFileModel) GetValues() []string {
	return []string{
		mdl.CustomerId,
		mdl.AccountId,
		mdl.UnlockResultCode,
		mdl.UnlockFailReason,
	}
}

func (mdl *BulkAccountUnlockResponseFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	pos, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.AccountId, &mdl.Typed_AccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.UnlockResultCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.UnlockFailReason, line, pos, 255, preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
