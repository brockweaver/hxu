package file

import (
	"fmt"
)

type CardEventNotificationFileModel struct {
	UserEventId                                                           string `json:"userEventId,omitempty"`                                                           // length 19
	CustomerId                                                            string `json:"customerId,omitempty"`                                                            // length 10
	CardId                                                                string `json:"cardId,omitempty"`                                                                // length 10
	TransactionId                                                         string `json:"transactionId,omitempty"`                                                         // length 19
	AuthorizationTransactionId                                            string `json:"authorizationTransactionId,omitempty"`                                            // length 19
	RequestTypeCode                                                       string `json:"requestTypeCode,omitempty"`                                                       // length 3
	Reserved                                                              string `json:"reserved,omitempty"`                                                              // length 12
	MerchantCategoryCode                                                  string `json:"merchantCategoryCode,omitempty"`                                                  // length 4
	MerchantGroupTypeCode                                                 string `json:"merchantGroupTypeCode,omitempty"`                                                 // length 6
	CashbackAmount                                                        string `json:"cashbackAmount,omitempty"`                                                        // length 10
	SurchargeAmount                                                       string `json:"surchargeAmount,omitempty"`                                                       // length 10
	CashDepositAmount                                                     string `json:"cashDepositAmount,omitempty"`                                                     // length 10
	CheckDepositAmount                                                    string `json:"checkDepositAmount,omitempty"`                                                    // length 10
	TerminalId                                                            string `json:"terminalId,omitempty"`                                                            // length 15
	MerchantId                                                            string `json:"merchantId,omitempty"`                                                            // length 15
	MerchantLocation                                                      string `json:"merchantLocation,omitempty"`                                                      // length 25
	MerchantCity                                                          string `json:"merchantCity,omitempty"`                                                          // length 13
	MerchantStateCode                                                     string `json:"merchantStateCode,omitempty"`                                                     // length 2
	MerchantZipCode                                                       string `json:"merchantZipCode,omitempty"`                                                       // length 9
	MerchantCountryCode                                                   string `json:"merchantCountryCode,omitempty"`                                                   // length 2
	PanEntryModeTypeCode                                                  string `json:"panEntryModeTypeCode,omitempty"`                                                  // length 10
	PinEntryModeTypeCode                                                  string `json:"pinEntryModeTypeCode,omitempty"`                                                  // length 10
	CardHolderPresenceTypeCode                                            string `json:"cardHolderPresenceTypeCode,omitempty"`                                            // length 10
	CardPresenceTypeCode                                                  string `json:"cardPresenceTypeCode,omitempty"`                                                  // length 10
	TerminalUnattendedTypeCode                                            string `json:"terminalUnattendedTypeCode,omitempty"`                                            // length 10
	TerminalPremisesTypeCode                                              string `json:"terminalPremisesTypeCode,omitempty"`                                              // length 10
	CustomerTag                                                           string `json:"customerTag,omitempty"`                                                           // length 50
	CardTag                                                               string `json:"cardTag,omitempty"`                                                               // length 50
	Amount                                                                string `json:"amount,omitempty"`                                                                // length 10
	AvailableDate                                                         string `json:"availableDate,omitempty"`                                                         // length 34
	CreatedDate                                                           string `json:"createdDate,omitempty"`                                                           // length 34
	CustomField1                                                          string `json:"customField1,omitempty"`                                                          // length 50
	Description                                                           string `json:"description,omitempty"`                                                           // length 255
	DenialReason                                                          string `json:"denialReason,omitempty"`                                                          // length 255
	FeeCode                                                               string `json:"feeCode,omitempty"`                                                               // length 3
	FeeDescription                                                        string `json:"feeDescription,omitempty"`                                                        // length 50
	FriendlyDescription                                                   string `json:"friendlyDescription,omitempty"`                                                   // length 255
	FromAccountAmount                                                     string `json:"fromAccountAmount,omitempty"`                                                     // length 10
	FromAccountId                                                         string `json:"fromAccountId,omitempty"`                                                         // length 10
	FromAccountNumberMasked                                               string `json:"fromAccountNumberMasked,omitempty"`                                               // length 50
	FromAvailableAmount                                                   string `json:"fromAvailableAmount,omitempty"`                                                   // length 10
	FromCategory                                                          string `json:"fromCategory,omitempty"`                                                          // length 50
	FromCreatedDate                                                       string `json:"fromCreatedDate,omitempty"`                                                       // length 34
	FromCustomField1                                                      string `json:"fromCustomField1,omitempty"`                                                      // length 50
	FromCustomField2                                                      string `json:"fromCustomField2,omitempty"`                                                      // length 50
	FromCustomField3                                                      string `json:"fromCustomField3,omitempty"`                                                      // length 50
	FromCustomField4                                                      string `json:"fromCustomField4,omitempty"`                                                      // length 50
	FromCustomField5                                                      string `json:"fromCustomField5,omitempty"`                                                      // length 50
	FromLegalName1                                                        string `json:"fromLegalName1,omitempty"`                                                        // length 100
	FromLegalName2                                                        string `json:"fromLegalName2,omitempty"`                                                        // length 100
	FromName                                                              string `json:"fromName,omitempty"`                                                              // length 50
	FromPendingAmount                                                     string `json:"fromPendingAmount,omitempty"`                                                     // length 10
	FromPrimaryCustomerId                                                 string `json:"fromPrimaryCustomerId,omitempty"`                                                 // length 10
	FromSubCategory                                                       string `json:"fromSubCategory,omitempty"`                                                       // length 50
	FromTag                                                               string `json:"fromTag,omitempty"`                                                               // length 50
	InstitutionName                                                       string `json:"institutionName,omitempty"`                                                       // length 50
	MasterId                                                              string `json:"masterId,omitempty"`                                                              // length 19
	ModifiedById                                                          string `json:"modifiedById,omitempty"`                                                          // length 10
	PayloadTypeId                                                         string `json:"payloadTypeId,omitempty"`                                                         // length 10
	SettledDate                                                           string `json:"settledDate,omitempty"`                                                           // length 34
	SubType                                                               string `json:"subType,omitempty"`                                                               // length 255
	SubTypeCode                                                           string `json:"subTypeCode,omitempty"`                                                           // length 6
	Tag                                                                   string `json:"tag,omitempty"`                                                                   // length 50
	ToAccountAmount                                                       string `json:"toAccountAmount,omitempty"`                                                       // length 10
	ToAccountId                                                           string `json:"toAccountId,omitempty"`                                                           // length 10
	ToAccountNumberMasked                                                 string `json:"toAccountNumberMasked,omitempty"`                                                 // length 50
	ToAvailableAmount                                                     string `json:"toAvailableAmount,omitempty"`                                                     // length 10
	ToCategory                                                            string `json:"toCategory,omitempty"`                                                            // length 50
	ToCreatedDate                                                         string `json:"toCreatedDate,omitempty"`                                                         // length 34
	ToCustomField1                                                        string `json:"toCustomField1,omitempty"`                                                        // length 50
	ToCustomField2                                                        string `json:"toCustomField2,omitempty"`                                                        // length 50
	ToCustomField3                                                        string `json:"toCustomField3,omitempty"`                                                        // length 50
	ToCustomField4                                                        string `json:"toCustomField4,omitempty"`                                                        // length 50
	ToCustomField5                                                        string `json:"toCustomField5,omitempty"`                                                        // length 50
	ToLegalName1                                                          string `json:"toLegalName1,omitempty"`                                                          // length 100
	ToLegalName2                                                          string `json:"toLegalName2,omitempty"`                                                          // length 100
	ToName                                                                string `json:"toName,omitempty"`                                                                // length 50
	ToPendingAmount                                                       string `json:"toPendingAmount,omitempty"`                                                       // length 10
	ToPrimaryCustomerId                                                   string `json:"toPrimaryCustomerId,omitempty"`                                                   // length 10
	ToSubCategory                                                         string `json:"toSubCategory,omitempty"`                                                         // length 50
	ToTag                                                                 string `json:"toTag,omitempty"`                                                                 // length 50
	Type                                                                  string `json:"type,omitempty"`                                                                  // length 50
	TypeCode                                                              string `json:"typeCode,omitempty"`                                                              // length 6
	EventTypeId                                                           string `json:"eventTypeId,omitempty"`                                                           // length 10
	NetworkProviderTypeId                                                 string `json:"networkProviderTypeId,omitempty"`                                                 // length 1
	PointOfServicePanEntryMode                                            string `json:"pointOfServicePanEntryMode,omitempty"`                                            // length 2
	PointOfServicePinEntryMode                                            string `json:"pointOfServicePinEntryMode,omitempty"`                                            // length 1
	Cvv2PresenceIndicator                                                 string `json:"cvv2PresenceIndicator,omitempty"`                                                 // length 1
	Cvv2Result                                                            string `json:"cvv2Result,omitempty"`                                                            // length 1
	Token                                                                 string `json:"token,omitempty"`                                                                 // length 19
	TokenAssuranceLevel                                                   string `json:"tokenAssuranceLevel,omitempty"`                                                   // length 2
	DigitalWalletTokenRequestorTypeId                                     string `json:"digitalWalletTokenRequestorTypeId,omitempty"`                                     // length 2
	TokenExpirationDate                                                   string `json:"tokenExpirationDate,omitempty"`                                                   // length 4
	PaymentAccountReferenceNumber                                         string `json:"paymentAccountReferenceNumber,omitempty"`                                         // length 29
	MessageTypeIndicator                                                  string `json:"messageTypeIndicator,omitempty"`                                                  // length 4
	OutputMessageTypeIndicator                                            string `json:"outputMessageTypeIndicator,omitempty"`                                            // length 4
	OutputResponseCode                                                    string `json:"outputResponseCode,omitempty"`                                                    // length 2
	SystemTraceAuditNumber                                                string `json:"systemTraceAuditNumber,omitempty"`                                                // length 6
	AcquirerInstitutionCountryCode                                        string `json:"acquirerInstitutionCountryCode,omitempty"`                                        // length 3
	AuthorizationIdentificationResponse                                   string `json:"authorizationIdentificationResponse,omitempty"`                                   // length 12
	ResponseCode                                                          string `json:"responseCode,omitempty"`                                                          // length 2
	PinValidationCode                                                     string `json:"pinValidationCode,omitempty"`                                                     // length 1
	AdditionalAmounts_Purchase                                            string `json:"additionalAmounts_Purchase,omitempty"`                                            // length 10
	AdditionalAmounts_Gratuity                                            string `json:"additionalAmounts_Gratuity,omitempty"`                                            // length 10
	Advice_OriginatorCode                                                 string `json:"advice_OriginatorCode,omitempty"`                                                 // length 1
	Advice_ReasonCode                                                     string `json:"advice_ReasonCode,omitempty"`                                                     // length 1
	PrivatelyDefinedData_IssuerNetworkIdCode                              string `json:"privatelyDefinedData_IssuerNetworkIdCode,omitempty"`                              // length 3
	PrivatelyDefinedData_AdditionalTransactionElement_FallbackIndicator   string `json:"privatelyDefinedData_AdditionalTransactionElement_FallbackIndicator,omitempty"`   // length 1
	PrivatelyDefinedData_ProcessingFlag_SpecialTransactionIndicator       string `json:"privatelyDefinedData_ProcessingFlag_SpecialTransactionIndicator,omitempty"`       // length 1
	PrivatelyDefinedData_ProcessingFlag_ISAIndicator                      string `json:"privatelyDefinedData_ProcessingFlag_ISAIndicator,omitempty"`                      // length 1
	PrivatelyDefinedData_ProcessingFlag_PartialAuthIndicator              string `json:"privatelyDefinedData_ProcessingFlag_PartialAuthIndicator,omitempty"`              // length 1
	PrivatelyDefinedData_RiskData_ScoreSource                             string `json:"privatelyDefinedData_RiskData_ScoreSource,omitempty"`                             // length 1
	PrivatelyDefinedData_RiskData_ScoreValue                              string `json:"privatelyDefinedData_RiskData_ScoreValue,omitempty"`                              // length 4
	PrivatelyDefinedData_RiskData_ResponseCode                            string `json:"privatelyDefinedData_RiskData_ResponseCode,omitempty"`                            // length 1
	PrivatelyDefinedData_RiskData_FalconReason1                           string `json:"privatelyDefinedData_RiskData_FalconReason1,omitempty"`                           // length 2
	PrivatelyDefinedData_RiskData_FalconReason2                           string `json:"privatelyDefinedData_RiskData_FalconReason2,omitempty"`                           // length 2
	PrivatelyDefinedData_RiskData_FalconReason3                           string `json:"privatelyDefinedData_RiskData_FalconReason3,omitempty"`                           // length 2
	PrivatelyDefinedData_RiskData_VisaRiskScore                           string `json:"privatelyDefinedData_RiskData_VisaRiskScore,omitempty"`                           // length 2
	PrivatelyDefinedData_RiskData_VisaRiskReason                          string `json:"privatelyDefinedData_RiskData_VisaRiskReason,omitempty"`                          // length 2
	PrivatelyDefinedData_RiskData_VisaRiskConditionCode1                  string `json:"privatelyDefinedData_RiskData_VisaRiskConditionCode1,omitempty"`                  // length 2
	PrivatelyDefinedData_RiskData_VisaRiskConditionCode2                  string `json:"privatelyDefinedData_RiskData_VisaRiskConditionCode2,omitempty"`                  // length 2
	PrivatelyDefinedData_RiskData_VisaRiskConditionCode3                  string `json:"privatelyDefinedData_RiskData_VisaRiskConditionCode3,omitempty"`                  // length 2
	PrivatelyDefinedData_RiskData_VAAConditionCode1Rank                   string `json:"privatelyDefinedData_RiskData_VAAConditionCode1Rank,omitempty"`                   // length 2
	PrivatelyDefinedData_RiskData_RTDResultCode                           string `json:"privatelyDefinedData_RiskData_RTDResultCode,omitempty"`                           // length 1
	PrivatelyDefinedData_RiskData_TravelStatusIndicator                   string `json:"privatelyDefinedData_RiskData_TravelStatusIndicator,omitempty"`                   // length 1
	TextInfo                                                              string `json:"textInfo,omitempty"`                                                              // length 255
	Track2Data_ServiceCode                                                string `json:"track2Data_ServiceCode,omitempty"`                                                // length 3
	RetrievalReferenceNumber                                              string `json:"retrievalReferenceNumber,omitempty"`                                              // length 12
	NetworkManagementInformationCode                                      string `json:"networkManagementInformationCode,omitempty"`                                      // length 3
	FalconCaseStatus                                                      string `json:"falconCaseStatus,omitempty"`                                                      // length 255
	FalconCaseSubStatus                                                   string `json:"falconCaseSubStatus,omitempty"`                                                   // length 255
	FalconBlockCode                                                       string `json:"falconBlockCode,omitempty"`                                                       // length 255
	FalconFraudCode                                                       string `json:"falconFraudCode,omitempty"`                                                       // length 255
	TransactionFeeAmount                                                  string `json:"transactionFeeAmount,omitempty"`                                                  // length 10
	SettlementFeeAmount                                                   string `json:"settlementFeeAmount,omitempty"`                                                   // length 10
	AdditionalFees_CCA                                                    string `json:"additionalFees_CCA,omitempty"`                                                    // length 10
	AdditionalFees_ICA                                                    string `json:"additionalFees_ICA,omitempty"`                                                    // length 10
	CardAcceptorRegionCode                                                string `json:"cardAcceptorRegionCode,omitempty"`                                                // length 2
	CardAcceptorCountryCode                                               string `json:"cardAcceptorCountryCode,omitempty"`                                               // length 2
	NationalPointOfServiceCondition_TerminalUnattended                    string `json:"nationalPointOfServiceCondition_TerminalUnattended,omitempty"`                    // length 1
	NationalPointOfServiceCondition_TerminalOperator                      string `json:"nationalPointOfServiceCondition_TerminalOperator,omitempty"`                      // length 1
	NationalPointOfServiceCondition_TerminalPremises                      string `json:"nationalPointOfServiceCondition_TerminalPremises,omitempty"`                      // length 1
	NationalPointOfServiceCondition_CardPresentation                      string `json:"nationalPointOfServiceCondition_CardPresentation,omitempty"`                      // length 1
	NationalPointOfServiceCondition_CardPresence                          string `json:"nationalPointOfServiceCondition_CardPresence,omitempty"`                          // length 1
	NationalPointOfServiceCondition_CardRetention                         string `json:"nationalPointOfServiceCondition_CardRetention,omitempty"`                         // length 1
	NationalPointOfServiceCondition_CardTransaction                       string `json:"nationalPointOfServiceCondition_CardTransaction,omitempty"`                       // length 1
	NationalPointOfServiceCondition_SecurityCondition                     string `json:"nationalPointOfServiceCondition_SecurityCondition,omitempty"`                     // length 1
	NationalPointOfServiceCondition_TerminalType                          string `json:"nationalPointOfServiceCondition_TerminalType,omitempty"`                          // length 2
	NationalPointOfServiceCondition_TerminalEntryCapability               string `json:"nationalPointOfServiceCondition_TerminalEntryCapability,omitempty"`               // length 1
	PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator       string `json:"privatelyDefinedData_TransactionLevel_CredentialOnFileIndicator,omitempty"`       // length 1
	PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator string `json:"privatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator,omitempty"` // length 1
	AvsResult                                                             string `json:"avsResult,omitempty"`                                                             // length 1

	Typed_UserEventId                int64   `json:"-"` // length 19
	Typed_CustomerId                 int64   `json:"-"` // length 10
	Typed_CardId                     int64   `json:"-"` // length 10
	Typed_TransactionId              int64   `json:"-"` // length 19
	Typed_AuthorizationTransactionId int64   `json:"-"` // length 19
	Typed_CashbackAmount             float64 `json:"-"` // length 10
	Typed_SurchargeAmount            float64 `json:"-"` // length 10
	Typed_CashDepositAmount          float64 `json:"-"` // length 10
	Typed_CheckDepositAmount         float64 `json:"-"` // length 10
	Typed_Amount                     float64 `json:"-"` // length 10
	Typed_FromAccountAmount          float64 `json:"-"` // length 10
	Typed_FromAccountId              int64   `json:"-"` // length 10
	Typed_FromAvailableAmount        float64 `json:"-"` // length 10
	Typed_FromPendingAmount          float64 `json:"-"` // length 10
	Typed_FromPrimaryCustomerId      int64   `json:"-"` // length 10
	Typed_MasterId                   int64   `json:"-"` // length long 19
	Typed_ModifiedById               int64   `json:"-"` // length 10
	Typed_PayloadTypeId              int64   `json:"-"` // length 10
	Typed_ToAccountAmount            float64 `json:"-"` // length 10
	Typed_ToAccountId                int64   `json:"-"` // length 10
	Typed_ToAvailableAmount          float64 `json:"-"` // length 10
	Typed_ToPendingAmount            float64 `json:"-"` // length 10
	Typed_ToPrimaryCustomerId        int64   `json:"-"` // length 10
	Typed_EventTypeId                int64   `json:"-"` // length 10
	Typed_NetworkProviderTypeId      int64   `json:"-"` // length 1
	Typed_AdditionalAmounts_Purchase float64 `json:"-"` // length 10
	Typed_AdditionalAmounts_Gratuity float64 `json:"-"` // length int 10
	Typed_TransactionFeeAmount       float64 `json:"-"` // length 10
	Typed_SettlementFeeAmount        float64 `json:"-"` // length 10
	Typed_AdditionalFees_CCA         float64 `json:"-"` // length 10
	Typed_AdditionalFees_ICA         float64 `json:"-"` // length 10
}

