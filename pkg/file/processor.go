package file

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/structs"
)

type EmitFormat int64

const (
	FormatJsonLines EmitFormat = iota
	FormatCSV
	FormatTsv
	FormatVerifyOnly
)

type FileParser interface {
	ParseNative(input string, strict bool, preserve bool) error
	GetValues() []string
	GetJsonStruct(preserve bool) interface{}
	IsTabDelimitedFile() bool
	HeaderFieldCount() int
}

func GetNames(v interface{}) []string {
	names := structs.Names(v)

	// remove those that start with "Typed_"
	rv := make([]string, 0)
	for _, v := range names {
		if !strings.HasPrefix(v, "Typed_") {
			rv = append(rv, v)
		}
	}

	return rv

}

func Process(input *bufio.Scanner, output *bufio.Writer, format EmitFormat, strict bool, preserve bool, noheader bool, windows bool) error {

	// first line is always header line.
	if !input.Scan() {
		// no content, nothing to do.
		return nil
	}

	lineEnding := "\n"
	if windows {
		lineEnding = "\r\n"
	}

	// helix files use windows-style line endings
	// (carriage return + line feed) so we must trim off the trailing \r
	headerLine := strings.TrimRight(input.Text(), "\r\n")

	fp, err := getFileParser(headerLine)
	if err != nil {
		return err
	}

	// if fp was a struct here, we wouldn't need Indirect. But since it's an interface, we need indirect.
	fpType := reflect.Indirect(reflect.ValueOf(fp)).Type()
	fpTypeName := fpType.Name()

	hdr := &HeaderModel{}
	if err = hdr.ParseNative(headerLine, preserve, fp); err != nil {
		return err
	}

	count := int64(0)
	var b []byte

	if format == FormatVerifyOnly {
		// should not emit anything; only parsing file to verify its contents.
		for input.Scan() {
			line := input.Text()
			if err = fp.ParseNative(line, strict, preserve); err != nil {
				return fmt.Errorf("parser error in file type %v on content line %v: %v", fpTypeName, count+1, err)
			}
			count++
		}
	} else if format == FormatJsonLines {

		// emit header
		if !noheader {
			b, _ = json.Marshal(hdr.GetJsonStruct(preserve))
			output.Write(append(b, lineEnding...))
		}

		// emit content
		for input.Scan() {
			line := input.Text()
			if err = fp.ParseNative(line, strict, preserve); err != nil {
				return fmt.Errorf("parser error in file type %v on content line %v: %v", fpTypeName, count+1, err)
			}
			// note fp.GetValues() is not necessary here since json.Marshal handles those details
			b, _ = json.Marshal(fp.GetJsonStruct(preserve))
			output.Write(append(b, lineEnding...))
			count++
		}
	} else {
		// must be CSV or TAB. either way we do almost exactly the same thing.

		c := csv.NewWriter(output)

		// if tab output, set "Comma" appropriately. otherwise defaults to comma.
		if format == FormatTsv {
			c.Comma = '\t'
		}

		c.UseCRLF = windows

		// emit header (not actual header content; content field names)
		if !noheader {
			_ = c.Write(GetNames(fp))
		}

		// emit content
		for input.Scan() {
			line := input.Text()
			if err = fp.ParseNative(line, strict, preserve); err != nil {
				return fmt.Errorf("parser error in file type %v on content line %v: %v", fpTypeName, count+1, err)
			}
			c.Write(fp.GetValues())
			count++
		}
	}

	if strict && count != hdr.Typed_RecordCount {
		return fmt.Errorf("in file type %v, header count of %v lines does not match actual line count of %v", fpTypeName, hdr.RecordCount, count)
	}

	// be sure to flush to underlying storage device
	if err = output.Flush(); err != nil {
		return err
	}

	return input.Err()
}

