package file

import (
	"fmt"
)

type BulkAccountLockResponseFileModel struct {
	CustomerId         string `json:"customerId,omitempty"`         // length 10
	AccountId          string `json:"accountId,omitempty"`          // length 10
	LockTypeCode       string `json:"lockTypeCode,omitempty"`       // length 3
	LockReasonTypeCode string `json:"lockReasonTypeCode,omitempty"` // length 3
	LockFailReason     string `json:"lockFailReason,omitempty"`     // length 255

	Typed_CustomerId int64 `json:"-"`
	Typed_AccountId  int64 `json:"-"`
}

type TypedBulkAccountLockResponseFileModel struct {
	CustomerId         int64  `json:"customerId,omitempty"`         // length 10
	AccountId          int64  `json:"accountId,omitempty"`          // length 10
	LockTypeCode       string `json:"lockTypeCode,omitempty"`       // length 3
	LockReasonTypeCode string `json:"lockReasonTypeCode,omitempty"` // length 3
	LockFailReason     string `json:"lockFailReason,omitempty"`     // length 255
}

func (mdl *BulkAccountLockResponseFileModel) IsTabDelimitedFile() bool {
	return false
}

func (mdl *BulkAccountLockResponseFileModel) HeaderFieldCount() int {
	return 9
}

func (mdl *BulkAccountLockResponseFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedBulkAccountLockResponseFileModel{
			CustomerId:         mdl.Typed_CustomerId,
			AccountId:          mdl.Typed_AccountId,
			LockTypeCode:       mdl.LockTypeCode,
			LockReasonTypeCode: mdl.LockReasonTypeCode,
			LockFailReason:     mdl.LockFailReason,
		}
	}
}

func (mdl *BulkAccountLockResponseFileModel) GetValues() []string {
	return []string{
		mdl.CustomerId,
		mdl.AccountId,
		mdl.LockTypeCode,
		mdl.LockReasonTypeCode,
		mdl.LockFailReason,
	}
}

func (mdl *BulkAccountLockResponseFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	pos, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.AccountId, &mdl.Typed_AccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.LockTypeCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.LockReasonTypeCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.LockFailReason, line, pos, 255, preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
