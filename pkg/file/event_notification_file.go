package file

import (
	"fmt"
)

type EventNotificationFileModel struct {
	CustomerId               string `json:"customerId,omitempty"`               // length 10
	CustomerTag              string `json:"customerTag,omitempty"`              // length 50
	AccountId                string `json:"accountId,omitempty"`                // length 10
	AccountTag               string `json:"accountTag,omitempty"`               // length 50
	ExternalAccountId        string `json:"externalAccountid,omitempty"`        // length 10
	ExternalAccountTag       string `json:"externalAccountTag,omitempty"`       // length 50
	TransactionId            string `json:"transactionId,omitempty"`            // length 19
	TransactionTag           string `json:"transactionTag,omitempty"`           // length 50
	Description              string `json:"description,omitempty"`              // length 50
	EventTypeId              string `json:"eventTypeId,omitempty"`              // length 10
	TransactionTypeCode      string `json:"transactionTypeCode,omitempty"`      // length 6
	TransactionSettledDate   string `json:"transactionSettledDate,omitempty"`   // length 34
	TransactionAvailableDate string `json:"transactionAvailableDate,omitempty"` // length 34
	UserEventId              string `json:"userEventId,omitempty"`              // length 19
	MasterId                 string `json:"masterId,omitempty"`                 // length 19
	TransactionAmount        string `json:"transactionAmount,omitempty"`        // length 10
	TransactionCreatedDate   string `json:"transactionCreatedDate,omitempty"`   // length 34
	TransactionSubTypeCode   string `json:"transactionSubTypeCode,omitempty"`   // length 6
	TransactionTypeCode2     string `json:"transactionTypeCode2,omitempty"`     // length 6
	FromAccountId            string `json:"fromAccountId,omitempty"`            // length 10
	FromAvailableBalance     string `json:"fromAvailableBalance,omitempty"`     // length 15
	FromAccountBalance       string `json:"fromAccountBalance,omitempty"`       // length 15
	FromPendingBalance       string `json:"fromPendingBalance,omitempty"`       // length 15
	ToAccountId              string `json:"toAccountId,omitempty"`              // length 10
	ToAvailableBalance       string `json:"toAvailableBalance,omitempty"`       // length 15
	ToAccountBalance         string `json:"toAccountBalance,omitempty"`         // length 15
	ToPendingBalance         string `json:"toPendingBalance,omitempty"`         // length 15
	ModifiedById             string `json:"modifiedById,omitempty"`             // length 10
	EventDate                string `json:"eventDate,omitempty"`                // length 34
	IsAdminUser              string `json:"isAdminUser,omitempty"`              // length 1

	Typed_CustomerId            int64   `json:"-"`
	Typed_AccountId             int64   `json:"-"`
	Typed_ExternalAccountId     int64   `json:"-"`
	Typed_TransactionId         int64   `json:"-"`
	Typed_EventTypeId           int64   `json:"-"`
	Typed_UserEventId           int64   `json:"-"`
	Typed_MasterId              int64   `json:"-"`
	Typed_TransactionAmount     float64 `json:"-"`
	Typed_ReturnedTransactionId int64   `json:"-"`
	Typed_FromAccountId         int64   `json:"-"`
	Typed_FromAvailableBalance  float64 `json:"-"`
	Typed_FromAccountBalance    float64 `json:"-"`
	Typed_FromPendingBalance    float64 `json:"-"`
	Typed_ToAccountId           int64   `json:"-"`
	Typed_ToAvailableBalance    float64 `json:"-"`
	Typed_ToAccountBalance      float64 `json:"-"`
	Typed_ToPendingBalance      float64 `json:"-"`
	Typed_ModifiedById          int64   `json:"-"`
	Typed_IsAdminUser           bool    `json:"-"`
}

