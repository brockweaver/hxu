package file

import (
	"fmt"
	"strings"
)

type TrialBalanceFileModel struct {
	ProgramName                  string `json:"programName,omitempty"`                  // length 50
	ClientName                   string `json:"clientName,omitempty"`                   // length 50
	CustomerId                   string `json:"customerId,omitempty"`                   // length 10
	FirstName                    string `json:"firstName,omitempty"`                    // length 64
	MiddleName                   string `json:"middleName,omitempty"`                   // length 64
	LastName                     string `json:"lastName,omitempty"`                     // length 128
	FullName                     string `json:"fullName,omitempty"`                     // length 256
	AccountCreatedDate           string `json:"accountCreatedDate,omitempty"`           // length 34
	AccountName                  string `json:"accountName,omitempty"`                  // length 50
	AccountNumber                string `json:"accountNumber,omitempty"`                // length 17
	EffectiveDateEndingBalance   string `json:"effectiveDateEndingBalance,omitempty"`   // length ???
	EffectiveDateInterestAccrued string `json:"effectiveDateInterestAccrued,omitempty"` // length ???
	PeriodAverageDailyBalance    string `json:"periodAverageDailyBalance,omitempty"`    // length ???
	PeriodInterestAccrued        string `json:"periodInterestAccrued,omitempty"`        // length ???
	PeriodRoundedInterestAccrued string `json:"periodRoundedInterestAccrued,omitempty"` // length ???
	PeriodInterestPaid           string `json:"periodInterestPaid,omitempty"`           // length ???
	YearToDateInterestPaid       string `json:"yearToDateInterestPaid,omitempty"`       // length ???
	InterestRate                 string `json:"interestRate,omitempty"`                 // length ???
	BeneficiaryCount             string `json:"beneficiaryCount,omitempty"`             // length ???
	ProductName                  string `json:"productName,omitempty"`                  // length 50
	TaxId                        string `json:"taxId,omitempty"`                        // length 10
	AccountId                    string `json:"accountId,omitempty"`                    // length 10
	ProductId                    string `json:"ProductId,omitempty"`                    // length 10

	Typed_CustomerId                   int64   `json:"-"`
	Typed_EffectiveDateEndingBalance   float64 `json:"-"`
	Typed_EffectiveDateInterestAccrued float64 `json:"-"`
	Typed_PeriodAverageDailyBalance    float64 `json:"-"`
	Typed_PeriodInterestAccrued        float64 `json:"-"`
	Typed_PeriodRoundedInterestAccrued float64 `json:"-"`
	Typed_PeriodInterestPaid           float64 `json:"-"`
	Typed_YearToDateInterestPaid       float64 `json:"-"`
	Typed_InterestRate                 float64 `json:"-"`
	Typed_BeneficiaryCount             int64   `json:"-"`
	Typed_AccountId                    int64   `json:"-"`
	Typed_ProductId                    int64   `json:"-"`
}

