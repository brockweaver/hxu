package file

import (
	"fmt"
)

type AccountBalanceFileModel struct {
	CustomerId               string `json:"customerId,omitempty"`               // length 10
	CustomerTag              string `json:"customerTag,omitempty"`              // length 50
	AccountId                string `json:"accountId,omitempty"`                // length 10
	AccountTag               string `json:"accountTag,omitempty"`               // length 50
	AccountName              string `json:"accountName,omitempty"`              // length 50
	AccountNumber            string `json:"accountNumber,omitempty"`            // length 50
	AccountType              string `json:"accountType,omitempty"`              // length 50
	AccountStatus            string `json:"accountStatus,omitempty"`            // length 50
	AccountBalance           string `json:"accountBalance,omitempty"`           // length 15
	CreatedDate              string `json:"createdDate,omitempty"`              // length 34
	ClosedDate               string `json:"closedDate,omitempty"`               // length 34
	TargetDate               string `json:"targetDate,omitempty"`               // length 8
	TargetAmount             string `json:"targetAmount,omitempty"`             // length 15
	Category                 string `json:"category,omitempty"`                 // length 50
	Subcategory              string `json:"subcategory,omitempty"`              // length 50
	TargetMetDate            string `json:"targetMetDate,omitempty"`            // length 34
	TargetMetPercent         string `json:"targetMetPercent,omitempty"`         // length 15
	IsPrimary                string `json:"isPrimary,omitempty"`                // length 1
	PrimaryCustomerId        string `json:"primaryCustomerId,omitempty"`        // length 10
	InterestRate             string `json:"interestRate,omitempty"`             // length 15
	ProductId                string `json:"productId,omitempty"`                // length 10
	AvailableBalance         string `json:"availableBalance,omitempty"`         // length 15
	PendingBalance           string `json:"pendingBalance,omitempty"`           // length 15
	AccountLockCode          string `json:"accountLockCode,omitempty"`          // length 3
	AccountLockEffectiveDate string `json:"accountLockEffectiveDate,omitempty"` // length 34

	Typed_CustomerId        int64   `json:"-"`
	Typed_AccountId         int64   `json:"-"`
	Typed_AccountBalance    float64 `json:"-"`
	Typed_TargetAmount      float64 `json:"-"`
	Typed_TargetMetPercent  float64 `json:"-"`
	Typed_IsPrimary         bool    `json:"-"`
	Typed_PrimaryCustomerId int64   `json:"-"`
	Typed_InterestRate      float64 `json:"-"`
	Typed_ProductId         int64   `json:"-"`
	Typed_AvailableBalance  float64 `json:"-"`
	Typed_PendingBalance    float64 `json:"-"`
}

type TypedAccountBalanceFileModel struct {
	CustomerId               int64   `json:"customerId,omitempty"`               // length 10
	CustomerTag              string  `json:"customerTag,omitempty"`              // length 50
	AccountId                int64   `json:"accountId,omitempty"`                // length 10
	AccountTag               string  `json:"accountTag,omitempty"`               // length 50
	AccountName              string  `json:"accountName,omitempty"`              // length 50
	AccountNumber            string  `json:"accountNumber,omitempty"`            // length 50
	AccountType              string  `json:"accountType,omitempty"`              // length 50
	AccountStatus            string  `json:"accountStatus,omitempty"`            // length 50
	AccountBalance           float64 `json:"accountBalance,omitempty"`           // length 15
	CreatedDate              string  `json:"createdDate,omitempty"`              // length 34
	ClosedDate               string  `json:"closedDate,omitempty"`               // length 34
	TargetDate               string  `json:"targetDate,omitempty"`               // length 8
	TargetAmount             float64 `json:"targetAmount,omitempty"`             // length 15
	Category                 string  `json:"category,omitempty"`                 // length 50
	Subcategory              string  `json:"subcategory,omitempty"`              // length 50
	TargetMetDate            string  `json:"targetMetDate,omitempty"`            // length 34
	TargetMetPercent         float64 `json:"targetMetPercent,omitempty"`         // length 15
	IsPrimary                bool    `json:"isPrimary,omitempty"`                // length 1
	PrimaryCustomerId        int64   `json:"primaryCustomerId,omitempty"`        // length 10
	InterestRate             float64 `json:"interestRate,omitempty"`             // length 15
	ProductId                int64   `json:"productId,omitempty"`                // length 10
	AvailableBalance         float64 `json:"availableBalance,omitempty"`         // length 15
	PendingBalance           float64 `json:"pendingBalance,omitempty"`           // length 15
	AccountLockCode          string  `json:"accountLockCode,omitempty"`          // length 3
	AccountLockEffectiveDate string  `json:"accountLockEffectiveDate,omitempty"` // length 34

}

