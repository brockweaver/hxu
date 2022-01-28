package file

import (
	"fmt"
	"strings"
)

type ExternalAccountFileModel struct {
	ExternalAccountId   string `json:"externalAccountId,omitempty"`   // length 10
	CustomerId          string `json:"customerId,omitempty"`          // length 10
	Tag                 string `json:"tag,omitempty"`                 // length 50
	Name                string `json:"name,omitempty"`                // length 50
	RoutingNumber       string `json:"routingNumber,omitempty"`       // length 10
	RoutingNumberMasked string `json:"routingNumberMasked,omitempty"` // length 10
	AccountNumber       string `json:"accountNumber,omitempty"`       // length 17
	AccountNumberMasked string `json:"accountNumberMasked,omitempty"` // length 17
	Type                string `json:"type,omitempty"`                // length 50
	NickName            string `json:"nickName,omitempty"`            // length 50
	Status              string `json:"status,omitempty"`              // length 20
	StatusDate          string `json:"statusDate,omitempty"`          // length 34
	LastModifiedDate    string `json:"lastModifiedDate,omitempty"`    // length 34
	NocCode             string `json:"nocCode,omitempty"`             // length 10
	IsActive            string `json:"isActive,omitempty"`            // length 1
	IsLocked            string `json:"isLocked,omitempty"`            // length 1
	LockedDate          string `json:"lockedDate,omitempty"`          // length 34
	LockedReason        string `json:"lockedReason,omitempty"`        // length 255
	CustomField1        string `json:"customField1,omitempty"`        // length 50
	CustomField2        string `json:"customField2,omitempty"`        // length 50
	CustomField3        string `json:"customField3,omitempty"`        // length 50
	CustomField4        string `json:"customField4,omitempty"`        // length 50
	CustomField5        string `json:"customField5,omitempty"`        // length 50
	BusinessName        string `json:"businessName,omitempty"`        // length 100

	Typed_ExternalAccountId int64 `json:"-"`
	Typed_CustomerId        int64 `json:"-"`
	Typed_IsActive          bool  `json:"-"`
	Typed_IsLocked          bool  `json:"-"`
}

type TypedExternalAccountFileModel struct {
	ExternalAccountId   int64  `json:"externalAccountId,omitempty"`   // length 10
	CustomerId          int64  `json:"customerId,omitempty"`          // length 10
	Tag                 string `json:"tag,omitempty"`                 // length 50
	Name                string `json:"name,omitempty"`                // length 50
	RoutingNumber       string `json:"routingNumber,omitempty"`       // length 10
	RoutingNumberMasked string `json:"routingNumberMasked,omitempty"` // length 10
	AccountNumber       string `json:"accountNumber,omitempty"`       // length 17
	AccountNumberMasked string `json:"accountNumberMasked,omitempty"` // length 17
	Type                string `json:"type,omitempty"`                // length 50
	NickName            string `json:"nickName,omitempty"`            // length 50
	Status              string `json:"status,omitempty"`              // length 20
	StatusDate          string `json:"statusDate,omitempty"`          // length 34
	LastModifiedDate    string `json:"lastModifiedDate,omitempty"`    // length 34
	NocCode             string `json:"nocCode,omitempty"`             // length 10
	IsActive            bool   `json:"isActive,omitempty"`            // length 1
	IsLocked            bool   `json:"isLocked,omitempty"`            // length 1
	LockedDate          string `json:"lockedDate,omitempty"`          // length 34
	LockedReason        string `json:"lockedReason,omitempty"`        // length 255
	CustomField1        string `json:"customField1,omitempty"`        // length 50
	CustomField2        string `json:"customField2,omitempty"`        // length 50
	CustomField3        string `json:"customField3,omitempty"`        // length 50
	CustomField4        string `json:"customField4,omitempty"`        // length 50
	CustomField5        string `json:"customField5,omitempty"`        // length 50
	BusinessName        string `json:"businessName,omitempty"`        // length 100
}

func (mdl *ExternalAccountFileModel) IsTabDelimitedFile() bool {
	return true
}

func (mdl *ExternalAccountFileModel) HeaderFieldCount() int {
	return 4
}

func (mdl *ExternalAccountFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return &TypedExternalAccountFileModel{
			ExternalAccountId:   mdl.Typed_ExternalAccountId,
			CustomerId:          mdl.Typed_CustomerId,
			Tag:                 mdl.Tag,
			Name:                mdl.Name,
			RoutingNumber:       mdl.RoutingNumber,
			RoutingNumberMasked: mdl.RoutingNumberMasked,
			AccountNumber:       mdl.AccountNumber,
			AccountNumberMasked: mdl.AccountNumberMasked,
			Type:                mdl.Type,
			NickName:            mdl.NickName,
			Status:              mdl.Status,
			StatusDate:          mdl.StatusDate,
			LastModifiedDate:    mdl.LastModifiedDate,
			NocCode:             mdl.NocCode,
			IsActive:            mdl.Typed_IsActive,
			IsLocked:            mdl.Typed_IsLocked,
			LockedDate:          mdl.LockedDate,
			LockedReason:        mdl.LockedReason,
			CustomField1:        mdl.CustomField1,
			CustomField2:        mdl.CustomField2,
			CustomField3:        mdl.CustomField3,
			CustomField4:        mdl.CustomField4,
			CustomField5:        mdl.CustomField5,
			BusinessName:        mdl.BusinessName,
		}
	}
}

