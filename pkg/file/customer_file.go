package file

import (
	"fmt"
	"strings"
)

type CustomerFileModel struct {
	CustomerId                          string `json:"customerId,omitempty"`                          // offset 0, length 10
	CustomerTag                         string `json:"customerTag,omitempty"`                         // offset 1, length 50
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

	Typed_CustomerId                   int64 `json:"-"`
	Typed_IsSubjectToBackupWithholding bool  `json:"-"`
	Typed_IsOptedInToBankCommunication bool  `json:"-"`
	Typed_IsDocumentsAccepted          bool  `json:"-"`
	Typed_IsLocked                     bool  `json:"-"`
}

type TypedCustomerFileModel struct {
	CustomerId                          int64  `json:"customerId,omitempty"`                          // offset 0, length 10
	CustomerTag                         string `json:"customerTag,omitempty"`                         // offset 1, length 50
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

}

func (mdl *CustomerFileModel) IsTabDelimitedFile() bool {
	return true
}

func (mdl *CustomerFileModel) HeaderFieldCount() int {
	return 4
}

func (mdl *CustomerFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return &TypedCustomerFileModel{
			CustomerId:                          mdl.Typed_CustomerId,
			CustomerTag:                         mdl.CustomerTag,
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
		}
	}
}

func (mdl *CustomerFileModel) GetValues() []string {
	return []string{
		mdl.CustomerId,
		mdl.CustomerTag,
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
	}
}

