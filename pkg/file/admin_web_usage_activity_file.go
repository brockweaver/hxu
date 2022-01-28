package file

import (
	"fmt"
	"strings"
)

type AdminWebUsageActivityFileModel struct {
	UserId       string `json:"userId,omitempty"`       // length 10
	Url          string `json:"url,omitempty"`          // length 2000
	Date         string `json:"date,omitempty"`         // length 34
	ProgramId    string `json:"programId,omitempty"`    // length 10
	ProgramName  string `json:"programName,omitempty"`  // length 50
	EmailAddress string `json:"emailAddress,omitempty"` // length 255

	Typed_UserId    int64 `json:"-"`
	Typed_ProgramId int64 `json:"-"`
}

type TypeAdminWebUsageActivityFileModel struct {
	UserId       int64  `json:"userId,omitempty"`       // length 10
	Url          string `json:"url,omitempty"`          // length 2000
	Date         string `json:"date,omitempty"`         // length 34
	ProgramId    int64  `json:"programId,omitempty"`    // length 10
	ProgramName  string `json:"programName,omitempty"`  // length 50
	EmailAddress string `json:"emailAddress,omitempty"` // length 255
}

func (mdl *AdminWebUsageActivityFileModel) IsTabDelimitedFile() bool {
	return true
}

func (mdl *AdminWebUsageActivityFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *AdminWebUsageActivityFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypeAdminWebUsageActivityFileModel{
			UserId:       mdl.Typed_UserId,
			Url:          mdl.Url,
			Date:         mdl.Date,
			ProgramId:    mdl.Typed_ProgramId,
			ProgramName:  mdl.ProgramName,
			EmailAddress: mdl.EmailAddress,
		}
	}
}

func (mdl *AdminWebUsageActivityFileModel) GetValues() []string {
	return []string{
		mdl.UserId,
		mdl.Url,
		mdl.Date,
		mdl.ProgramId,
		mdl.ProgramName,
		mdl.EmailAddress,
	}
}

func (mdl *AdminWebUsageActivityFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	// special: this is a tab-delimited file
	splits := strings.Split(line, "\t")

	const expectedFieldCount = 6
	actualFieldCount := len(splits)

	if strict && actualFieldCount != expectedFieldCount {
		err = fmt.Errorf("expected exactly %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else if !strict && actualFieldCount < expectedFieldCount {
		err = fmt.Errorf("expected at least %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else {

		_, err = ConsumeIntField(&mdl.UserId, &mdl.Typed_UserId, splits[0], pos, len(splits[0]), preserve, err)
		_, err = ConsumeField(&mdl.Url, splits[1], pos, len(splits[1]), preserve, err)
		_, err = ConsumeField(&mdl.Date, splits[2], pos, len(splits[2]), preserve, err)
		_, err = ConsumeIntField(&mdl.ProgramId, &mdl.Typed_ProgramId, splits[3], pos, len(splits[3]), preserve, err)
		_, err = ConsumeField(&mdl.ProgramName, splits[4], pos, len(splits[4]), preserve, err)
		_, err = ConsumeField(&mdl.EmailAddress, splits[5], pos, len(splits[5]), preserve, err)

	}

	return err
}