type TypedCardEventNotificationFileModel struct {
	UserEventId                                                           int64   `json:"userEventId,omitempty"`                                                           // length 19
	CustomerId                                                            int64   `json:"customerId,omitempty"`                                                            // length 10
	CardId                                                                int64   `json:"cardId,omitempty"`                                                                // length 10
	TransactionId                                                         int64   `json:"transactionId,omitempty"`                                                         // length 19
	AuthorizationTransactionId                                            int64   `json:"authorizationTransactionId,omitempty"`                                            // length 19
	RequestTypeCode                                                       string  `json:"requestTypeCode,omitempty"`                                                       // length 3
	Reserved                                                              string  `json:"reserved,omitempty"`                                                              // length 12
	MerchantCategoryCode                                                  string  `json:"merchantCategoryCode,omitempty"`                                                  // length 4
	MerchantGroupTypeCode                                                 string  `json:"merchantGroupTypeCode,omitempty"`                                                 // length 6
	CashbackAmount                                                        float64 `json:"cashbackAmount,omitempty"`                                                        // length 10
	SurchargeAmount                                                       float64 `json:"surchargeAmount,omitempty"`                                                       // length 10
	CashDepositAmount                                                     float64 `json:"cashDepositAmount,omitempty"`                                                     // length 10
	CheckDepositAmount                                                    float64 `json:"checkDepositAmount,omitempty"`                                                    // length 10
	TerminalId                                                            string  `json:"terminalId,omitempty"`                                                            // length 15
	MerchantId                                                            string  `json:"merchantId,omitempty"`                                                            // length 15
	MerchantLocation                                                      string  `json:"merchantLocation,omitempty"`                                                      // length 25
	MerchantCity                                                          string  `json:"merchantCity,omitempty"`                                                          // length 13
	MerchantStateCode                                                     string  `json:"merchantStateCode,omitempty"`                                                     // length 2
	MerchantZipCode                                                       string  `json:"merchantZipCode,omitempty"`                                                       // length 9
	MerchantCountryCode                                                   string  `json:"merchantCountryCode,omitempty"`                                                   // length 2
	PanEntryModeTypeCode                                                  string  `json:"panEntryModeTypeCode,omitempty"`                                                  // length 10
	PinEntryModeTypeCode                                                  string  `json:"pinEntryModeTypeCode,omitempty"`                                                  // length 10
	CardHolderPresenceTypeCode                                            string  `json:"cardHolderPresenceTypeCode,omitempty"`                                            // length 10
	CardPresenceTypeCode                                                  string  `json:"cardPresenceTypeCode,omitempty"`                                                  // length 10
	TerminalUnattendedTypeCode                                            string  `json:"terminalUnattendedTypeCode,omitempty"`                                            // length 10
	TerminalPremisesTypeCode                                              string  `json:"terminalPremisesTypeCode,omitempty"`                                              // length 10
	CustomerTag                                                           string  `json:"customerTag,omitempty"`                                                           // length 50
	CardTag                                                               string  `json:"cardTag,omitempty"`                                                               // length 50
	Amount                                                                float64 `json:"amount,omitempty"`                                                                // length 10
	AvailableDate                                                         string  `json:"availableDate,omitempty"`                                                         // length 34
	CreatedDate                                                           string  `json:"createdDate,omitempty"`                                                           // length 34
	CustomField1                                                          string  `json:"customField1,omitempty"`                                                          // length 50
	Description                                                           string  `json:"description,omitempty"`                                                           // length 255
	DenialReason                                                          string  `json:"denialReason,omitempty"`                                                          // length 255
	FeeCode                                                               string  `json:"feeCode,omitempty"`                                                               // length 3
	FeeDescription                                                        string  `json:"feeDescription,omitempty"`                                                        // length 50
	FriendlyDescription                                                   string  `json:"friendlyDescription,omitempty"`                                                   // length 255
	FromAccountAmount                                                     float64 `json:"fromAccountAmount,omitempty"`                                                     // length 10
	FromAccountId                                                         int64   `json:"fromAccountId,omitempty"`                                                         // length 10
	FromAccountNumberMasked                                               string  `json:"fromAccountNumberMasked,omitempty"`                                               // length 50
	FromAvailableAmount                                                   float64 `json:"fromAvailableAmount,omitempty"`                                                   // length 10
	FromCategory                                                          string  `json:"fromCategory,omitempty"`                                                          // length 50
	FromCreatedDate                                                       string  `json:"fromCreatedDate,omitempty"`                                                       // length 34
	FromCustomField1                                                      string  `json:"fromCustomField1,omitempty"`                                                      // length 50
	FromCustomField2                                                      string  `json:"fromCustomField2,omitempty"`                                                      // length 50
	FromCustomField3                                                      string  `json:"fromCustomField3,omitempty"`                                                      // length 50
	FromCustomField4                                                      string  `json:"fromCustomField4,omitempty"`                                                      // length 50
	FromCustomField5                                                      string  `json:"fromCustomField5,omitempty"`                                                      // length 50
	FromLegalName1                                                        string  `json:"fromLegalName1,omitempty"`                                                        // length 100
	FromLegalName2                                                        string  `json:"fromLegalName2,omitempty"`                                                        // length 100
	FromName                                                              string  `json:"fromName,omitempty"`                                                              // length 50
	FromPendingAmount                                                     float64 `json:"fromPendingAmount,omitempty"`                                                     // length 10
	FromPrimaryCustomerId                                                 int64   `json:"fromPrimaryCustomerId,omitempty"`                                                 // length 10
	FromSubCategory                                                       string  `json:"fromSubCategory,omitempty"`                                                       // length 50
	FromTag                                                               string  `json:"fromTag,omitempty"`                                                               // length 50
	InstitutionName                                                       string  `json:"institutionName,omitempty"`                                                       // length 50
	MasterId                                                              int64   `json:"masterId,omitempty"`                                                              // length 19
	ModifiedById                                                          int64   `json:"modifiedById,omitempty"`                                                          // length 10
	PayloadTypeId                                                         int64   `json:"payloadTypeId,omitempty"`                                                         // length 10
	SettledDate                                                           string  `json:"settledDate,omitempty"`                                                           // length 34
	SubType                                                               string  `json:"subType,omitempty"`                                                               // length 255
	SubTypeCode                                                           string  `json:"subTypeCode,omitempty"`                                                           // length 6
	Tag                                                                   string  `json:"tag,omitempty"`                                                                   // length 50
	ToAccountAmount                                                       float64 `json:"toAccountAmount,omitempty"`                                                       // length 10
	ToAccountId                                                           int64   `json:"toAccountId,omitempty"`                                                           // length 10
	ToAccountNumberMasked                                                 string  `json:"toAccountNumberMasked,omitempty"`                                                 // length 50
	ToAvailableAmount                                                     float64 `json:"toAvailableAmount,omitempty"`                                                     // length 10
	ToCategory                                                            string  `json:"toCategory,omitempty"`                                                            // length 50
	ToCreatedDate                                                         string  `json:"toCreatedDate,omitempty"`                                                         // length 34
	ToCustomField1                                                        string  `json:"toCustomField1,omitempty"`                                                        // length 50
	ToCustomField2                                                        string  `json:"toCustomField2,omitempty"`                                                        // length 50
	ToCustomField3                                                        string  `json:"toCustomField3,omitempty"`                                                        // length 50
	ToCustomField4                                                        string  `json:"toCustomField4,omitempty"`                                                        // length 50
	ToCustomField5                                                        string  `json:"toCustomField5,omitempty"`                                                        // length 50
	ToLegalName1                                                          string  `json:"toLegalName1,omitempty"`                                                          // length 100
	ToLegalName2                                                          string  `json:"toLegalName2,omitempty"`                                                          // length 100
	ToName                                                                string  `json:"toName,omitempty"`                                                                // length 50
	ToPendingAmount                                                       float64 `json:"toPendingAmount,omitempty"`                                                       // length 10
	ToPrimaryCustomerId                                                   int64   `json:"toPrimaryCustomerId,omitempty"`                                                   // length 10
	ToSubCategory                                                         string  `json:"toSubCategory,omitempty"`                                                         // length 50
	ToTag                                                                 string  `json:"toTag,omitempty"`                                                                 // length 50
	Type                                                                  string  `json:"type,omitempty"`                                                                  // length 50
	TypeCode                                                              string  `json:"typeCode,omitempty"`                                                              // length 6
	EventTypeId                                                           int64   `json:"eventTypeId,omitempty"`                                                           // length 10
	NetworkProviderTypeId                                                 int64   `json:"networkProviderTypeId,omitempty"`                                                 // length 1
	PointOfServicePanEntryMode                                            string  `json:"pointOfServicePanEntryMode,omitempty"`                                            // length 2
	PointOfServicePinEntryMode                                            string  `json:"pointOfServicePinEntryMode,omitempty"`                                            // length 1
	Cvv2PresenceIndicator                                                 string  `json:"cvv2PresenceIndicator,omitempty"`                                                 // length 1
	Cvv2Result                                                            string  `json:"cvv2Result,omitempty"`                                                            // length 1
	Token                                                                 string  `json:"token,omitempty"`                                                                 // length 19
	TokenAssuranceLevel                                                   string  `json:"tokenAssuranceLevel,omitempty"`                                                   // length 2
	DigitalWalletTokenRequestorTypeId                                     string  `json:"digitalWalletTokenRequestorTypeId,omitempty"`                                     // length 2
	TokenExpirationDate                                                   string  `json:"tokenExpirationDate,omitempty"`                                                   // length 4
	PaymentAccountReferenceNumber                                         string  `json:"paymentAccountReferenceNumber,omitempty"`                                         // length 29
	MessageTypeIndicator                                                  string  `json:"messageTypeIndicator,omitempty"`                                                  // length 4
	OutputMessageTypeIndicator                                            string  `json:"outputMessageTypeIndicator,omitempty"`                                            // length 4
	OutputResponseCode                                                    string  `json:"outputResponseCode,omitempty"`                                                    // length 2
	SystemTraceAuditNumber                                                string  `json:"systemTraceAuditNumber,omitempty"`                                                // length 6
	AcquirerInstitutionCountryCode                                        string  `json:"acquirerInstitutionCountryCode,omitempty"`                                        // length 3
	AuthorizationIdentificationResponse                                   string  `json:"authorizationIdentificationResponse,omitempty"`                                   // length 12
	ResponseCode                                                          string  `json:"responseCode,omitempty"`                                                          // length 2
	PinValidationCode                                                     string  `json:"pinValidationCode,omitempty"`                                                     // length 1
	AdditionalAmounts_Purchase                                            float64 `json:"additionalAmounts_Purchase,omitempty"`                                            // length 10
	AdditionalAmounts_Gratuity                                            float64 `json:"additionalAmounts_Gratuity,omitempty"`                                            // length 10
	Advice_OriginatorCode                                                 string  `json:"advice_OriginatorCode,omitempty"`                                                 // length 1
	Advice_ReasonCode                                                     string  `json:"advice_ReasonCode,omitempty"`                                                     // length 1
	PrivatelyDefinedData_IssuerNetworkIdCode                              string  `json:"privatelyDefinedData_IssuerNetworkIdCode,omitempty"`                              // length 3
	PrivatelyDefinedData_AdditionalTransactionElement_FallbackIndicator   string  `json:"privatelyDefinedData_AdditionalTransactionElement_FallbackIndicator,omitempty"`   // length 1
	PrivatelyDefinedData_ProcessingFlag_SpecialTransactionIndicator       string  `json:"privatelyDefinedData_ProcessingFlag_SpecialTransactionIndicator,omitempty"`       // length 1
	PrivatelyDefinedData_ProcessingFlag_ISAIndicator                      string  `json:"privatelyDefinedData_ProcessingFlag_ISAIndicator,omitempty"`                      // length 1
	PrivatelyDefinedData_ProcessingFlag_PartialAuthIndicator              string  `json:"privatelyDefinedData_ProcessingFlag_PartialAuthIndicator,omitempty"`              // length 1
	PrivatelyDefinedData_RiskData_ScoreSource                             string  `json:"privatelyDefinedData_RiskData_ScoreSource,omitempty"`                             // length 1
	PrivatelyDefinedData_RiskData_ScoreValue                              string  `json:"privatelyDefinedData_RiskData_ScoreValue,omitempty"`                              // length 4
	PrivatelyDefinedData_RiskData_ResponseCode                            string  `json:"privatelyDefinedData_RiskData_ResponseCode,omitempty"`                            // length 1
	PrivatelyDefinedData_RiskData_FalconReason1                           string  `json:"privatelyDefinedData_RiskData_FalconReason1,omitempty"`                           // length 2
	PrivatelyDefinedData_RiskData_FalconReason2                           string  `json:"privatelyDefinedData_RiskData_FalconReason2,omitempty"`                           // length 2
	PrivatelyDefinedData_RiskData_FalconReason3                           string  `json:"privatelyDefinedData_RiskData_FalconReason3,omitempty"`                           // length 2
	PrivatelyDefinedData_RiskData_VisaRiskScore                           string  `json:"privatelyDefinedData_RiskData_VisaRiskScore,omitempty"`                           // length 2
	PrivatelyDefinedData_RiskData_VisaRiskReason                          string  `json:"privatelyDefinedData_RiskData_VisaRiskReason,omitempty"`                          // length 2
	PrivatelyDefinedData_RiskData_VisaRiskConditionCode1                  string  `json:"privatelyDefinedData_RiskData_VisaRiskConditionCode1,omitempty"`                  // length 2
	PrivatelyDefinedData_RiskData_VisaRiskConditionCode2                  string  `json:"privatelyDefinedData_RiskData_VisaRiskConditionCode2,omitempty"`                  // length 2
	PrivatelyDefinedData_RiskData_VisaRiskConditionCode3                  string  `json:"privatelyDefinedData_RiskData_VisaRiskConditionCode3,omitempty"`                  // length 2
	PrivatelyDefinedData_RiskData_VAAConditionCode1Rank                   string  `json:"privatelyDefinedData_RiskData_VAAConditionCode1Rank,omitempty"`                   // length 2
	PrivatelyDefinedData_RiskData_RTDResultCode                           string  `json:"privatelyDefinedData_RiskData_RTDResultCode,omitempty"`                           // length 1
	PrivatelyDefinedData_RiskData_TravelStatusIndicator                   string  `json:"privatelyDefinedData_RiskData_TravelStatusIndicator,omitempty"`                   // length 1
	TextInfo                                                              string  `json:"textInfo,omitempty"`                                                              // length 255
	Track2Data_ServiceCode                                                string  `json:"track2Data_ServiceCode,omitempty"`                                                // length 3
	RetrievalReferenceNumber                                              string  `json:"retrievalReferenceNumber,omitempty"`                                              // length 12
	NetworkManagementInformationCode                                      string  `json:"networkManagementInformationCode,omitempty"`                                      // length 3
	FalconCaseStatus                                                      string  `json:"falconCaseStatus,omitempty"`                                                      // length 255
	FalconCaseSubStatus                                                   string  `json:"falconCaseSubStatus,omitempty"`                                                   // length 255
	FalconBlockCode                                                       string  `json:"falconBlockCode,omitempty"`                                                       // length 255
	FalconFraudCode                                                       string  `json:"falconFraudCode,omitempty"`                                                       // length 255
	TransactionFeeAmount                                                  float64 `json:"transactionFeeAmount,omitempty"`                                                  // length 10
	SettlementFeeAmount                                                   float64 `json:"settlementFeeAmount,omitempty"`                                                   // length 10
	AdditionalFees_CCA                                                    float64 `json:"additionalFees_CCA,omitempty"`                                                    // length 10
	AdditionalFees_ICA                                                    float64 `json:"additionalFees_ICA,omitempty"`                                                    // length 10
	CardAcceptorRegionCode                                                string  `json:"cardAcceptorRegionCode,omitempty"`                                                // length 2
	CardAcceptorCountryCode                                               string  `json:"cardAcceptorCountryCode,omitempty"`                                               // length 2
	NationalPointOfServiceCondition_TerminalUnattended                    string  `json:"nationalPointOfServiceCondition_TerminalUnattended,omitempty"`                    // length 1
	NationalPointOfServiceCondition_TerminalOperator                      string  `json:"nationalPointOfServiceCondition_TerminalOperator,omitempty"`                      // length 1
	NationalPointOfServiceCondition_TerminalPremises                      string  `json:"nationalPointOfServiceCondition_TerminalPremises,omitempty"`                      // length 1
	NationalPointOfServiceCondition_CardPresentation                      string  `json:"nationalPointOfServiceCondition_CardPresentation,omitempty"`                      // length 1
	NationalPointOfServiceCondition_CardPresence                          string  `json:"nationalPointOfServiceCondition_CardPresence,omitempty"`                          // length 1
	NationalPointOfServiceCondition_CardRetention                         string  `json:"nationalPointOfServiceCondition_CardRetention,omitempty"`                         // length 1
	NationalPointOfServiceCondition_CardTransaction                       string  `json:"nationalPointOfServiceCondition_CardTransaction,omitempty"`                       // length 1
	NationalPointOfServiceCondition_SecurityCondition                     string  `json:"nationalPointOfServiceCondition_SecurityCondition,omitempty"`                     // length 1
	NationalPointOfServiceCondition_TerminalType                          string  `json:"nationalPointOfServiceCondition_TerminalType,omitempty"`                          // length 2
	NationalPointOfServiceCondition_TerminalEntryCapability               string  `json:"nationalPointOfServiceCondition_TerminalEntryCapability,omitempty"`               // length 1
	PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator       string  `json:"privatelyDefinedData_TransactionLevel_CredentialOnFileIndicator,omitempty"`       // length 1
	PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator string  `json:"privatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator,omitempty"` // length 1
	AvsResult                                                             string  `json:"avsResult,omitempty"`                                                             // length 1
}