type TypedEventNotificationFileModel struct {
	CustomerId               int64   `json:"customerId,omitempty"`               // length 10
	CustomerTag              string  `json:"customerTag,omitempty"`              // length 50
	AccountId                int64   `json:"accountId,omitempty"`                // length 10
	AccountTag               string  `json:"accountTag,omitempty"`               // length 50
	ExternalAccountId        int64   `json:"externalAccountid,omitempty"`        // length 10
	ExternalAccountTag       string  `json:"externalAccountTag,omitempty"`       // length 50
	TransactionId            int64   `json:"transactionId,omitempty"`            // length 19
	TransactionTag           string  `json:"transactionTag,omitempty"`           // length 50
	Description              string  `json:"description,omitempty"`              // length 50
	EventTypeId              int64   `json:"eventTypeId,omitempty"`              // length 10
	TransactionTypeCode      string  `json:"transactionTypeCode,omitempty"`      // length 6
	TransactionSettledDate   string  `json:"transactionSettledDate,omitempty"`   // length 34
	TransactionAvailableDate string  `json:"transactionAvailableDate,omitempty"` // length 34
	UserEventId              int64   `json:"userEventId,omitempty"`              // length 19
	MasterId                 int64   `json:"masterId,omitempty"`                 // length 19
	TransactionAmount        float64 `json:"transactionAmount,omitempty"`        // length 10
	TransactionCreatedDate   string  `json:"transactionCreatedDate,omitempty"`   // length 34
	TransactionSubTypeCode   string  `json:"transactionSubTypeCode,omitempty"`   // length 6
	TransactionTypeCode2     string  `json:"transactionTypeCode2,omitempty"`     // length 6
	FromAccountId            int64   `json:"fromAccountId,omitempty"`            // length 10
	FromAvailableBalance     float64 `json:"fromAvailableBalance,omitempty"`     // length 15
	FromAccountBalance       float64 `json:"fromAccountBalance,omitempty"`       // length 15
	FromPendingBalance       float64 `json:"fromPendingBalance,omitempty"`       // length 15
	ToAccountId              int64   `json:"toAccountId,omitempty"`              // length 10
	ToAvailableBalance       float64 `json:"toAvailableBalance,omitempty"`       // length 15
	ToAccountBalance         float64 `json:"toAccountBalance,omitempty"`         // length 15
	ToPendingBalance         float64 `json:"toPendingBalance,omitempty"`         // length 15
	ModifiedById             int64   `json:"modifiedById,omitempty"`             // length 10
	EventDate                string  `json:"eventDate,omitempty"`                // length 34
	IsAdminUser              bool    `json:"isAdminUser,omitempty"`              // length 1
}

func (mdl *EventNotificationFileModel) IsTabDelimitedFile() bool {
	return true
}

func (mdl *EventNotificationFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *EventNotificationFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedEventNotificationFileModel{
			CustomerId:               mdl.Typed_CustomerId,
			CustomerTag:              mdl.CustomerTag,
			AccountId:                mdl.Typed_AccountId,
			AccountTag:               mdl.AccountTag,
			ExternalAccountId:        mdl.Typed_ExternalAccountId,
			ExternalAccountTag:       mdl.ExternalAccountTag,
			TransactionId:            mdl.Typed_TransactionId,
			TransactionTag:           mdl.TransactionTag,
			Description:              mdl.Description,
			EventTypeId:              mdl.Typed_EventTypeId,
			TransactionTypeCode:      mdl.TransactionTypeCode,
			TransactionSettledDate:   mdl.TransactionSettledDate,
			TransactionAvailableDate: mdl.TransactionAvailableDate,
			UserEventId:              mdl.Typed_UserEventId,
			MasterId:                 mdl.Typed_MasterId,
			TransactionAmount:        mdl.Typed_TransactionAmount,
			TransactionCreatedDate:   mdl.TransactionCreatedDate,
			TransactionSubTypeCode:   mdl.TransactionSubTypeCode,
			TransactionTypeCode2:     mdl.TransactionTypeCode2,
			FromAccountId:            mdl.Typed_FromAccountId,
			FromAvailableBalance:     mdl.Typed_FromAvailableBalance,
			FromAccountBalance:       mdl.Typed_FromAccountBalance,
			FromPendingBalance:       mdl.Typed_FromPendingBalance,
			ToAccountId:              mdl.Typed_ToAccountId,
			ToAvailableBalance:       mdl.Typed_ToAvailableBalance,
			ToAccountBalance:         mdl.Typed_ToAccountBalance,
			ToPendingBalance:         mdl.Typed_ToPendingBalance,
			ModifiedById:             mdl.Typed_ModifiedById,
			EventDate:                mdl.EventDate,
			IsAdminUser:              mdl.Typed_IsAdminUser,
		}
	}
}

