package file

import (
	"fmt"
	"strings"
)

type CustomerArchivedFileModel struct {
	CustomerId                          string `json:"customerId,omitempty"`                          // offset 0, length 10
	AccountId                           string `json:"accountId,omitempty"`                           // offset 1, length 10
	FirstName                           string `json:"firstName,omitempty"`                           // offset 2, length 64
	MiddleName                          string `json:"middleName,omitempty"`                          // offset 3, length 64
	LastName                            string `json:"lastName,omitempty"`                            // offset 4, length 128
	Suffix                              string `json:"suffix,omitempty"`                              // offset 5, length 20
	BirthDate                           string `json:"birthDate,omitempty"`                           // offset 6, length 10
	Gender                              string `json:"gender,omitempty"`                              // offset 7, length 1
	Culture                             string `json:"culture,omitempty"`                             // offset 8, length 50
	TaxId                               string `json:"taxId,omitempty"`                               // offset 9, length 30
	CustomerStatus                      string `json:"customerStatus,omitempty"`                      // offset 10, length 50
	CreatedDate                         string `json:"createdDate,omitempty"`                         // offset 11, length 34
	ArchivedDate                        string `json:"archivedDate,omitempty"`                        // offset 12, length 34
	DeceasedDate                        string `json:"deceasedDate,omitempty"`                        // offset 13, length 34
	IdVerificationDate                  string `json:"idVerificationDate,omitempty"`                  // offset 14, length 34
	IdVerificationDocumentsReceivedDate string `json:"idVerificationDocumentsReceivedDate,omitempty"` // offset 15, length 34
	DriversLicenseNumber                string `json:"driversLicenseNumber,omitempty"`                // offset 16, length 30
	DriversLicenseState                 string `json:"driversLicenseState,omitempty"`                 // offset 17, length 2
	DriversLicenseIssueDate             string `json:"driversLicenseIssueDate,omitempty"`             // offset 18, length 34
	DriversLicenseExpireDate            string `json:"driversLicenseExpireDate,omitempty"`            // offset 19, length 34
	PassportNumber                      string `json:"passportNumber,omitempty"`                      // offset 20, length 30
	PassportCountry                     string `json:"passportCountry,omitempty"`                     // offset 21, length 5
	PassportIssueDate                   string `json:"passportIssueDate,omitempty"`                   // offset 22, length 34
	PassportExpireDate                  string `json:"passportExpireDate,omitempty"`                  // offset 23, length 34
	EmailAddress                        string `json:"emailAddress,omitempty"`                        // offset 24, length 255
	IsSubjectToBackupWithholding        string `json:"isSubjectToBackupWithholding,omitempty"`        // offset 25, length 1
	IsOptedInToBankCommunication        string `json:"isOptedInToBankCommunication,omitempty"`        // offset 26, length 1
	IsDocumentsAccepted                 string `json:"isDocumentsAccepted,omitempty"`                 // offset 27, length 1
	DocumentsAcceptedDate               string `json:"documentsAcceptedDate,omitempty"`               // offset 28, length 34
	IsLocked                            string `json:"isLocked,omitempty"`                            // offset 29, length 1
	LockedDate                          string `json:"lockedDate,omitempty"`                          // offset 30, length 34
	LockedReason                        string `json:"lockedReason,omitempty"`                        // offset 31, length 255
	ResidenceLine1                      string `json:"residenceLine1,omitempty"`                      // offset 32, length 100
	ResidenceLine2                      string `json:"residenceLine2,omitempty"`                      // offset 33, length 100
	ResidenceLine3                      string `json:"residenceLine3,omitempty"`                      // offset 34, length 100
	ResidenceLine4                      string `json:"residenceLine4,omitempty"`                      // offset 35, length 100
	ResidenceCity                       string `json:"residenceCity,omitempty"`                       // offset 36, length 50
	ResidenceState                      string `json:"residenceState,omitempty"`                      // offset 37, length 2
	ResidencePostalCode                 string `json:"residencePostalCode,omitempty"`                 // offset 38, length 50
	ResidenceCountry                    string `json:"residenceCountry,omitempty"`                    // offset 39, length 50
	MailingLine1                        string `json:"mailingLine1,omitempty"`                        // offset 40, length 100
	MailingLine2                        string `json:"mailingLine2,omitempty"`                        // offset 41, length 100
	MailingLine3                        string `json:"mailingLine3,omitempty"`                        // offset 42, length 100
	MailingLine4                        string `json:"mailingLine4,omitempty"`                        // offset 43, length 100
	MailingCity                         string `json:"mailingCity,omitempty"`                         // offset 44, length 50
	MailingState                        string `json:"mailingState,omitempty"`                        // offset 45, length 2
	MailingPostalCode                   string `json:"mailingPostalCode,omitempty"`                   // offset 46, length 50
	MailingCountry                      string `json:"mailingCountry,omitempty"`                      // offset 47, length 50
	HomePhone                           string `json:"homePhone,omitempty"`                           // offset 48, length 50
	MobilePhone                         string `json:"mobilePhone,omitempty"`                         // offset 49, length 50
	OfficePhone                         string `json:"officePhone,omitempty"`                         // offset 50, length 50
	CustomField1                        string `json:"customField1,omitempty"`                        // offset 51, length 50
	CustomField2                        string `json:"customField2,omitempty"`                        // offset 52, length 50
	CustomField3                        string `json:"customField3,omitempty"`                        // offset 53, length 50
	CustomField4                        string `json:"customField4,omitempty"`                        // offset 54, length 50
	CustomField5                        string `json:"customField5,omitempty"`                        // offset 55, length 50
	LastActivityDate                    string `json:"lastActivityDate,omitempty"`                    // offset 56, length 34
	ArchivedReason                      string `json:"archivedReason,omitempty"`                      // offset 57, length 255
	AccountNumber                       string `json:"accountNumber,omitempty"`                       // offset 58, length 10

	Typed_CustomerId                   int64 `json:"-"`
	Typed_AccountId                    int64 `json:"-"`
	Typed_IsSubjectToBackupWithholding bool  `json:"-"`
	Typed_IsOptedInToBankCommunication bool  `json:"-"`
	Typed_IsDocumentsAccepted          bool  `json:"-"`
	Typed_IsLocked                     bool  `json:"-"`
}