type TypedTrialBalanceFileModel struct {
	ProgramName                  string  `json:"programName,omitempty"`                  // length 50
	ClientName                   string  `json:"clientName,omitempty"`                   // length 50
	CustomerId                   int64   `json:"customerId,omitempty"`                   // length 10
	FirstName                    string  `json:"firstName,omitempty"`                    // length 64
	MiddleName                   string  `json:"middleName,omitempty"`                   // length 64
	LastName                     string  `json:"lastName,omitempty"`                     // length 128
	FullName                     string  `json:"fullName,omitempty"`                     // length 256
	AccountCreatedDate           string  `json:"accountCreatedDate,omitempty"`           // length 34
	AccountName                  string  `json:"accountName,omitempty"`                  // length 50
	AccountNumber                string  `json:"accountNumber,omitempty"`                // length 17
	EffectiveDateEndingBalance   float64 `json:"effectiveDateEndingBalance,omitempty"`   // length ???
	EffectiveDateInterestAccrued float64 `json:"effectiveDateInterestAccrued,omitempty"` // length ???
	PeriodAverageDailyBalance    float64 `json:"periodAverageDailyBalance,omitempty"`    // length ???
	PeriodInterestAccrued        float64 `json:"periodInterestAccrued,omitempty"`        // length ???
	PeriodRoundedInterestAccrued float64 `json:"periodRoundedInterestAccrued,omitempty"` // length ???
	PeriodInterestPaid           float64 `json:"periodInterestPaid,omitempty"`           // length ???
	YearToDateInterestPaid       float64 `json:"yearToDateInterestPaid,omitempty"`       // length ???
	InterestRate                 float64 `json:"interestRate,omitempty"`                 // length ???
	BeneficiaryCount             int64   `json:"beneficiaryCount,omitempty"`             // length ???
	ProductName                  string  `json:"productName,omitempty"`                  // length 50
	TaxId                        string  `json:"taxId,omitempty"`                        // length 10
	AccountId                    int64   `json:"accountId,omitempty"`                    // length 10
	ProductId                    int64   `json:"ProductId,omitempty"`                    // length 10
}

func (mdl *TrialBalanceFileModel) IsTabDelimitedFile() bool {
	return true
}

func (mdl *TrialBalanceFileModel) HeaderFieldCount() int {
	return 4
}

func (mdl *TrialBalanceFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedTrialBalanceFileModel{
			ProgramName:                  mdl.ProgramName,
			ClientName:                   mdl.ClientName,
			CustomerId:                   mdl.Typed_CustomerId,
			FirstName:                    mdl.FirstName,
			MiddleName:                   mdl.MiddleName,
			LastName:                     mdl.LastName,
			FullName:                     mdl.FullName,
			AccountCreatedDate:           mdl.AccountCreatedDate,
			AccountName:                  mdl.AccountName,
			AccountNumber:                mdl.AccountNumber,
			EffectiveDateEndingBalance:   mdl.Typed_EffectiveDateEndingBalance,
			EffectiveDateInterestAccrued: mdl.Typed_EffectiveDateInterestAccrued,
			PeriodAverageDailyBalance:    mdl.Typed_PeriodAverageDailyBalance,
			PeriodInterestAccrued:        mdl.Typed_PeriodInterestAccrued,
			PeriodRoundedInterestAccrued: mdl.Typed_PeriodRoundedInterestAccrued,
			PeriodInterestPaid:           mdl.Typed_PeriodInterestPaid,
			YearToDateInterestPaid:       mdl.Typed_YearToDateInterestPaid,
			InterestRate:                 mdl.Typed_InterestRate,
			BeneficiaryCount:             mdl.Typed_BeneficiaryCount,
			ProductName:                  mdl.ProductName,
			TaxId:                        mdl.TaxId,
			AccountId:                    mdl.Typed_AccountId,
			ProductId:                    mdl.Typed_ProductId,
		}
	}
}

func (mdl *TrialBalanceFileModel) GetValues() []string {
	return []string{
		mdl.ProgramName,
		mdl.ClientName,
		mdl.CustomerId,
		mdl.FirstName,
		mdl.MiddleName,
		mdl.LastName,
		mdl.FullName,
		mdl.AccountCreatedDate,
		mdl.AccountName,
		mdl.AccountNumber,
		mdl.EffectiveDateEndingBalance,
		mdl.EffectiveDateInterestAccrued,
		mdl.PeriodAverageDailyBalance,
		mdl.PeriodInterestAccrued,
		mdl.PeriodRoundedInterestAccrued,
		mdl.PeriodInterestAccrued,
		mdl.YearToDateInterestPaid,
		mdl.InterestRate,
		mdl.BeneficiaryCount,
		mdl.ProductName,
		mdl.TaxId,
		mdl.AccountId,
		mdl.ProductId,
	}
}