func getFileParser(headerLine string) (FileParser, error) {

	test := strings.ToUpper(headerLine)

	if strings.Contains(test, "_ACCOUNTBALANCE.TXT") {
		// this is an account balance file. call its handler.
		return &AccountBalanceFileModel{}, nil
	} else if strings.Contains(test, "_ACHTRANSACTION.TXT") {
		return &AchTransactionFileModel{}, nil
	} else if strings.Contains(test, "_ADMINUSERS.TXT") {
		return &AdminUsersFileModel{}, nil
	} else if strings.Contains(test, "_ADMINLOGINACTIVITY.TXT") {
		return &AdminLoginActivityFileModel{}, nil
	} else if strings.Contains(test, "_ADMINCUSTOMERSEARCHACTIVITY.TXT") {
		return &AdminCustomerSearchActivityFileModel{}, nil
	} else if strings.Contains(test, "_ADMINWEBUSAGEACTIVITY.TXT") {
		return &AdminWebUsageActivityFileModel{}, nil
	} else if strings.Contains(test, "_BULKACCOUNTLOCK.TXT") {
		return &BulkAccountLockRequestFileModel{}, nil
	} else if strings.Contains(test, "_BULKACCOUNTLOCKRESPONSE.TXT") {
		return &BulkAccountLockResponseFileModel{}, nil
	} else if strings.Contains(test, "_BULKACCOUNTUNLOCK.TXT") {
		return &BulkAccountUnlockRequestFileModel{}, nil
	} else if strings.Contains(test, "_BULKACCOUNTUNLOCKRESPONSE.TXT") {
		return &BulkAccountUnlockResponseFileModel{}, nil
	} else if strings.Contains(test, "_BULKTRANSFERINITIATE.TXT") {
		return &BulkTransferInitiateFileModel{}, nil
	} else if strings.Contains(test, "_BULKTRANSFER.TXT") {
		return &BulkTransferRequestFileModel{}, nil
	} else if strings.Contains(test, "_BULKTRANSFERRESPONSE.TXT") {
		return &BulkTransferResponseFileModel{}, nil
	} else if strings.Contains(test, "_DEBITCARDEVENTNOTIFICATION.TXT") {
		return &CardEventNotificationFileModel{}, nil
	} else if strings.Contains(test, "_DEBITCARDTRANSACTION.TXT") {
		return &CardTransactionFileModel{}, nil
	} else if strings.Contains(test, "_CUSTOMER.TXT") {
		return &CustomerFileModel{}, nil
	} else if strings.Contains(test, "_CUSTOMERREGISTRATION.TXT") {
		return &CustomerRegistrationFileModel{}, nil
	} else if strings.Contains(test, "_EVENTNOTIFICATION.TXT") {
		return &EventNotificationFileModel{}, nil
	} else if strings.Contains(test, "_EXTERNALACCOUNTEXPORT.TXT") {
		return &ExternalAccountFileModel{}, nil
	} else if strings.Contains(test, "_POSTEDTRANSACTION.TXT") {
		return &PostedTransactionFileModel{}, nil
	} else if strings.Contains(test, "_STATEMENTEVENTNOTIFICATION.TXT") {
		return &StatementEventNotificationFileModel{}, nil
	} else if strings.Contains(test, "_TRIALBALANCEEXPORT_") {
		return &TrialBalanceFileModel{}, nil
	} else if strings.Contains(test, "_ARCHIVEDCUSTOMER.TXT") {
		return &CustomerArchivedFileModel{}, nil
	}

	return nil, fmt.Errorf("could not find FileParser instance for header=%s", test)
}

func consume(line string, start int, length int, defaultValue string) (string, int, error) {
	var position int
	rv := defaultValue
	var err error
	if len(line) >= start+length {
		// line contains entire field
		rv = line[start : start+length]
	} else if len(line) > start && len(line) < start+length {
		// line contains part of the field (invalid file, pull what we can)
		rv = line[start:]
		err = fmt.Errorf("file width is %v, field starting at byte %v of length %v is only partially available", len(line), start, length)
	} else if len(line) <= start {
		// we're short the entire field.
		err = fmt.Errorf("file width is %v, field starting at byte %v of length %v is not available", len(line), start, length)
	} else {
		// should never get here.
		err = fmt.Errorf("UNEXPECTED CASE: file width is %v, field starting at byte %v of length %v is not available", len(line), start, length)
	}

	// regardless of file content, always jump to correct byte offset
	position = start + length
	return strings.TrimSpace(rv), position, err
}

func ConsumeField(origField *string, line string, start int, length int, preserve bool, previousError error) (int, error) {
	var position int
	if previousError != nil {
		return position, previousError
	}

	var err error
	*origField, position, err = consume(line, start, length, "")
	return position, err
}

func ConsumeBoolField(origField *string, typedField *bool, line string, start int, length int, trueValue string, preserve bool, previousError error) (int, error) {
	var position int
	var err error
	if position, err = ConsumeField(origField, line, start, length, preserve, previousError); err != nil {
		return position, err
	} else {

		*typedField = strings.EqualFold(*origField, trueValue)

		if !preserve {
			*origField = fmt.Sprintf("%v", *typedField)
		}
	}
	return position, nil
}

func ConsumeIntField(origField *string, typedField *int64, line string, start int, length int, preserve bool, previousError error) (int, error) {
	var position int
	var err error
	if position, err = ConsumeField(origField, line, start, length, preserve, previousError); err != nil {
		return position, err
	} else {
		if *typedField, err = strconv.ParseInt(*origField, 10, 64); err != nil {
			//*typedField = 0
			return position, err
		}
		if !preserve {
			*origField = fmt.Sprintf("%v", *typedField)
		}
	}
	return position, nil
}

func ConsumeFloatField(origField *string, typedField *float64, line string, start int, length int, assumedDigitCount int, significantDigitCount int, preserve bool, previousError error) (int, error) {
	var position int
	var err error
	if position, err = ConsumeField(origField, line, start, length, preserve, previousError); err != nil {
		return position, err
	} else {

		val := *origField

		// negative fields are sometimes zero-padded left of the negative sign.
		// e.g.:  "-1234" field of length 10 will appear as "00000-1234".
		// if this is the case, we want to trim off the zeroes left of the negative sign prior to attempting to parse that float.
		if strings.Contains(val, "-") {
			val = strings.TrimLeft(val, "0")
		}

		if *typedField, err = strconv.ParseFloat(val, 64); err != nil {
			//*typedField = 0.0
			return position, err
		}

		if assumedDigitCount > 0 {
			*typedField = *typedField / math.Pow10(assumedDigitCount)
		}

		if !preserve {
			f := "%f"
			if significantDigitCount > 0 {
				f = "%." + strconv.Itoa(significantDigitCount) + "f"
			}
			*origField = fmt.Sprintf(f, *typedField)
		}
	}

	return position, nil
}