type TypedCustomerArchivedFileModel struct {
	CustomerId                          int64  `json:"customerId,omitempty"`                          // offset 0, length 10
	AccountId                           int64  `json:"accountId,omitempty"`                           // offset 1, length 10
	FirstName                           string `json:"firstName,omitempty"`                           // offset 2, length 64
	MiddleName                          string `json:"middleName,omitempty"`                          // offset 3, length 64
	LastName                            string `json:"lastName,omitempty"`                            // offset 4, length 128
	Suffix                              string `json:"suffix,omitempty"`                              // offset 5, length 20
	BirthDate                           string `json:"birthDate,omitempty"`                           // offset 6, length 10
	Gender                              string `json:"gender,omitempty"`                              // offset 7, length 1
	Culture                             string `json:"culture,omitempty"`                             // offset 8, length 50
	TaxId                               string `json:"taxId,omitempty"`                               // offset 9, length 30
	CustomerStatus                      string `json:"customerStatus,omitempty"`                      // offset 10, length 50
	CreatedDate                         string `json:"createdDate,omitempty"`                         // offset 11, length 34
	ArchivedDate                        string `json:"archivedDate,omitempty"`                        // offset 12, length 34
	DeceasedDate                        string `json:"deceasedDate,omitempty"`                        // offset 13, length 34
	IdVerificationDate                  string `json:"idVerificationDate,omitempty"`                  // offset 14, length 34
	IdVerificationDocumentsReceivedDate string `json:"idVerificationDocumentsReceivedDate,omitempty"` // offset 15, length 34
	DriversLicenseNumber                string `json:"driversLicenseNumber,omitempty"`                // offset 16, length 30
	DriversLicenseState                 string `json:"driversLicenseState,omitempty"`                 // offset 17, length 2
	DriversLicenseIssueDate             string `json:"driversLicenseIssueDate,omitempty"`             // offset 18, length 34
	DriversLicenseExpireDate            string `json:"driversLicenseExpireDate,omitempty"`            // offset 19, length 34
	PassportNumber                      string `json:"passportNumber,omitempty"`                      // offset 20, length 30
	PassportCountry                     string `json:"passportCountry,omitempty"`                     // offset 21, length 5
	PassportIssueDate                   string `json:"passportIssueDate,omitempty"`                   // offset 22, length 34
	PassportExpireDate                  string `json:"passportExpireDate,omitempty"`                  // offset 23, length 34
	EmailAddress                        string `json:"emailAddress,omitempty"`                        // offset 24, length 255
	IsSubjectToBackupWithholding        bool   `json:"isSubjectToBackupWithholding,omitempty"`        // offset 25, length 1
	IsOptedInToBankCommunication        bool   `json:"isOptedInToBankCommunication,omitempty"`        // offset 26, length 1
	IsDocumentsAccepted                 bool   `json:"isDocumentsAccepted,omitempty"`                 // offset 27, length 1
	DocumentsAcceptedDate               string `json:"documentsAcceptedDate,omitempty"`               // offset 28, length 34
	IsLocked                            bool   `json:"isLocked,omitempty"`                            // offset 29, length 1
	LockedDate                          string `json:"lockedDate,omitempty"`                          // offset 30, length 34
	LockedReason                        string `json:"lockedReason,omitempty"`                        // offset 31, length 255
	ResidenceLine1                      string `json:"residenceLine1,omitempty"`                      // offset 32, length 100
	ResidenceLine2                      string `json:"residenceLine2,omitempty"`                      // offset 33, length 100
	ResidenceLine3                      string `json:"residenceLine3,omitempty"`                      // offset 34, length 100
	ResidenceLine4                      string `json:"residenceLine4,omitempty"`                      // offset 35, length 100
	ResidenceCity                       string `json:"residenceCity,omitempty"`                       // offset 36, length 50
	ResidenceState                      string `json:"residenceState,omitempty"`                      // offset 37, length 2
	ResidencePostalCode                 string `json:"residencePostalCode,omitempty"`                 // offset 38, length 50
	ResidenceCountry                    string `json:"residenceCountry,omitempty"`                    // offset 39, length 50
	MailingLine1                        string `json:"mailingLine1,omitempty"`                        // offset 40, length 100
	MailingLine2                        string `json:"mailingLine2,omitempty"`                        // offset 41, length 100
	MailingLine3                        string `json:"mailingLine3,omitempty"`                        // offset 42, length 100
	MailingLine4                        string `json:"mailingLine4,omitempty"`                        // offset 43, length 100
	MailingCity                         string `json:"mailingCity,omitempty"`                         // offset 44, length 50
	MailingState                        string `json:"mailingState,omitempty"`                        // offset 45, length 2
	MailingPostalCode                   string `json:"mailingPostalCode,omitempty"`                   // offset 46, length 50
	MailingCountry                      string `json:"mailingCountry,omitempty"`                      // offset 47, length 50
	HomePhone                           string `json:"homePhone,omitempty"`                           // offset 48, length 50
	MobilePhone                         string `json:"mobilePhone,omitempty"`                         // offset 49, length 50
	OfficePhone                         string `json:"officePhone,omitempty"`                         // offset 50, length 50
	CustomField1                        string `json:"customField1,omitempty"`                        // offset 51, length 50
	CustomField2                        string `json:"customField2,omitempty"`                        // offset 52, length 50
	CustomField3                        string `json:"customField3,omitempty"`                        // offset 53, length 50
	CustomField4                        string `json:"customField4,omitempty"`                        // offset 54, length 50
	CustomField5                        string `json:"customField5,omitempty"`                        // offset 55, length 50
	LastActivityDate                    string `json:"lastActivityDate,omitempty"`                    // offset 56, length 34
	ArchivedReason                      string `json:"archivedReason,omitempty"`                      // offset 57, length 255
	AccountNumber                       string `json:"accountNumber,omitempty"`                       // offset 58, length 10

}