func (mdl *TrialBalanceFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	// special: this is a tab-delimited file
	splits := strings.Split(line, "\t")

	const expectedFieldCount = 23
	actualFieldCount := len(splits)

	if strict && actualFieldCount != expectedFieldCount {
		err = fmt.Errorf("expected exactly %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else if !strict && actualFieldCount < expectedFieldCount {
		err = fmt.Errorf("expected at least %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else {

		_, err = ConsumeField(&mdl.ProgramName, splits[0], pos, len(splits[0]), preserve, err)
		_, err = ConsumeField(&mdl.ClientName, splits[1], pos, len(splits[1]), preserve, err)
		_, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, splits[2], pos, len(splits[2]), preserve, err)
		_, err = ConsumeField(&mdl.FirstName, splits[3], pos, len(splits[3]), preserve, err)
		_, err = ConsumeField(&mdl.MiddleName, splits[4], pos, len(splits[4]), preserve, err)
		_, err = ConsumeField(&mdl.LastName, splits[5], pos, len(splits[5]), preserve, err)
		_, err = ConsumeField(&mdl.FullName, splits[6], pos, len(splits[6]), preserve, err)
		_, err = ConsumeField(&mdl.AccountCreatedDate, splits[7], pos, len(splits[7]), preserve, err)
		_, err = ConsumeField(&mdl.AccountName, splits[8], pos, len(splits[8]), preserve, err)
		_, err = ConsumeField(&mdl.AccountNumber, splits[9], pos, len(splits[9]), preserve, err)
		_, err = ConsumeFloatField(&mdl.EffectiveDateEndingBalance, &mdl.Typed_EffectiveDateEndingBalance, splits[10], pos, len(splits[10]), 0, 2, preserve, err)
		_, err = ConsumeFloatField(&mdl.EffectiveDateInterestAccrued, &mdl.Typed_EffectiveDateInterestAccrued, splits[11], pos, len(splits[11]), 0, 4, preserve, err)
		_, err = ConsumeFloatField(&mdl.PeriodAverageDailyBalance, &mdl.Typed_PeriodAverageDailyBalance, splits[12], pos, len(splits[12]), 0, 2, preserve, err)
		_, err = ConsumeFloatField(&mdl.PeriodInterestAccrued, &mdl.Typed_PeriodInterestAccrued, splits[13], pos, len(splits[13]), 0, 4, preserve, err)
		_, err = ConsumeFloatField(&mdl.PeriodRoundedInterestAccrued, &mdl.Typed_PeriodRoundedInterestAccrued, splits[14], pos, len(splits[14]), 0, 4, preserve, err)
		_, err = ConsumeFloatField(&mdl.PeriodInterestPaid, &mdl.Typed_PeriodInterestPaid, splits[15], pos, len(splits[15]), 0, 2, preserve, err)
		_, err = ConsumeFloatField(&mdl.YearToDateInterestPaid, &mdl.Typed_YearToDateInterestPaid, splits[16], pos, len(splits[16]), 0, 2, preserve, err)
		_, err = ConsumeFloatField(&mdl.InterestRate, &mdl.Typed_InterestRate, splits[17], pos, len(splits[17]), 0, 4, preserve, err)
		_, err = ConsumeIntField(&mdl.BeneficiaryCount, &mdl.Typed_BeneficiaryCount, splits[18], pos, len(splits[18]), preserve, err)
		_, err = ConsumeField(&mdl.ProductName, splits[19], pos, len(splits[19]), preserve, err)
		_, err = ConsumeField(&mdl.TaxId, splits[20], pos, len(splits[20]), preserve, err)
		_, err = ConsumeIntField(&mdl.AccountId, &mdl.Typed_AccountId, splits[21], pos, len(splits[21]), preserve, err)
		_, err = ConsumeIntField(&mdl.ProductId, &mdl.Typed_ProductId, splits[22], pos, len(splits[22]), preserve, err)

	}

	return err
}