func (mdl *CardEventNotificationFileModel) IsTabDelimitedFile() bool {
	return false
}

func (mdl *CardEventNotificationFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *CardEventNotificationFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedCardEventNotificationFileModel{
			UserEventId:                              mdl.Typed_UserEventId,
			CustomerId:                               mdl.Typed_CustomerId,
			CardId:                                   mdl.Typed_CardId,
			TransactionId:                            mdl.Typed_TransactionId,
			AuthorizationTransactionId:               mdl.Typed_AuthorizationTransactionId,
			RequestTypeCode:                          mdl.RequestTypeCode,
			Reserved:                                 mdl.Reserved,
			MerchantCategoryCode:                     mdl.MerchantCategoryCode,
			MerchantGroupTypeCode:                    mdl.MerchantGroupTypeCode,
			CashbackAmount:                           mdl.Typed_CashbackAmount,
			SurchargeAmount:                          mdl.Typed_SurchargeAmount,
			CashDepositAmount:                        mdl.Typed_CashDepositAmount,
			CheckDepositAmount:                       mdl.Typed_CheckDepositAmount,
			TerminalId:                               mdl.TerminalId,
			MerchantId:                               mdl.MerchantId,
			MerchantLocation:                         mdl.MerchantLocation,
			MerchantCity:                             mdl.MerchantCity,
			MerchantStateCode:                        mdl.MerchantStateCode,
			MerchantZipCode:                          mdl.MerchantZipCode,
			MerchantCountryCode:                      mdl.MerchantCountryCode,
			PanEntryModeTypeCode:                     mdl.PanEntryModeTypeCode,
			PinEntryModeTypeCode:                     mdl.PinEntryModeTypeCode,
			CardHolderPresenceTypeCode:               mdl.CardHolderPresenceTypeCode,
			CardPresenceTypeCode:                     mdl.CardPresenceTypeCode,
			TerminalUnattendedTypeCode:               mdl.TerminalUnattendedTypeCode,
			TerminalPremisesTypeCode:                 mdl.TerminalPremisesTypeCode,
			CustomerTag:                              mdl.CustomerTag,
			CardTag:                                  mdl.CardTag,
			Amount:                                   mdl.Typed_Amount,
			AvailableDate:                            mdl.AvailableDate,
			CreatedDate:                              mdl.CreatedDate,
			CustomField1:                             mdl.CustomField1,
			Description:                              mdl.Description,
			DenialReason:                             mdl.DenialReason,
			FeeCode:                                  mdl.FeeCode,
			FeeDescription:                           mdl.FeeDescription,
			FriendlyDescription:                      mdl.FriendlyDescription,
			FromAccountAmount:                        mdl.Typed_FromAccountAmount,
			FromAccountId:                            mdl.Typed_FromAccountId,
			FromAccountNumberMasked:                  mdl.FromAccountNumberMasked,
			FromAvailableAmount:                      mdl.Typed_FromAvailableAmount,
			FromCategory:                             mdl.FromCategory,
			FromCreatedDate:                          mdl.FromCreatedDate,
			FromCustomField1:                         mdl.FromCustomField1,
			FromCustomField2:                         mdl.FromCustomField2,
			FromCustomField3:                         mdl.FromCustomField3,
			FromCustomField4:                         mdl.FromCustomField4,
			FromCustomField5:                         mdl.FromCustomField5,
			FromLegalName1:                           mdl.FromLegalName1,
			FromLegalName2:                           mdl.FromLegalName2,
			FromName:                                 mdl.FromName,
			FromPendingAmount:                        mdl.Typed_FromPendingAmount,
			FromPrimaryCustomerId:                    mdl.Typed_FromPrimaryCustomerId,
			FromSubCategory:                          mdl.FromSubCategory,
			FromTag:                                  mdl.FromTag,
			InstitutionName:                          mdl.InstitutionName,
			MasterId:                                 mdl.Typed_MasterId,
			ModifiedById:                             mdl.Typed_ModifiedById,
			PayloadTypeId:                            mdl.Typed_PayloadTypeId,
			SettledDate:                              mdl.SettledDate,
			SubType:                                  mdl.SubType,
			SubTypeCode:                              mdl.SubTypeCode,
			Tag:                                      mdl.Tag,
			ToAccountAmount:                          mdl.Typed_ToAccountAmount,
			ToAccountId:                              mdl.Typed_ToAccountId,
			ToAccountNumberMasked:                    mdl.ToAccountNumberMasked,
			ToAvailableAmount:                        mdl.Typed_ToAvailableAmount,
			ToCategory:                               mdl.ToCategory,
			ToCreatedDate:                            mdl.ToCreatedDate,
			ToCustomField1:                           mdl.ToCustomField1,
			ToCustomField2:                           mdl.ToCustomField2,
			ToCustomField3:                           mdl.ToCustomField3,
			ToCustomField4:                           mdl.ToCustomField4,
			ToCustomField5:                           mdl.ToCustomField5,
			ToLegalName1:                             mdl.ToLegalName1,
			ToLegalName2:                             mdl.ToLegalName2,
			ToName:                                   mdl.ToName,
			ToPendingAmount:                          mdl.Typed_ToPendingAmount,
			ToPrimaryCustomerId:                      mdl.Typed_ToPrimaryCustomerId,
			ToSubCategory:                            mdl.ToSubCategory,
			ToTag:                                    mdl.ToTag,
			Type:                                     mdl.Type,
			TypeCode:                                 mdl.TypeCode,
			EventTypeId:                              mdl.Typed_EventTypeId,
			NetworkProviderTypeId:                    mdl.Typed_NetworkProviderTypeId,
			PointOfServicePanEntryMode:               mdl.PointOfServicePanEntryMode,
			PointOfServicePinEntryMode:               mdl.PointOfServicePinEntryMode,
			Cvv2PresenceIndicator:                    mdl.Cvv2PresenceIndicator,
			Cvv2Result:                               mdl.Cvv2Result,
			Token:                                    mdl.Token,
			TokenAssuranceLevel:                      mdl.TokenAssuranceLevel,
			DigitalWalletTokenRequestorTypeId:        mdl.DigitalWalletTokenRequestorTypeId,
			TokenExpirationDate:                      mdl.TokenExpirationDate,
			PaymentAccountReferenceNumber:            mdl.PaymentAccountReferenceNumber,
			MessageTypeIndicator:                     mdl.MessageTypeIndicator,
			OutputMessageTypeIndicator:               mdl.OutputMessageTypeIndicator,
			OutputResponseCode:                       mdl.OutputResponseCode,
			SystemTraceAuditNumber:                   mdl.SystemTraceAuditNumber,
			AcquirerInstitutionCountryCode:           mdl.AcquirerInstitutionCountryCode,
			AuthorizationIdentificationResponse:      mdl.AuthorizationIdentificationResponse,
			ResponseCode:                             mdl.ResponseCode,
			PinValidationCode:                        mdl.PinValidationCode,
			AdditionalAmounts_Purchase:               mdl.Typed_AdditionalAmounts_Purchase,
			AdditionalAmounts_Gratuity:               mdl.Typed_AdditionalAmounts_Gratuity,
			Advice_OriginatorCode:                    mdl.Advice_OriginatorCode,
			Advice_ReasonCode:                        mdl.Advice_ReasonCode,
			PrivatelyDefinedData_IssuerNetworkIdCode: mdl.PrivatelyDefinedData_IssuerNetworkIdCode,
			PrivatelyDefinedData_AdditionalTransactionElement_FallbackIndicator: mdl.PrivatelyDefinedData_AdditionalTransactionElement_FallbackIndicator,
			PrivatelyDefinedData_ProcessingFlag_SpecialTransactionIndicator:     mdl.PrivatelyDefinedData_ProcessingFlag_SpecialTransactionIndicator,
			PrivatelyDefinedData_ProcessingFlag_ISAIndicator:                    mdl.PrivatelyDefinedData_ProcessingFlag_ISAIndicator,
			PrivatelyDefinedData_ProcessingFlag_PartialAuthIndicator:            mdl.PrivatelyDefinedData_ProcessingFlag_PartialAuthIndicator,
			PrivatelyDefinedData_RiskData_ScoreSource:                           mdl.PrivatelyDefinedData_RiskData_ScoreSource,
			PrivatelyDefinedData_RiskData_ScoreValue:                            mdl.PrivatelyDefinedData_RiskData_ScoreValue,
			PrivatelyDefinedData_RiskData_ResponseCode:                          mdl.PrivatelyDefinedData_RiskData_ResponseCode,
			PrivatelyDefinedData_RiskData_FalconReason1:                         mdl.PrivatelyDefinedData_RiskData_FalconReason1,
			PrivatelyDefinedData_RiskData_FalconReason2:                         mdl.PrivatelyDefinedData_RiskData_FalconReason2,
			PrivatelyDefinedData_RiskData_FalconReason3:                         mdl.PrivatelyDefinedData_RiskData_FalconReason3,
			PrivatelyDefinedData_RiskData_VisaRiskScore:                         mdl.PrivatelyDefinedData_RiskData_VisaRiskScore,
			PrivatelyDefinedData_RiskData_VisaRiskReason:                        mdl.PrivatelyDefinedData_RiskData_VisaRiskReason,
			PrivatelyDefinedData_RiskData_VisaRiskConditionCode1:                mdl.PrivatelyDefinedData_RiskData_VisaRiskConditionCode1,
			PrivatelyDefinedData_RiskData_VisaRiskConditionCode2:                mdl.PrivatelyDefinedData_RiskData_VisaRiskConditionCode2,
			PrivatelyDefinedData_RiskData_VisaRiskConditionCode3:                mdl.PrivatelyDefinedData_RiskData_VisaRiskConditionCode3,
			PrivatelyDefinedData_RiskData_VAAConditionCode1Rank:                 mdl.PrivatelyDefinedData_RiskData_VAAConditionCode1Rank,
			PrivatelyDefinedData_RiskData_RTDResultCode:                         mdl.PrivatelyDefinedData_RiskData_RTDResultCode,
			PrivatelyDefinedData_RiskData_TravelStatusIndicator:                 mdl.PrivatelyDefinedData_RiskData_TravelStatusIndicator,
			TextInfo:                         mdl.TextInfo,
			Track2Data_ServiceCode:           mdl.Track2Data_ServiceCode,
			RetrievalReferenceNumber:         mdl.RetrievalReferenceNumber,
			NetworkManagementInformationCode: mdl.NetworkManagementInformationCode,
			FalconCaseStatus:                 mdl.FalconCaseStatus,
			FalconCaseSubStatus:              mdl.FalconCaseSubStatus,
			FalconBlockCode:                  mdl.FalconBlockCode,
			FalconFraudCode:                  mdl.FalconFraudCode,
			TransactionFeeAmount:             mdl.Typed_TransactionFeeAmount,
			SettlementFeeAmount:              mdl.Typed_SettlementFeeAmount,
			AdditionalFees_CCA:               mdl.Typed_AdditionalFees_CCA,
			AdditionalFees_ICA:               mdl.Typed_AdditionalFees_ICA,
			CardAcceptorRegionCode:           mdl.CardAcceptorRegionCode,
			CardAcceptorCountryCode:          mdl.CardAcceptorCountryCode,
			NationalPointOfServiceCondition_TerminalUnattended:                    mdl.NationalPointOfServiceCondition_TerminalUnattended,
			NationalPointOfServiceCondition_TerminalOperator:                      mdl.NationalPointOfServiceCondition_TerminalOperator,
			NationalPointOfServiceCondition_TerminalPremises:                      mdl.NationalPointOfServiceCondition_TerminalPremises,
			NationalPointOfServiceCondition_CardPresentation:                      mdl.NationalPointOfServiceCondition_CardPresentation,
			NationalPointOfServiceCondition_CardPresence:                          mdl.NationalPointOfServiceCondition_CardPresence,
			NationalPointOfServiceCondition_CardRetention:                         mdl.NationalPointOfServiceCondition_CardRetention,
			NationalPointOfServiceCondition_CardTransaction:                       mdl.NationalPointOfServiceCondition_CardTransaction,
			NationalPointOfServiceCondition_SecurityCondition:                     mdl.NationalPointOfServiceCondition_SecurityCondition,
			NationalPointOfServiceCondition_TerminalType:                          mdl.NationalPointOfServiceCondition_TerminalType,
			NationalPointOfServiceCondition_TerminalEntryCapability:               mdl.NationalPointOfServiceCondition_TerminalEntryCapability,
			PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator:       mdl.PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator,
			PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator: mdl.PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator,
			AvsResult: mdl.AvsResult,
		}
	}
}