func (mdl *CustomerArchivedFileModel) IsTabDelimitedFile() bool {
	return true
}

func (mdl *CustomerArchivedFileModel) HeaderFieldCount() int {
	return 4
}

func (mdl *CustomerArchivedFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return &TypedCustomerArchivedFileModel{
			CustomerId:                          mdl.Typed_CustomerId,
			AccountId:                           mdl.Typed_AccountId,
			FirstName:                           mdl.FirstName,
			MiddleName:                          mdl.MiddleName,
			LastName:                            mdl.LastName,
			Suffix:                              mdl.Suffix,
			BirthDate:                           mdl.BirthDate,
			Gender:                              mdl.Gender,
			Culture:                             mdl.Culture,
			TaxId:                               mdl.TaxId,
			CustomerStatus:                      mdl.CustomerStatus,
			CreatedDate:                         mdl.CreatedDate,
			ArchivedDate:                        mdl.ArchivedDate,
			DeceasedDate:                        mdl.DeceasedDate,
			IdVerificationDate:                  mdl.IdVerificationDate,
			IdVerificationDocumentsReceivedDate: mdl.IdVerificationDocumentsReceivedDate,
			DriversLicenseNumber:                mdl.DriversLicenseNumber,
			DriversLicenseState:                 mdl.DriversLicenseState,
			DriversLicenseIssueDate:             mdl.DriversLicenseIssueDate,
			DriversLicenseExpireDate:            mdl.DriversLicenseExpireDate,
			PassportNumber:                      mdl.PassportNumber,
			PassportCountry:                     mdl.PassportCountry,
			PassportIssueDate:                   mdl.PassportIssueDate,
			PassportExpireDate:                  mdl.PassportExpireDate,
			EmailAddress:                        mdl.EmailAddress,
			IsSubjectToBackupWithholding:        mdl.Typed_IsSubjectToBackupWithholding,
			IsOptedInToBankCommunication:        mdl.Typed_IsOptedInToBankCommunication,
			IsDocumentsAccepted:                 mdl.Typed_IsDocumentsAccepted,
			DocumentsAcceptedDate:               mdl.DocumentsAcceptedDate,
			IsLocked:                            mdl.Typed_IsLocked,
			LockedDate:                          mdl.LockedDate,
			LockedReason:                        mdl.LockedReason,
			ResidenceLine1:                      mdl.ResidenceLine1,
			ResidenceLine2:                      mdl.ResidenceLine2,
			ResidenceLine3:                      mdl.ResidenceLine3,
			ResidenceLine4:                      mdl.ResidenceLine4,
			ResidenceCity:                       mdl.ResidenceCity,
			ResidenceState:                      mdl.ResidenceState,
			ResidencePostalCode:                 mdl.ResidencePostalCode,
			ResidenceCountry:                    mdl.ResidenceCountry,
			MailingLine1:                        mdl.MailingLine1,
			MailingLine2:                        mdl.MailingLine2,
			MailingLine3:                        mdl.MailingLine3,
			MailingLine4:                        mdl.MailingLine4,
			MailingCity:                         mdl.MailingCity,
			MailingState:                        mdl.MailingState,
			MailingPostalCode:                   mdl.MailingPostalCode,
			MailingCountry:                      mdl.MailingCountry,
			HomePhone:                           mdl.HomePhone,
			MobilePhone:                         mdl.MobilePhone,
			OfficePhone:                         mdl.OfficePhone,
			CustomField1:                        mdl.CustomField1,
			CustomField2:                        mdl.CustomField2,
			CustomField3:                        mdl.CustomField3,
			CustomField4:                        mdl.CustomField4,
			CustomField5:                        mdl.CustomField5,
			LastActivityDate:                    mdl.LastActivityDate,
			ArchivedReason:                      mdl.ArchivedReason,
			AccountNumber:                       mdl.AccountNumber,
		}
	}
}