func (mdl *EventNotificationFileModel) GetValues() []string {
	return []string{
		mdl.CustomerId,
		mdl.CustomerTag,
		mdl.AccountId,
		mdl.AccountTag,
		mdl.ExternalAccountId,
		mdl.ExternalAccountTag,
		mdl.TransactionId,
		mdl.TransactionTag,
		mdl.Description,
		mdl.EventTypeId,
		mdl.TransactionTypeCode,
		mdl.TransactionSettledDate,
		mdl.TransactionAvailableDate,
		mdl.UserEventId,
		mdl.MasterId,
		mdl.TransactionAmount,
		mdl.TransactionCreatedDate,
		mdl.TransactionSubTypeCode,
		mdl.TransactionTypeCode2,
		mdl.FromAccountId,
		mdl.FromAvailableBalance,
		mdl.FromAccountBalance,
		mdl.FromPendingBalance,
		mdl.ToAccountId,
		mdl.ToAvailableBalance,
		mdl.ToAccountBalance,
		mdl.ToPendingBalance,
		mdl.ModifiedById,
		mdl.EventDate,
		mdl.IsAdminUser,
	}
}

func (mdl *EventNotificationFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	pos, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.CustomerTag, line, pos, 50, preserve, err)
	pos, err = ConsumeIntField(&mdl.AccountId, &mdl.Typed_AccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.AccountTag, line, pos, 50, preserve, err)
	pos, err = ConsumeIntField(&mdl.ExternalAccountId, &mdl.Typed_ExternalAccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.ExternalAccountTag, line, pos, 50, preserve, err)
	pos, err = ConsumeIntField(&mdl.TransactionId, &mdl.Typed_TransactionId, line, pos, 19, preserve, err)
	pos, err = ConsumeField(&mdl.TransactionTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.Description, line, pos, 50, preserve, err)
	pos, err = ConsumeIntField(&mdl.EventTypeId, &mdl.Typed_EventTypeId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.TransactionTypeCode, line, pos, 6, preserve, err)
	pos, err = ConsumeField(&mdl.TransactionSettledDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.TransactionAvailableDate, line, pos, 34, preserve, err)
	pos, err = ConsumeIntField(&mdl.UserEventId, &mdl.Typed_UserEventId, line, pos, 19, preserve, err)
	pos, err = ConsumeIntField(&mdl.MasterId, &mdl.Typed_MasterId, line, pos, 19, preserve, err)
	pos, err = ConsumeFloatField(&mdl.TransactionAmount, &mdl.Typed_TransactionAmount, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.TransactionCreatedDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.TransactionSubTypeCode, line, pos, 6, preserve, err)
	pos, err = ConsumeField(&mdl.TransactionTypeCode2, line, pos, 6, preserve, err)
	pos, err = ConsumeIntField(&mdl.FromAccountId, &mdl.Typed_FromAccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeFloatField(&mdl.FromAvailableBalance, &mdl.Typed_FromAvailableBalance, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.FromAccountBalance, &mdl.Typed_FromAccountBalance, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.FromPendingBalance, &mdl.Typed_FromPendingBalance, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeIntField(&mdl.ToAccountId, &mdl.Typed_ToAccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeFloatField(&mdl.ToAvailableBalance, &mdl.Typed_ToAvailableBalance, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.ToAccountBalance, &mdl.Typed_ToAccountBalance, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.ToPendingBalance, &mdl.Typed_ToPendingBalance, line, pos, 15, 2, 2, preserve, err)

	pos, err = ConsumeIntField(&mdl.ModifiedById, &mdl.Typed_ModifiedById, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.EventDate, line, pos, 34, preserve, err)
	pos, err = ConsumeBoolField(&mdl.IsAdminUser, &mdl.Typed_IsAdminUser, line, pos, 1, "1", preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
