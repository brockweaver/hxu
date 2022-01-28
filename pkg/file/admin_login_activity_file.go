package file

import (
	"fmt"
	"strings"
)

type AdminLoginActivityFileModel struct {
	UserId        string `json:"userId,omitempty"`        // length 10
	UserName      string `json:"userName,omitempty"`      // length 255
	RemoteAddress string `json:"remoteAddress,omitempty"` // length 200
	Headers       string `json:"headers,omitempty"`       // length 400
	Date          string `json:"date,omitempty"`          // length 34
	Status        string `json:"status,omitempty"`        // length 1
	ProgramId     string `json:"programId,omitempty"`     // length 10
	ProgramName   string `json:"programName,omitempty"`   // length 50

	Typed_UserId    int64 `json:"-"`
	Typed_ProgramId int64 `json:"-"`
}

type TypedAdminLoginActivityFileModel struct {
	UserId        int64  `json:"userId,omitempty"`        // length 10
	UserName      string `json:"userName,omitempty"`      // length 255
	RemoteAddress string `json:"remoteAddress,omitempty"` // length 200
	Headers       string `json:"headers,omitempty"`       // length 400
	Date          string `json:"date,omitempty"`          // length 34
	Status        string `json:"status,omitempty"`        // length 1
	ProgramId     int64  `json:"programId,omitempty"`     // length 10
	ProgramName   string `json:"programName,omitempty"`   // length 50

}

func (mdl *AdminLoginActivityFileModel) IsTabDelimitedFile() bool {
	return true
}

func (mdl *AdminLoginActivityFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *AdminLoginActivityFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return &TypedAdminLoginActivityFileModel{
			UserId:        mdl.Typed_UserId,
			UserName:      mdl.UserName,
			RemoteAddress: mdl.RemoteAddress,
			Headers:       mdl.Headers,
			Date:          mdl.Date,
			Status:        mdl.Status,
			ProgramId:     mdl.Typed_ProgramId,
			ProgramName:   mdl.ProgramName,
		}
	}
}

func (mdl *AdminLoginActivityFileModel) GetValues() []string {
	return []string{
		mdl.UserId,
		mdl.UserName,
		mdl.RemoteAddress,
		mdl.Headers,
		mdl.Date,
		mdl.Status,
		mdl.ProgramId,
		mdl.ProgramName,
	}
}

func (mdl *AdminLoginActivityFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	// special: this is a tab-delimited file
	splits := strings.Split(line, "\t")

	const expectedFieldCount = 8
	actualFieldCount := len(splits)

	if strict && actualFieldCount != expectedFieldCount {
		err = fmt.Errorf("expected exactly %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else if !strict && actualFieldCount < expectedFieldCount {
		err = fmt.Errorf("expected at least %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else {

		_, err = ConsumeIntField(&mdl.UserId, &mdl.Typed_UserId, splits[0], pos, len(splits[0]), preserve, err)
		_, err = ConsumeField(&mdl.UserName, splits[1], pos, len(splits[1]), preserve, err)
		_, err = ConsumeField(&mdl.RemoteAddress, splits[2], pos, len(splits[2]), preserve, err)
		_, err = ConsumeField(&mdl.Headers, splits[3], pos, len(splits[3]), preserve, err)
		_, err = ConsumeField(&mdl.Date, splits[4], pos, len(splits[4]), preserve, err)
		_, err = ConsumeField(&mdl.Status, splits[5], pos, len(splits[5]), preserve, err)
		_, err = ConsumeIntField(&mdl.ProgramId, &mdl.Typed_ProgramId, splits[6], pos, len(splits[6]), preserve, err)
		_, err = ConsumeField(&mdl.ProgramName, splits[7], pos, len(splits[7]), preserve, err)

	}

	return err
}
