package file

import (
	"fmt"
	"strings"
)

type HeaderModel struct {
	RecordType        string `json:"recordType,omitempty"`        // length 1
	FileName          string `json:"fileName,omitempty"`          // length 50
	RecordCount       string `json:"recordCount,omitempty"`       // length 10
	FileCreatedDate   string `json:"fileCreatedDate,omitempty"`   // length 34
	FileEffectiveDate string `json:"fileEffectiveDate,omitempty"` // length 34
	ReferenceId       string `json:"referenceId,omitempty"`       // length 50, optional
	SuccessCount      string `json:"successCount,omitempty"`      // length 10, optional
	FailedCount       string `json:"failedCount,omitempty"`       // length 10, optional
	ProcessedCount    string `json:"processedCount,omitempty"`    // length 10, optional

	Typed_RecordCount    int64 `json:"-"`
	Typed_SuccessCount   int64 `json:"-"`
	Typed_FailedCount    int64 `json:"-"`
	Typed_ProcessedCount int64 `json:"-"`
}

type TypedHeaderModel struct {
	RecordType        string `json:"recordType,omitempty"`        // length 1
	FileName          string `json:"fileName,omitempty"`          // length 50
	RecordCount       int64  `json:"recordCount,omitempty"`       // length 10
	FileCreatedDate   string `json:"fileCreatedDate,omitempty"`   // length 34
	FileEffectiveDate string `json:"fileEffectiveDate,omitempty"` // length 34
	ReferenceId       string `json:"referenceId,omitempty"`       // length 50, optional
	SuccessCount      int64  `json:"successCount,omitempty"`      // length 10, optional
	FailedCount       int64  `json:"failedCount,omitempty"`       // length 10, optional
	ProcessedCount    int64  `json:"processedCount,omitempty"`    // length 10, optional

}

func (mdl *HeaderModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedHeaderModel{
			RecordType:        mdl.RecordType,
			FileName:          mdl.FileName,
			RecordCount:       mdl.Typed_RecordCount,
			FileCreatedDate:   mdl.FileCreatedDate,
			FileEffectiveDate: mdl.FileEffectiveDate,
			ReferenceId:       mdl.ReferenceId,
			SuccessCount:      mdl.Typed_SuccessCount,
			FailedCount:       mdl.Typed_FailedCount,
			ProcessedCount:    mdl.Typed_ProcessedCount,
		}
	}
}

func (mdl *HeaderModel) ParseNative(input string, preserve bool, parser FileParser) error {
	pos := 0
	var err error

	if parser.IsTabDelimitedFile() {
		splits := strings.Split(input, "\t")

		expectedFieldCount := parser.HeaderFieldCount()
		actualFieldCount := len(splits)

		if actualFieldCount < expectedFieldCount {
			err = fmt.Errorf("expected exactly %v header fields, got %v", expectedFieldCount, actualFieldCount)
		} else {
			switch actualFieldCount {
			case 4:
				_, err = ConsumeField(&mdl.FileName, splits[0], pos, len(splits[0]), preserve, err)
				_, err = ConsumeIntField(&mdl.RecordCount, &mdl.Typed_RecordCount, splits[1], pos, len(splits[1]), preserve, err)
				_, err = ConsumeField(&mdl.FileCreatedDate, splits[2], pos, len(splits[2]), preserve, err)
				_, err = ConsumeField(&mdl.FileEffectiveDate, splits[3], pos, len(splits[3]), preserve, err)
			case 5:
				_, err = ConsumeField(&mdl.RecordType, splits[0], pos, len(splits[0]), preserve, err)
				_, err = ConsumeField(&mdl.FileName, splits[1], pos, len(splits[1]), preserve, err)
				_, err = ConsumeIntField(&mdl.RecordCount, &mdl.Typed_RecordCount, splits[2], pos, len(splits[2]), preserve, err)
				_, err = ConsumeField(&mdl.FileCreatedDate, splits[3], pos, len(splits[3]), preserve, err)
				_, err = ConsumeField(&mdl.FileEffectiveDate, splits[4], pos, len(splits[4]), preserve, err)
			default:
				return fmt.Errorf("not implemented: %v tabbed header fields", actualFieldCount)
			}
		}
	} else {

		pos, err = ConsumeField(&mdl.RecordType, input, pos, 1, preserve, err)
		pos, err = ConsumeField(&mdl.FileName, input, pos, 50, preserve, err)
		pos, err = ConsumeIntField(&mdl.RecordCount, &mdl.Typed_RecordCount, input, pos, 10, preserve, err)
		pos, err = ConsumeField(&mdl.FileCreatedDate, input, pos, 34, preserve, err)
		pos, err = ConsumeField(&mdl.FileEffectiveDate, input, pos, 34, preserve, err)

		switch parser.HeaderFieldCount() {
		case 5:
			// nothing more to do for this case.
		case 6:
			_, err = ConsumeField(&mdl.ReferenceId, input, pos, 50, preserve, err)
		case 9:
			pos, err = ConsumeField(&mdl.ReferenceId, input, pos, 50, preserve, err)
			pos, err = ConsumeIntField(&mdl.SuccessCount, &mdl.Typed_SuccessCount, input, pos, 10, preserve, err)
			pos, err = ConsumeIntField(&mdl.FailedCount, &mdl.Typed_FailedCount, input, pos, 10, preserve, err)
			_, err = ConsumeIntField(&mdl.ProcessedCount, &mdl.Typed_ProcessedCount, input, pos, 10, preserve, err)
		default:
			return fmt.Errorf("not implemented: %v fixed-length header fields", parser.HeaderFieldCount())
		}
	}

	return err
}