func (mdl *CardEventNotificationFileModel) GetValues() []string {
	return []string{}
}

func (mdl *CardEventNotificationFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	pos, err = ConsumeIntField(&mdl.UserEventId, &mdl.Typed_UserEventId, line, pos, 19, preserve, err)
	pos, err = ConsumeIntField(&mdl.CustomerId, &mdl.Typed_CustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.CardId, &mdl.Typed_CardId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.TransactionId, &mdl.Typed_TransactionId, line, pos, 19, preserve, err)
	pos, err = ConsumeIntField(&mdl.AuthorizationTransactionId, &mdl.Typed_AuthorizationTransactionId, line, pos, 19, preserve, err)
	pos, err = ConsumeField(&mdl.RequestTypeCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.Reserved, line, pos, 12, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantCategoryCode, line, pos, 4, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantGroupTypeCode, line, pos, 6, preserve, err)
	pos, err = ConsumeFloatField(&mdl.CashbackAmount, &mdl.Typed_CashbackAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.SurchargeAmount, &mdl.Typed_SurchargeAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.CashDepositAmount, &mdl.Typed_CashDepositAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.CheckDepositAmount, &mdl.Typed_CheckDepositAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.TerminalId, line, pos, 15, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantId, line, pos, 15, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantLocation, line, pos, 25, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantCity, line, pos, 13, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantStateCode, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantZipCode, line, pos, 9, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantCountryCode, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PanEntryModeTypeCode, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.PinEntryModeTypeCode, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.CardHolderPresenceTypeCode, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.CardPresenceTypeCode, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.TerminalUnattendedTypeCode, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.TerminalPremisesTypeCode, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.CustomerTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.CardTag, line, pos, 50, preserve, err)
	pos, err = ConsumeFloatField(&mdl.Amount, &mdl.Typed_Amount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.AvailableDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.CreatedDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.CustomField1, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.Description, line, pos, 255, preserve, err)
	pos, err = ConsumeField(&mdl.DenialReason, line, pos, 255, preserve, err)
	pos, err = ConsumeField(&mdl.FeeCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.FeeDescription, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.FriendlyDescription, line, pos, 255, preserve, err)
	pos, err = ConsumeFloatField(&mdl.FromAccountAmount, &mdl.Typed_FromAccountAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeIntField(&mdl.FromAccountId, &mdl.Typed_FromAccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.FromAccountNumberMasked, line, pos, 50, preserve, err)
	pos, err = ConsumeFloatField(&mdl.FromAvailableAmount, &mdl.Typed_FromAvailableAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.FromCategory, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.FromCreatedDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.FromCustomField1, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.FromCustomField2, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.FromCustomField3, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.FromCustomField4, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.FromCustomField5, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.FromLegalName1, line, pos, 100, preserve, err)
	pos, err = ConsumeField(&mdl.FromLegalName2, line, pos, 100, preserve, err)
	pos, err = ConsumeField(&mdl.FromName, line, pos, 50, preserve, err)
	pos, err = ConsumeFloatField(&mdl.FromPendingAmount, &mdl.Typed_FromPendingAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeIntField(&mdl.FromPrimaryCustomerId, &mdl.Typed_FromPrimaryCustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.FromSubCategory, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.FromTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.InstitutionName, line, pos, 50, preserve, err)
	pos, err = ConsumeIntField(&mdl.MasterId, &mdl.Typed_MasterId, line, pos, 19, preserve, err)
	pos, err = ConsumeIntField(&mdl.ModifiedById, &mdl.Typed_ModifiedById, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.PayloadTypeId, &mdl.Typed_PayloadTypeId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.SettledDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.SubType, line, pos, 255, preserve, err)
	pos, err = ConsumeField(&mdl.SubTypeCode, line, pos, 6, preserve, err)
	pos, err = ConsumeField(&mdl.Tag, line, pos, 50, preserve, err)
	pos, err = ConsumeFloatField(&mdl.ToAccountAmount, &mdl.Typed_ToAccountAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeIntField(&mdl.ToAccountId, &mdl.Typed_ToAccountId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.ToAccountNumberMasked, line, pos, 50, preserve, err)
	pos, err = ConsumeFloatField(&mdl.ToAvailableAmount, &mdl.Typed_ToAvailableAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.ToCategory, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.ToCreatedDate, line, pos, 34, preserve, err)
	pos, err = ConsumeField(&mdl.ToCustomField1, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.ToCustomField2, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.ToCustomField3, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.ToCustomField4, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.ToCustomField5, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.ToLegalName1, line, pos, 100, preserve, err)
	pos, err = ConsumeField(&mdl.ToLegalName2, line, pos, 100, preserve, err)
	pos, err = ConsumeField(&mdl.ToName, line, pos, 50, preserve, err)
	pos, err = ConsumeFloatField(&mdl.ToPendingAmount, &mdl.Typed_ToPendingAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeIntField(&mdl.ToPrimaryCustomerId, &mdl.Typed_ToPrimaryCustomerId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.ToSubCategory, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.ToTag, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.Type, line, pos, 50, preserve, err)
	pos, err = ConsumeField(&mdl.TypeCode, line, pos, 6, preserve, err)
	pos, err = ConsumeIntField(&mdl.EventTypeId, &mdl.Typed_EventTypeId, line, pos, 10, preserve, err)
	pos, err = ConsumeIntField(&mdl.NetworkProviderTypeId, &mdl.Typed_NetworkProviderTypeId, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PointOfServicePanEntryMode, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PointOfServicePinEntryMode, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.Cvv2PresenceIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.Cvv2Result, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.Token, line, pos, 19, preserve, err)
	pos, err = ConsumeField(&mdl.TokenAssuranceLevel, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.DigitalWalletTokenRequestorTypeId, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.TokenExpirationDate, line, pos, 4, preserve, err)
	pos, err = ConsumeField(&mdl.PaymentAccountReferenceNumber, line, pos, 29, preserve, err)
	pos, err = ConsumeField(&mdl.MessageTypeIndicator, line, pos, 4, preserve, err)
	pos, err = ConsumeField(&mdl.OutputMessageTypeIndicator, line, pos, 4, preserve, err)
	pos, err = ConsumeField(&mdl.OutputResponseCode, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.SystemTraceAuditNumber, line, pos, 6, preserve, err)
	pos, err = ConsumeField(&mdl.AcquirerInstitutionCountryCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.AuthorizationIdentificationResponse, line, pos, 12, preserve, err)
	pos, err = ConsumeField(&mdl.ResponseCode, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PinValidationCode, line, pos, 1, preserve, err)
	pos, err = ConsumeFloatField(&mdl.AdditionalAmounts_Purchase, &mdl.Typed_AdditionalAmounts_Purchase, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.AdditionalAmounts_Gratuity, &mdl.Typed_AdditionalAmounts_Gratuity, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.Advice_OriginatorCode, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.Advice_ReasonCode, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_IssuerNetworkIdCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_AdditionalTransactionElement_FallbackIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_ProcessingFlag_SpecialTransactionIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_ProcessingFlag_ISAIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_ProcessingFlag_PartialAuthIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_ScoreSource, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_ScoreValue, line, pos, 4, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_ResponseCode, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_FalconReason1, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_FalconReason2, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_FalconReason3, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_VisaRiskScore, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_VisaRiskReason, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_VisaRiskConditionCode1, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_VisaRiskConditionCode2, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_VisaRiskConditionCode3, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_VAAConditionCode1Rank, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_RTDResultCode, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_RiskData_TravelStatusIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.TextInfo, line, pos, 255, preserve, err)
	pos, err = ConsumeField(&mdl.Track2Data_ServiceCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.RetrievalReferenceNumber, line, pos, 12, preserve, err)
	pos, err = ConsumeField(&mdl.NetworkManagementInformationCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.FalconCaseStatus, line, pos, 255, preserve, err)
	pos, err = ConsumeField(&mdl.FalconCaseSubStatus, line, pos, 255, preserve, err)
	pos, err = ConsumeField(&mdl.FalconBlockCode, line, pos, 255, preserve, err)
	pos, err = ConsumeField(&mdl.FalconFraudCode, line, pos, 255, preserve, err)
	pos, err = ConsumeFloatField(&mdl.TransactionFeeAmount, &mdl.Typed_TransactionFeeAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.SettlementFeeAmount, &mdl.Typed_SettlementFeeAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.AdditionalFees_CCA, &mdl.Typed_AdditionalFees_CCA, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.AdditionalFees_ICA, &mdl.Typed_AdditionalFees_ICA, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.CardAcceptorRegionCode, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.CardAcceptorCountryCode, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.NationalPointOfServiceCondition_TerminalUnattended, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.NationalPointOfServiceCondition_TerminalOperator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.NationalPointOfServiceCondition_TerminalPremises, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.NationalPointOfServiceCondition_CardPresentation, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.NationalPointOfServiceCondition_CardPresence, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.NationalPointOfServiceCondition_CardRetention, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.NationalPointOfServiceCondition_CardTransaction, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.NationalPointOfServiceCondition_SecurityCondition, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.NationalPointOfServiceCondition_TerminalType, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.NationalPointOfServiceCondition_TerminalEntryCapability, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.AvsResult, line, pos, 1, preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
