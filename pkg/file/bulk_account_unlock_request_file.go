package file

import (
	"fmt"
)

type BulkAccountUnlockRequestFileModel struct {
	CustomerId string `json:"customerId,omitempty"` // length 10
	AccountId  string `json:"accountId,omitempty"`  // length 10
	Notes      string `json:"notes,omitempty"`      // length 256

	Typed_CustomerId int64 `json:"-"`
	Typed_AccountId  int64 `json:"-"`
}

type TypedBulkAccountUnlockRequestFileModel struct {
	CustomerId int64  `json:"customerId,omitempty"` // length 10
	AccountId  int64  `json:"accountId,omitempty"`  // length 10
	Notes      string `json:"notes,omitempty"`      // length 256
}

func (mdl *BulkAccountUnlockRequestFileModel) IsTabDelimitedFile() bool {
	return false
}

func (mdl *BulkAccountUnlockRequestFileModel) HeaderFieldCount() int {
	return 6
}

func (mdl *BulkAccountUnlockRequestFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedBulkAccountUnlockRequestFileModel{
			CustomerId: mdl.Typed_CustomerId,
			AccountId:  mdl.Typed_AccountId,
			Notes:      mdl.Notes,
		}
	}
}

func (mdl *BulkAccountUnlockRequestFileModel) GetValues() []string {
	return []string{
		mdl.CustomerId,
		mdl.AccountId,
		mdl.Notes,
	}
}

func (mdl *BulkAccountUnlockRequestFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	pos, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.AccountId, &mdl.Typed_AccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.Notes, line, pos, 256, preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