func (mdl *CustomerArchivedFileModel) GetValues() []string {
	return []string{
		mdl.CustomerId,
		mdl.AccountId,
		mdl.FirstName,
		mdl.MiddleName,
		mdl.LastName,
		mdl.Suffix,
		mdl.BirthDate,
		mdl.Gender,
		mdl.Culture,
		mdl.TaxId,
		mdl.CustomerStatus,
		mdl.CreatedDate,
		mdl.ArchivedDate,
		mdl.DeceasedDate,
		mdl.IdVerificationDate,
		mdl.IdVerificationDocumentsReceivedDate,
		mdl.DriversLicenseNumber,
		mdl.DriversLicenseState,
		mdl.DriversLicenseIssueDate,
		mdl.DriversLicenseExpireDate,
		mdl.PassportNumber,
		mdl.PassportCountry,
		mdl.PassportIssueDate,
		mdl.PassportExpireDate,
		mdl.EmailAddress,
		mdl.IsSubjectToBackupWithholding,
		mdl.IsOptedInToBankCommunication,
		mdl.IsDocumentsAccepted,
		mdl.DocumentsAcceptedDate,
		mdl.IsLocked,
		mdl.LockedDate,
		mdl.LockedReason,
		mdl.ResidenceLine1,
		mdl.ResidenceLine2,
		mdl.ResidenceLine3,
		mdl.ResidenceLine4,
		mdl.ResidenceCity,
		mdl.ResidenceState,
		mdl.ResidencePostalCode,
		mdl.ResidenceCountry,
		mdl.MailingLine1,
		mdl.MailingLine2,
		mdl.MailingLine3,
		mdl.MailingLine4,
		mdl.MailingCity,
		mdl.MailingState,
		mdl.MailingPostalCode,
		mdl.MailingCountry,
		mdl.HomePhone,
		mdl.MobilePhone,
		mdl.OfficePhone,
		mdl.CustomField1,
		mdl.CustomField2,
		mdl.CustomField3,
		mdl.CustomField4,
		mdl.CustomField5,
		mdl.LastActivityDate,
		mdl.ArchivedReason,
		mdl.AccountNumber,
	}
}