func (mdl *CustomerFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	// special: this is a tab-delimited file
	splits := strings.Split(line, "\t")

	const expectedFieldCount = 57
	actualFieldCount := len(splits)

	if strict && actualFieldCount != expectedFieldCount {
		err = fmt.Errorf("expected exactly %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else if !strict && actualFieldCount < expectedFieldCount {
		err = fmt.Errorf("expected at least %v fields, got %v", expectedFieldCount, actualFieldCount)
	} else {

		_, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, splits[0], pos, len(splits[0]), preserve, err)
		_, err = ConsumeField(&mdl.CustomerTag, splits[1], pos, len(splits[1]), preserve, err)
		_, err = ConsumeField(&mdl.FirstName, splits[2], pos, len(splits[2]), preserve, err)
		_, err = ConsumeField(&mdl.MiddleName, splits[3], pos, len(splits[3]), preserve, err)
		_, err = ConsumeField(&mdl.LastName, splits[4], pos, len(splits[4]), preserve, err)
		_, err = ConsumeField(&mdl.Suffix, splits[5], pos, len(splits[5]), preserve, err)
		_, err = ConsumeField(&mdl.BirthDate, splits[6], pos, len(splits[6]), preserve, err)
		_, err = ConsumeField(&mdl.Gender, splits[7], pos, len(splits[7]), preserve, err)
		_, err = ConsumeField(&mdl.Culture, splits[8], pos, len(splits[8]), preserve, err)
		_, err = ConsumeField(&mdl.TaxId, splits[9], pos, len(splits[9]), preserve, err)
		_, err = ConsumeField(&mdl.CustomerStatus, splits[10], pos, len(splits[10]), preserve, err)
		_, err = ConsumeField(&mdl.CreatedDate, splits[11], pos, len(splits[11]), preserve, err)
		_, err = ConsumeField(&mdl.ArchivedDate, splits[12], pos, len(splits[12]), preserve, err)
		_, err = ConsumeField(&mdl.DeceasedDate, splits[13], pos, len(splits[13]), preserve, err)
		_, err = ConsumeField(&mdl.IdVerificationDate, splits[14], pos, len(splits[14]), preserve, err)
		_, err = ConsumeField(&mdl.IdVerificationDocumentsReceivedDate, splits[15], pos, len(splits[15]), preserve, err)
		_, err = ConsumeField(&mdl.DriversLicenseNumber, splits[16], pos, len(splits[16]), preserve, err)
		_, err = ConsumeField(&mdl.DriversLicenseState, splits[17], pos, len(splits[17]), preserve, err)
		_, err = ConsumeField(&mdl.DriversLicenseIssueDate, splits[18], pos, len(splits[18]), preserve, err)
		_, err = ConsumeField(&mdl.DriversLicenseExpireDate, splits[19], pos, len(splits[19]), preserve, err)
		_, err = ConsumeField(&mdl.PassportNumber, splits[20], pos, len(splits[20]), preserve, err)
		_, err = ConsumeField(&mdl.PassportCountry, splits[21], pos, len(splits[21]), preserve, err)
		_, err = ConsumeField(&mdl.PassportIssueDate, splits[22], pos, len(splits[22]), preserve, err)
		_, err = ConsumeField(&mdl.PassportExpireDate, splits[23], pos, len(splits[23]), preserve, err)
		_, err = ConsumeField(&mdl.EmailAddress, splits[24], pos, len(splits[24]), preserve, err)
		_, err = ConsumeBoolField(&mdl.IsSubjectToBackupWithholding, &mdl.Typed_IsSubjectToBackupWithholding, splits[25], pos, len(splits[25]), "Y", preserve, err)
		_, err = ConsumeBoolField(&mdl.IsOptedInToBankCommunication, &mdl.Typed_IsOptedInToBankCommunication, splits[26], pos, len(splits[26]), "Y", preserve, err)
		_, err = ConsumeBoolField(&mdl.IsDocumentsAccepted, &mdl.Typed_IsDocumentsAccepted, splits[27], pos, len(splits[27]), "Y", preserve, err)
		_, err = ConsumeField(&mdl.DocumentsAcceptedDate, splits[28], pos, len(splits[28]), preserve, err)
		_, err = ConsumeBoolField(&mdl.IsLocked, &mdl.Typed_IsLocked, splits[29], pos, len(splits[29]), "Y", preserve, err)
		_, err = ConsumeField(&mdl.LockedDate, splits[30], pos, len(splits[30]), preserve, err)
		_, err = ConsumeField(&mdl.LockedReason, splits[31], pos, len(splits[31]), preserve, err)
		_, err = ConsumeField(&mdl.ResidenceLine1, splits[32], pos, len(splits[32]), preserve, err)
		_, err = ConsumeField(&mdl.ResidenceLine2, splits[33], pos, len(splits[33]), preserve, err)
		_, err = ConsumeField(&mdl.ResidenceLine3, splits[34], pos, len(splits[34]), preserve, err)
		_, err = ConsumeField(&mdl.ResidenceLine4, splits[35], pos, len(splits[35]), preserve, err)
		_, err = ConsumeField(&mdl.ResidenceCity, splits[36], pos, len(splits[36]), preserve, err)
		_, err = ConsumeField(&mdl.ResidenceState, splits[37], pos, len(splits[37]), preserve, err)
		_, err = ConsumeField(&mdl.ResidencePostalCode, splits[38], pos, len(splits[38]), preserve, err)
		_, err = ConsumeField(&mdl.ResidenceCountry, splits[39], pos, len(splits[39]), preserve, err)
		_, err = ConsumeField(&mdl.MailingLine1, splits[40], pos, len(splits[40]), preserve, err)
		_, err = ConsumeField(&mdl.MailingLine2, splits[41], pos, len(splits[41]), preserve, err)
		_, err = ConsumeField(&mdl.MailingLine3, splits[42], pos, len(splits[42]), preserve, err)
		_, err = ConsumeField(&mdl.MailingLine4, splits[43], pos, len(splits[43]), preserve, err)
		_, err = ConsumeField(&mdl.MailingCity, splits[44], pos, len(splits[44]), preserve, err)
		_, err = ConsumeField(&mdl.MailingState, splits[45], pos, len(splits[45]), preserve, err)
		_, err = ConsumeField(&mdl.MailingPostalCode, splits[46], pos, len(splits[46]), preserve, err)
		_, err = ConsumeField(&mdl.MailingCountry, splits[47], pos, len(splits[47]), preserve, err)
		_, err = ConsumeField(&mdl.HomePhone, splits[48], pos, len(splits[48]), preserve, err)
		_, err = ConsumeField(&mdl.MobilePhone, splits[49], pos, len(splits[49]), preserve, err)
		_, err = ConsumeField(&mdl.OfficePhone, splits[50], pos, len(splits[50]), preserve, err)
		_, err = ConsumeField(&mdl.CustomField1, splits[51], pos, len(splits[51]), preserve, err)
		_, err = ConsumeField(&mdl.CustomField2, splits[52], pos, len(splits[52]), preserve, err)
		_, err = ConsumeField(&mdl.CustomField3, splits[53], pos, len(splits[53]), preserve, err)
		_, err = ConsumeField(&mdl.CustomField4, splits[54], pos, len(splits[54]), preserve, err)
		_, err = ConsumeField(&mdl.CustomField5, splits[55], pos, len(splits[55]), preserve, err)
		_, err = ConsumeField(&mdl.LastActivityDate, splits[56], pos, len(splits[56]), preserve, err)
	}

	return err
}