func (mdl *AccountBalanceFileModel) IsTabDelimitedFile() bool {
	return false
}

func (mdl *AccountBalanceFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *AccountBalanceFileModel) GetJsonStruct(preserve bool) interface{} {

	if preserve {
		return mdl
	} else {
		return &TypedAccountBalanceFileModel{
			CustomerId:               mdl.Typed_CustomerId,
			CustomerTag:              mdl.CustomerTag,
			AccountId:                mdl.Typed_AccountId,
			AccountTag:               mdl.AccountTag,
			AccountName:              mdl.AccountName,
			AccountNumber:            mdl.AccountNumber,
			AccountType:              mdl.AccountType,
			AccountStatus:            mdl.AccountStatus,
			AccountBalance:           mdl.Typed_AccountBalance,
			CreatedDate:              mdl.CreatedDate,
			ClosedDate:               mdl.ClosedDate,
			TargetDate:               mdl.TargetDate,
			TargetAmount:             mdl.Typed_TargetAmount,
			Category:                 mdl.Category,
			Subcategory:              mdl.Subcategory,
			TargetMetDate:            mdl.TargetMetDate,
			TargetMetPercent:         mdl.Typed_TargetMetPercent,
			IsPrimary:                mdl.Typed_IsPrimary,
			PrimaryCustomerId:        mdl.Typed_PrimaryCustomerId,
			InterestRate:             mdl.Typed_InterestRate,
			ProductId:                mdl.Typed_ProductId,
			AvailableBalance:         mdl.Typed_AvailableBalance,
			PendingBalance:           mdl.Typed_PendingBalance,
			AccountLockCode:          mdl.AccountLockCode,
			AccountLockEffectiveDate: mdl.AccountLockEffectiveDate,
		}

	}

}

func (mdl *AccountBalanceFileModel) GetValues() []string {
	return []string{
		mdl.CustomerId,
		mdl.CustomerTag,
		mdl.AccountId,
		mdl.AccountTag,
		mdl.AccountName,
		mdl.AccountNumber,
		mdl.AccountType,
		mdl.AccountStatus,
		mdl.AccountBalance,
		mdl.CreatedDate,
		mdl.ClosedDate,
		mdl.TargetDate,
		mdl.TargetAmount,
		mdl.Category,
		mdl.Subcategory,
		mdl.TargetMetDate,
		mdl.TargetMetPercent,
		mdl.IsPrimary,
		mdl.PrimaryCustomerId,
		mdl.InterestRate,
		mdl.ProductId,
		mdl.AvailableBalance,
		mdl.PendingBalance,
		mdl.AccountLockCode,
		mdl.AccountLockEffectiveDate,
	}
}

func (mdl *AccountBalanceFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	pos, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.CustomerTag, line, pos, 50, preserve, err)
	pos, err = ConsumeIntField(&mdl.AccountId, &mdl.Typed_AccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.AccountTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.AccountName, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.AccountNumber, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.AccountType, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.AccountStatus, line, pos, 50, preserve, err)
	pos, err = ConsumeFloatField(&mdl.AccountBalance, &mdl.Typed_AccountBalance, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.CreatedDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.ClosedDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.TargetDate, line, pos, 8, preserve, err)
	pos, err = ConsumeFloatField(&mdl.TargetAmount, &mdl.Typed_TargetAmount, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.Category, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.Subcategory, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.TargetMetDate, line, pos, 34, preserve, err)
	pos, err = ConsumeFloatField(&mdl.TargetMetPercent, &mdl.Typed_TargetMetPercent, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeBoolField(&mdl.IsPrimary, &mdl.Typed_IsPrimary, line, pos, 1, "Y", preserve, err)
	pos, err = ConsumeIntField(&mdl.PrimaryCustomerId, &mdl.Typed_PrimaryCustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeFloatField(&mdl.InterestRate, &mdl.Typed_InterestRate, line, pos, 15, 0, 11, preserve, err)
	pos, err = ConsumeIntField(&mdl.ProductId, &mdl.Typed_ProductId, line, pos, 10, preserve, err)
	pos, err = ConsumeFloatField(&mdl.AvailableBalance, &mdl.Typed_AvailableBalance, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.PendingBalance, &mdl.Typed_PendingBalance, line, pos, 15, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.AccountLockCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.AccountLockEffectiveDate, line, pos, 34, preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