func (mdl *CustomerArchivedFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	// special: this is a tab-delimited file
	splits := strings.Split(line, "\t")

	const expectedFieldCount = 59
	actualFieldCount := len(splits)

	if strict && actualFieldCount != expectedFieldCount {
		err = fmt.Errorf("expected exactly %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else if !strict && actualFieldCount < expectedFieldCount {
		err = fmt.Errorf("expected at least %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else {
		_, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, line, pos, 10, preserve, err)
		_, err = ConsumeIntField(&mdl.AccountId, &mdl.Typed_AccountId, line, pos, 10, preserve, err)
		_, err = ConsumeField(&mdl.FirstName, line, pos, 64, preserve, err)
		_, err = ConsumeField(&mdl.MiddleName, line, pos, 64, preserve, err)
		_, err = ConsumeField(&mdl.LastName, line, pos, 128, preserve, err)
		_, err = ConsumeField(&mdl.Suffix, line, pos, 20, preserve, err)
		_, err = ConsumeField(&mdl.BirthDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.Gender, line, pos, 1, preserve, err)
		_, err = ConsumeField(&mdl.Culture, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.TaxId, line, pos, 30, preserve, err)
		_, err = ConsumeField(&mdl.CustomerStatus, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.CreatedDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.ArchivedDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.DeceasedDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.IdVerificationDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.IdVerificationDocumentsReceivedDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.DriversLicenseNumber, line, pos, 30, preserve, err)
		_, err = ConsumeField(&mdl.DriversLicenseState, line, pos, 2, preserve, err)
		_, err = ConsumeField(&mdl.DriversLicenseIssueDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.DriversLicenseExpireDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.PassportNumber, line, pos, 30, preserve, err)
		_, err = ConsumeField(&mdl.PassportCountry, line, pos, 5, preserve, err)
		_, err = ConsumeField(&mdl.PassportIssueDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.PassportExpireDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.EmailAddress, line, pos, 255, preserve, err)
		_, err = ConsumeBoolField(&mdl.IsSubjectToBackupWithholding, &mdl.Typed_IsSubjectToBackupWithholding, line, pos, 1, "Y", preserve, err)
		_, err = ConsumeBoolField(&mdl.IsOptedInToBankCommunication, &mdl.Typed_IsOptedInToBankCommunication, line, pos, 1, "Y", preserve, err)
		_, err = ConsumeBoolField(&mdl.IsDocumentsAccepted, &mdl.Typed_IsDocumentsAccepted, line, pos, 1, "Y", preserve, err)
		_, err = ConsumeField(&mdl.DocumentsAcceptedDate, line, pos, 34, preserve, err)
		_, err = ConsumeBoolField(&mdl.IsLocked, &mdl.Typed_IsLocked, line, pos, 1, "Y", preserve, err)
		_, err = ConsumeField(&mdl.LockedDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.LockedReason, line, pos, 255, preserve, err)
		_, err = ConsumeField(&mdl.ResidenceLine1, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.ResidenceLine2, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.ResidenceLine3, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.ResidenceLine4, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.ResidenceCity, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.ResidenceState, line, pos, 2, preserve, err)
		_, err = ConsumeField(&mdl.ResidencePostalCode, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.ResidenceCountry, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.MailingLine1, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.MailingLine2, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.MailingLine3, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.MailingLine4, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.MailingCity, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.MailingState, line, pos, 2, preserve, err)
		_, err = ConsumeField(&mdl.MailingPostalCode, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.MailingCountry, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.HomePhone, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.MobilePhone, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.OfficePhone, line, pos, 50, preserve, err)
		_, err = ConsumeField(&mdl.CustomField1, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.CustomField2, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.CustomField3, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.CustomField4, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.CustomField5, line, pos, 100, preserve, err)
		_, err = ConsumeField(&mdl.LastActivityDate, line, pos, 34, preserve, err)
		_, err = ConsumeField(&mdl.ArchivedReason, line, pos, 255, preserve, err)
		_, err = ConsumeField(&mdl.AccountNumber, line, pos, 10, preserve, err)

	}

	return err
}
