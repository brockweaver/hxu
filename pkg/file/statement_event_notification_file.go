package file

import (
	"fmt"
)

type StatementEventNotificationFileModel struct {
	UserEventId      string `json:"userEventId,omitempty"`      // length 19
	CustomerId       string `json:"customerId,omitempty"`       // length 10
	AccountId        string `json:"accountId,omitempty"`        // length 10
	Month            string `json:"month,omitempty"`            // length 2
	Year             string `json:"year,omitempty"`             // length 4
	NumberOfAccounts string `json:"numberOfAccounts,omitempty"` // length 4
	EventTypeId      string `json:"eventTypeId,omitempty"`      // length 10
	EventDate        string `json:"eventDate,omitempty"`        // length 34

	Typed_UserEventId      int64 `json:"-"`
	Typed_CustomerId       int64 `json:"-"`
	Typed_AccountId        int64 `json:"-"`
	Typed_Month            int64 `json:"-"`
	Typed_Year             int64 `json:"-"`
	Typed_NumberOfAccounts int64 `json:"-"`
	Typed_EventTypeId      int64 `json:"-"`
}

type TypedStatementEventNotificationFileModel struct {
	UserEventId      int64  `json:"userEventId,omitempty"`      // length 19
	CustomerId       int64  `json:"customerId,omitempty"`       // length 10
	AccountId        int64  `json:"accountId,omitempty"`        // length 10
	Month            int64  `json:"month,omitempty"`            // length 2
	Year             int64  `json:"year,omitempty"`             // length 4
	NumberOfAccounts int64  `json:"numberOfAccounts,omitempty"` // length 4
	EventTypeId      int64  `json:"eventTypeId,omitempty"`      // length 10
	EventDate        string `json:"eventDate,omitempty"`        // length 34
}

func (mdl *StatementEventNotificationFileModel) IsTabDelimitedFile() bool {
	return false
}

func (mdl *StatementEventNotificationFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *StatementEventNotificationFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedStatementEventNotificationFileModel{
			UserEventId:      mdl.Typed_UserEventId,
			CustomerId:       mdl.Typed_CustomerId,
			AccountId:        mdl.Typed_AccountId,
			Month:            mdl.Typed_Month,
			Year:             mdl.Typed_Year,
			NumberOfAccounts: mdl.Typed_NumberOfAccounts,
			EventTypeId:      mdl.Typed_EventTypeId,
			EventDate:        mdl.EventDate,
		}
	}
}

func (mdl *StatementEventNotificationFileModel) GetValues() []string {
	return []string{
		mdl.UserEventId,
		mdl.CustomerId,
		mdl.AccountId,
		mdl.Month,
		mdl.Year,
		mdl.NumberOfAccounts,
		mdl.EventTypeId,
		mdl.EventDate,
	}
}

func (mdl *StatementEventNotificationFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	pos, err = ConsumeIntField(&mdl.UserEventId, &mdl.Typed_UserEventId, line, pos, 19, preserve, err)
	pos, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.AccountId, &mdl.Typed_AccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.Month, &mdl.Typed_Month, line, pos, 2, preserve, err)
	pos, err = ConsumeIntField(&mdl.Year, &mdl.Typed_Year, line, pos, 4, preserve, err)
	pos, err = ConsumeIntField(&mdl.NumberOfAccounts, &mdl.Typed_NumberOfAccounts, line, pos, 4, preserve, err)
	pos, err = ConsumeIntField(&mdl.EventTypeId, &mdl.Typed_EventTypeId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.EventDate, line, pos, 34, preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
