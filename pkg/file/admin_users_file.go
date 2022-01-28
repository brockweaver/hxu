package file

import (
	"fmt"
	"strings"
)

type AdminUsersFileModel struct {
	UserId         string `json:"userId,omitempty"`         // length 10
	Email          string `json:"email,omitempty"`          // length 255
	FirstName      string `json:"firstName,omitempty"`      // length 255
	LastName       string `json:"lastName,omitempty"`       // length 255
	Phone          string `json:"phone,omitempty"`          // length 255
	EffectiveDate  string `json:"effectiveDate,omitempty"`  // length 34
	IsActive       string `json:"isActive,omitempty"`       // length 10
	CreatedDate    string `json:"createdDate,omitempty"`    // length 34
	TerminatedDate string `json:"terminatedDate,omitempty"` // length 8
	ProgramId      string `json:"programId,omitempty"`      // length 10
	ProgramName    string `json:"programName,omitempty"`    // length 50
	CreatorUserId  string `json:"creatorUserId,omitempty"`  // length 10
	CreatorEmail   string `json:"creatorEmail,omitempty"`   // length 255

	Typed_UserId        int64 `json:"-"`
	Typed_IsActive      bool  `json:"-"`
	Typed_ProgramId     int64 `json:"-"`
	Typed_CreatorUserId int64 `json:"-"`
}

type TypedAdminUsersFileModel struct {
	UserId         int64  `json:"userId,omitempty"`         // length 10
	Email          string `json:"email,omitempty"`          // length 255
	FirstName      string `json:"firstName,omitempty"`      // length 255
	LastName       string `json:"lastName,omitempty"`       // length 255
	Phone          string `json:"phone,omitempty"`          // length 255
	EffectiveDate  string `json:"effectiveDate,omitempty"`  // length 34
	IsActive       bool   `json:"isActive,omitempty"`       // length 10
	CreatedDate    string `json:"createdDate,omitempty"`    // length 34
	TerminatedDate string `json:"terminatedDate,omitempty"` // length 8
	ProgramId      int64  `json:"programId,omitempty"`      // length 10
	ProgramName    string `json:"programName,omitempty"`    // length 50
	CreatorUserId  int64  `json:"creatorUserId,omitempty"`  // length 10
	CreatorEmail   string `json:"creatorEmail,omitempty"`   // length 255
}

func (mdl *AdminUsersFileModel) IsTabDelimitedFile() bool {
	return true
}

func (mdl *AdminUsersFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *AdminUsersFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return &TypedAdminUsersFileModel{
			UserId:         mdl.Typed_UserId,
			Email:          mdl.Email,
			FirstName:      mdl.FirstName,
			LastName:       mdl.LastName,
			Phone:          mdl.Phone,
			EffectiveDate:  mdl.EffectiveDate,
			IsActive:       mdl.Typed_IsActive,
			CreatedDate:    mdl.CreatedDate,
			TerminatedDate: mdl.TerminatedDate,
			ProgramId:      mdl.Typed_ProgramId,
			ProgramName:    mdl.ProgramName,
			CreatorUserId:  mdl.Typed_CreatorUserId,
			CreatorEmail:   mdl.CreatorEmail,
		}
	}
}

func (mdl *AdminUsersFileModel) GetValues() []string {
	return []string{
		mdl.UserId,
		mdl.Email,
		mdl.FirstName,
		mdl.LastName,
		mdl.Phone,
		mdl.EffectiveDate,
		mdl.IsActive,
		mdl.CreatedDate,
		mdl.TerminatedDate,
		mdl.ProgramId,
		mdl.ProgramName,
		mdl.CreatorUserId,
		mdl.CreatorEmail,
	}
}

func (mdl *AdminUsersFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	// special: this is a tab-delimited file
	splits := strings.Split(line, "\t")

	const expectedFieldCount = 13
	actualFieldCount := len(splits)

	if strict && actualFieldCount != expectedFieldCount {
		err = fmt.Errorf("expected exactly %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else if !strict && actualFieldCount < expectedFieldCount {
		err = fmt.Errorf("expected at least %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else {

		_, err = ConsumeIntField(&mdl.UserId, &mdl.Typed_UserId, splits[0], pos, len(splits[0]), preserve, err)
		_, err = ConsumeField(&mdl.Email, splits[1], pos, len(splits[1]), preserve, err)
		_, err = ConsumeField(&mdl.FirstName, splits[2], pos, len(splits[2]), preserve, err)
		_, err = ConsumeField(&mdl.LastName, splits[3], pos, len(splits[3]), preserve, err)
		_, err = ConsumeField(&mdl.Phone, splits[4], pos, len(splits[4]), preserve, err)
		_, err = ConsumeField(&mdl.EffectiveDate, splits[5], pos, len(splits[5]), preserve, err)
		_, err = ConsumeBoolField(&mdl.IsActive, &mdl.Typed_IsActive, splits[6], pos, len(splits[6]), "1", preserve, err)
		_, err = ConsumeField(&mdl.CreatedDate, splits[7], pos, len(splits[7]), preserve, err)
		_, err = ConsumeField(&mdl.TerminatedDate, splits[8], pos, len(splits[8]), preserve, err)
		_, err = ConsumeIntField(&mdl.ProgramId, &mdl.Typed_ProgramId, splits[9], pos, len(splits[9]), preserve, err)
		_, err = ConsumeField(&mdl.ProgramName, splits[10], pos, len(splits[10]), preserve, err)
		_, err = ConsumeIntField(&mdl.CreatorUserId, &mdl.Typed_CreatorUserId, splits[11], pos, len(splits[11]), preserve, err)
		_, err = ConsumeField(&mdl.CreatorEmail, splits[12], pos, len(splits[12]), preserve, err)

	}

	return err
}