func (mdl *ExternalAccountFileModel) GetValues() []string {
	return []string{
		mdl.ExternalAccountId,
		mdl.CustomerId,
		mdl.Tag,
		mdl.Name,
		mdl.RoutingNumber,
		mdl.RoutingNumberMasked,
		mdl.AccountNumber,
		mdl.AccountNumberMasked,
		mdl.Type,
		mdl.NickName,
		mdl.Status,
		mdl.StatusDate,
		mdl.LastModifiedDate,
		mdl.NocCode,
		mdl.IsActive,
		mdl.IsLocked,
		mdl.LockedDate,
		mdl.LockedReason,
		mdl.CustomField1,
		mdl.CustomField2,
		mdl.CustomField3,
		mdl.CustomField4,
		mdl.CustomField5,
		mdl.BusinessName,
	}
}

func (mdl *ExternalAccountFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	// special: this is a tab-delimited file
	splits := strings.Split(line, "\t")

	const expectedFieldCount = 24
	actualFieldCount := len(splits)

	if strict && actualFieldCount != expectedFieldCount {
		err = fmt.Errorf("expected exactly %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else if !strict && actualFieldCount < expectedFieldCount {
		err = fmt.Errorf("expected at least %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else {

		_, err = ConsumeIntField(&mdl.ExternalAccountId, &mdl.Typed_ExternalAccountId, splits[0], pos, len(splits[0]), preserve, err)
		_, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, splits[1], pos, len(splits[1]), preserve, err)
		_, err = ConsumeField(&mdl.Tag, splits[2], pos, len(splits[2]), preserve, err)
		_, err = ConsumeField(&mdl.Name, splits[3], pos, len(splits[3]), preserve, err)
		_, err = ConsumeField(&mdl.RoutingNumber, splits[4], pos, len(splits[4]), preserve, err)
		_, err = ConsumeField(&mdl.RoutingNumberMasked, splits[5], pos, len(splits[5]), preserve, err)
		_, err = ConsumeField(&mdl.AccountNumber, splits[6], pos, len(splits[6]), preserve, err)
		_, err = ConsumeField(&mdl.AccountNumberMasked, splits[7], pos, len(splits[7]), preserve, err)
		_, err = ConsumeField(&mdl.Type, splits[8], pos, len(splits[8]), preserve, err)
		_, err = ConsumeField(&mdl.NickName, splits[9], pos, len(splits[9]), preserve, err)
		_, err = ConsumeField(&mdl.Status, splits[10], pos, len(splits[10]), preserve, err)
		_, err = ConsumeField(&mdl.StatusDate, splits[11], pos, len(splits[11]), preserve, err)
		_, err = ConsumeField(&mdl.LastModifiedDate, splits[12], pos, len(splits[12]), preserve, err)
		_, err = ConsumeField(&mdl.NocCode, splits[13], pos, len(splits[13]), preserve, err)
		_, err = ConsumeBoolField(&mdl.IsActive, &mdl.Typed_IsActive, splits[14], pos, len(splits[14]), "1", preserve, err)
		_, err = ConsumeBoolField(&mdl.IsLocked, &mdl.Typed_IsLocked, splits[15], pos, len(splits[15]), "1", preserve, err)
		_, err = ConsumeField(&mdl.LockedDate, splits[16], pos, len(splits[16]), preserve, err)
		_, err = ConsumeField(&mdl.LockedReason, splits[17], pos, len(splits[17]), preserve, err)
		_, err = ConsumeField(&mdl.CustomField1, splits[18], pos, len(splits[18]), preserve, err)
		_, err = ConsumeField(&mdl.CustomField2, splits[19], pos, len(splits[19]), preserve, err)
		_, err = ConsumeField(&mdl.CustomField3, splits[20], pos, len(splits[20]), preserve, err)
		_, err = ConsumeField(&mdl.CustomField4, splits[21], pos, len(splits[21]), preserve, err)
		_, err = ConsumeField(&mdl.CustomField5, splits[22], pos, len(splits[22]), preserve, err)
		_, err = ConsumeField(&mdl.BusinessName, splits[23], pos, len(splits[23]), preserve, err)
	}

	return err
}
