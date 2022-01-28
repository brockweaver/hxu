package file

import (
	"fmt"
)

type CardTransactionFileModel struct {
	TransactionId                                                         string `json:"transactionId,omitempty"`                                                         // length long 19
	MasterId                                                              string `json:"masterId,omitempty"`                                                              // length long 19
	CardId                                                                string `json:"cardId,omitempty"`                                                                // length 10
	PanLastFour                                                           string `json:"panLastFour,omitempty"`                                                           // length 4
	CardHolderFirstName                                                   string `json:"cardHolderFirstName,omitempty"`                                                   // length 64
	CardHolderMiddleName                                                  string `json:"cardHolderMiddleName,omitempty"`                                                  // length 64
	CardHolderLastName                                                    string `json:"cardHolderLastName,omitempty"`                                                    // length 128
	CardAcceptorLocation                                                  string `json:"cardAcceptorLocation,omitempty"`                                                  // length 25
	CardAcceptorCity                                                      string `json:"cardAcceptorCity,omitempty"`                                                      // length 13
	CardAcceptorState                                                     string `json:"cardAcceptorState,omitempty"`                                                     // length 2
	CardAcceptorZip                                                       string `json:"cardAcceptorZip,omitempty"`                                                       // length 9
	RetrievalReferenceNumber                                              string `json:"retrievalReferenceNumber,omitempty"`                                              // length 12
	SystemTraceAuditNumber                                                string `json:"systemTraceAuditNumber,omitempty"`                                                // length 6
	MerchantId                                                            string `json:"merchantId,omitempty"`                                                            // length 15
	SubTypeCode                                                           string `json:"subTypeCode,omitempty"`                                                           // length 6
	MerchantGroupCode                                                     string `json:"merchantGroupCode,omitempty"`                                                     // length 6
	TerminalId                                                            string `json:"terminalId,omitempty"`                                                            // length 15
	CashbackAmount                                                        string `json:"cashbackAmount,omitempty"`                                                        // length int 10
	SurchargeAmount                                                       string `json:"surchargeAmount,omitempty"`                                                       // length int 10
	UnverifiedDepositAmount                                               string `json:"unverifiedDepositAmount,omitempty"`                                               // length int 10
	CashDepositAmount                                                     string `json:"cashDepositAmount,omitempty"`                                                     // length int 10
	AtmNetworkIndicator                                                   string `json:"atmNetworkIndicator,omitempty"`                                                   // length 3
	NetworkProviderTypeId                                                 string `json:"networkProviderTypeId,omitempty"`                                                 // length int 1
	TransactionFeeAmount                                                  string `json:"transactionFeeAmount,omitempty"`                                                  // length 10
	SettlementFeeAmount                                                   string `json:"settlementFeeAmount,omitempty"`                                                   // length 10
	AdditionalFees_CCA                                                    string `json:"additionalFees_CCA,omitempty"`                                                    // length 10
	AdditionalFees_ICA                                                    string `json:"additionalFees_ICA,omitempty"`                                                    // length 10
	MessageTypeIndicator                                                  string `json:"messageTypeIndicator,omitempty"`                                                  // length 4
	OutputMessageTypeIndicator                                            string `json:"outputMessageTypeIndicator,omitempty"`                                            // length 4
	OutputResponseCode                                                    string `json:"outputResponseCode,omitempty"`                                                    // length 2
	AcquirerInstitutionCountryCode                                        string `json:"acquirerInstitutionCountryCode,omitempty"`                                        // length 3
	AuthorizationIdentificationResponse                                   string `json:"authorizationIdentificationResponse,omitempty"`                                   // length 12
	ResponseCode                                                          string `json:"responseCode,omitempty"`                                                          // length 2
	PinValidationCode                                                     string `json:"pinValidationCode,omitempty"`                                                     // length 1
	AdditionalAmounts_Purchase                                            string `json:"additionalAmounts_Purchase,omitempty"`                                            // length 10
	AdditionalAmounts_Gratuity                                            string `json:"additionalAmounts_Gratuity,omitempty"`                                            // length 10
	Advice_OriginatorCode                                                 string `json:"advice_OriginatorCode,omitempty"`                                                 // length 1
	Advice_ReasonCode                                                     string `json:"advice_ReasonCode,omitempty"`                                                     // length 1
	PrivatelyDefintedData_IssuerNetworkIdCode                             string `json:"privatelyDefintedData_IssuerNetworkIdCode,omitempty"`                             // length 3
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
	CardAcceptorTerminalIdentification                                    string `json:"cardAcceptorTerminalIdentification,omitempty"`                                    // length 8
	Cvv2PresenceIndicator                                                 string `json:"cvv2PresenceIndicator,omitempty"`                                                 // length 1
	Cvv2Result                                                            string `json:"cvv2Result,omitempty"`                                                            // length 1
	Token                                                                 string `json:"token,omitempty"`                                                                 // length 19
	TokenAssuranceLevel                                                   string `json:"tokenAssuranceLevel,omitempty"`                                                   // length 2
	DigitalWalletTokenRequestorTypeId                                     string `json:"digitalWalletTokenRequestorTypeId,omitempty"`                                     // length 10
	TokenExpirationDate                                                   string `json:"tokenExpirationDate,omitempty"`                                                   // length 4
	PaymentAccountReferenceNumber                                         string `json:"paymentAccountReferenceNumber,omitempty"`                                         // length 29
	MerchantCategoryCode                                                  string `json:"merchantCategoryCode,omitempty"`                                                  // length 4
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
	PointOfServicePanEntryMode                                            string `json:"pointOfServicePanEntryMode,omitempty"`                                            // length 2
	PointOfServicePinEntryMode                                            string `json:"pointOfServicePinEntryMode,omitempty"`                                            // length 1
	PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator       string `json:"privatelyDefinedData_TransactionLevel_CredentialOnFileIndicator,omitempty"`       // length 1
	PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator string `json:"privatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator,omitempty"` // length 1
	AvsResult                                                             string `json:"avsResult,omitempty"`                                                             // length 1

	Typed_TransactionId              int64   `json:"-"` // length long 19
	Typed_MasterId                   int64   `json:"-"` // length long 19
	Typed_CardId                     int64   `json:"-"` // length 10
	Typed_CashbackAmount             float64 `json:"-"` // length int 10
	Typed_SurchargeAmount            float64 `json:"-"` // length int 10
	Typed_UnverifiedDepositAmount    float64 `json:"-"` // length int 10
	Typed_CashDepositAmount          float64 `json:"-"` // length int 10
	Typed_NetworkProviderTypeId      int64   `json:"-"` // length int 1
	Typed_TransactionFeeAmount       float64 `json:"-"` // length 10
	Typed_SettlementFeeAmount        float64 `json:"-"` // length 10
	Typed_AdditionalFees_CCA         float64 `json:"-"` // length 10
	Typed_AdditionalFees_ICA         float64 `json:"-"` // length 10
	Typed_AdditionalAmounts_Purchase float64 `json:"-"` // length 10
	Typed_AdditionalAmounts_Gratuity float64 `json:"-"` // length 10

}

type TypedCardTransactionFileModel struct {
	TransactionId                                                         int64   `json:"transactionId,omitempty"`                                                         // length long 19
	MasterId                                                              int64   `json:"masterId,omitempty"`                                                              // length long 19
	CardId                                                                int64   `json:"cardId,omitempty"`                                                                // length 10
	PanLastFour                                                           string  `json:"panLastFour,omitempty"`                                                           // length 4
	CardHolderFirstName                                                   string  `json:"cardHolderFirstName,omitempty"`                                                   // length 64
	CardHolderMiddleName                                                  string  `json:"cardHolderMiddleName,omitempty"`                                                  // length 64
	CardHolderLastName                                                    string  `json:"cardHolderLastName,omitempty"`                                                    // length 128
	CardAcceptorLocation                                                  string  `json:"cardAcceptorLocation,omitempty"`                                                  // length 25
	CardAcceptorCity                                                      string  `json:"cardAcceptorCity,omitempty"`                                                      // length 13
	CardAcceptorState                                                     string  `json:"cardAcceptorState,omitempty"`                                                     // length 2
	CardAcceptorZip                                                       string  `json:"cardAcceptorZip,omitempty"`                                                       // length 9
	RetrievalReferenceNumber                                              string  `json:"retrievalReferenceNumber,omitempty"`                                              // length 12
	SystemTraceAuditNumber                                                string  `json:"systemTraceAuditNumber,omitempty"`                                                // length 6
	MerchantId                                                            string  `json:"merchantId,omitempty"`                                                            // length 15
	SubTypeCode                                                           string  `json:"subTypeCode,omitempty"`                                                           // length 6
	MerchantGroupCode                                                     string  `json:"merchantGroupCode,omitempty"`                                                     // length 6
	TerminalId                                                            string  `json:"terminalId,omitempty"`                                                            // length 15
	CashbackAmount                                                        float64 `json:"cashbackAmount,omitempty"`                                                        // length int 10
	SurchargeAmount                                                       float64 `json:"surchargeAmount,omitempty"`                                                       // length int 10
	UnverifiedDepositAmount                                               float64 `json:"unverifiedDepositAmount,omitempty"`                                               // length int 10
	CashDepositAmount                                                     float64 `json:"cashDepositAmount,omitempty"`                                                     // length int 10
	AtmNetworkIndicator                                                   string  `json:"atmNetworkIndicator,omitempty"`                                                   // length 3
	NetworkProviderTypeId                                                 int64   `json:"networkProviderTypeId,omitempty"`                                                 // length int 1
	TransactionFeeAmount                                                  float64 `json:"transactionFeeAmount,omitempty"`                                                  // length 10
	SettlementFeeAmount                                                   float64 `json:"settlementFeeAmount,omitempty"`                                                   // length 10
	AdditionalFees_CCA                                                    float64 `json:"additionalFees_CCA,omitempty"`                                                    // length 10
	AdditionalFees_ICA                                                    float64 `json:"additionalFees_ICA,omitempty"`                                                    // length 10
	MessageTypeIndicator                                                  string  `json:"messageTypeIndicator,omitempty"`                                                  // length 4
	OutputMessageTypeIndicator                                            string  `json:"outputMessageTypeIndicator,omitempty"`                                            // length 4
	OutputResponseCode                                                    string  `json:"outputResponseCode,omitempty"`                                                    // length 2
	AcquirerInstitutionCountryCode                                        string  `json:"acquirerInstitutionCountryCode,omitempty"`                                        // length 3
	AuthorizationIdentificationResponse                                   string  `json:"authorizationIdentificationResponse,omitempty"`                                   // length 12
	ResponseCode                                                          string  `json:"responseCode,omitempty"`                                                          // length 2
	PinValidationCode                                                     string  `json:"pinValidationCode,omitempty"`                                                     // length 1
	AdditionalAmounts_Purchase                                            float64 `json:"additionalAmounts_Purchase,omitempty"`                                            // length 10
	AdditionalAmounts_Gratuity                                            float64 `json:"additionalAmounts_Gratuity,omitempty"`                                            // length 10
	Advice_OriginatorCode                                                 string  `json:"advice_OriginatorCode,omitempty"`                                                 // length 1
	Advice_ReasonCode                                                     string  `json:"advice_ReasonCode,omitempty"`                                                     // length 1
	PrivatelyDefintedData_IssuerNetworkIdCode                             string  `json:"privatelyDefintedData_IssuerNetworkIdCode,omitempty"`                             // length 3
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
	CardAcceptorTerminalIdentification                                    string  `json:"cardAcceptorTerminalIdentification,omitempty"`                                    // length 8
	Cvv2PresenceIndicator                                                 string  `json:"cvv2PresenceIndicator,omitempty"`                                                 // length 1
	Cvv2Result                                                            string  `json:"cvv2Result,omitempty"`                                                            // length 1
	Token                                                                 string  `json:"token,omitempty"`                                                                 // length 19
	TokenAssuranceLevel                                                   string  `json:"tokenAssuranceLevel,omitempty"`                                                   // length 2
	DigitalWalletTokenRequestorTypeId                                     string  `json:"digitalWalletTokenRequestorTypeId,omitempty"`                                     // length 10
	TokenExpirationDate                                                   string  `json:"tokenExpirationDate,omitempty"`                                                   // length 4
	PaymentAccountReferenceNumber                                         string  `json:"paymentAccountReferenceNumber,omitempty"`                                         // length 29
	MerchantCategoryCode                                                  string  `json:"merchantCategoryCode,omitempty"`                                                  // length 4
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
	PointOfServicePanEntryMode                                            string  `json:"pointOfServicePanEntryMode,omitempty"`                                            // length 2
	PointOfServicePinEntryMode                                            string  `json:"pointOfServicePinEntryMode,omitempty"`                                            // length 1
	PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator       string  `json:"privatelyDefinedData_TransactionLevel_CredentialOnFileIndicator,omitempty"`       // length 1
	PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator string  `json:"privatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator,omitempty"` // length 1
	AvsResult                                                             string  `json:"avsResult,omitempty"`                                                             // length 1

}

func (mdl *CardTransactionFileModel) IsTabDelimitedFile() bool {
	return false
}

func (mdl *CardTransactionFileModel) HeaderFieldCount() int {
	return 5
}

func (mdl *CardTransactionFileModel) GetJsonStruct(preserve bool) interface{} {
	if preserve {
		return mdl
	} else {
		return TypedCardTransactionFileModel{
			TransactionId:                             mdl.Typed_TransactionId,
			MasterId:                                  mdl.Typed_MasterId,
			CardId:                                    mdl.Typed_CardId,
			PanLastFour:                               mdl.PanLastFour,
			CardHolderFirstName:                       mdl.CardHolderFirstName,
			CardHolderMiddleName:                      mdl.CardHolderMiddleName,
			CardHolderLastName:                        mdl.CardHolderLastName,
			CardAcceptorLocation:                      mdl.CardAcceptorLocation,
			CardAcceptorCity:                          mdl.CardAcceptorCity,
			CardAcceptorState:                         mdl.CardAcceptorState,
			CardAcceptorZip:                           mdl.CardAcceptorZip,
			RetrievalReferenceNumber:                  mdl.RetrievalReferenceNumber,
			SystemTraceAuditNumber:                    mdl.SystemTraceAuditNumber,
			MerchantId:                                mdl.MerchantId,
			SubTypeCode:                               mdl.SubTypeCode,
			MerchantGroupCode:                         mdl.MerchantGroupCode,
			TerminalId:                                mdl.TerminalId,
			CashbackAmount:                            mdl.Typed_CashbackAmount,
			SurchargeAmount:                           mdl.Typed_SurchargeAmount,
			UnverifiedDepositAmount:                   mdl.Typed_UnverifiedDepositAmount,
			CashDepositAmount:                         mdl.Typed_CashDepositAmount,
			AtmNetworkIndicator:                       mdl.AtmNetworkIndicator,
			NetworkProviderTypeId:                     mdl.Typed_NetworkProviderTypeId,
			TransactionFeeAmount:                      mdl.Typed_TransactionFeeAmount,
			SettlementFeeAmount:                       mdl.Typed_SettlementFeeAmount,
			AdditionalFees_CCA:                        mdl.Typed_AdditionalFees_CCA,
			AdditionalFees_ICA:                        mdl.Typed_AdditionalFees_ICA,
			MessageTypeIndicator:                      mdl.MessageTypeIndicator,
			OutputMessageTypeIndicator:                mdl.OutputMessageTypeIndicator,
			OutputResponseCode:                        mdl.OutputResponseCode,
			AcquirerInstitutionCountryCode:            mdl.AcquirerInstitutionCountryCode,
			AuthorizationIdentificationResponse:       mdl.AuthorizationIdentificationResponse,
			ResponseCode:                              mdl.ResponseCode,
			PinValidationCode:                         mdl.PinValidationCode,
			AdditionalAmounts_Purchase:                mdl.Typed_AdditionalAmounts_Purchase,
			AdditionalAmounts_Gratuity:                mdl.Typed_AdditionalAmounts_Gratuity,
			Advice_OriginatorCode:                     mdl.Advice_OriginatorCode,
			Advice_ReasonCode:                         mdl.Advice_ReasonCode,
			PrivatelyDefintedData_IssuerNetworkIdCode: mdl.PrivatelyDefintedData_IssuerNetworkIdCode,
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
			TextInfo:                           mdl.TextInfo,
			Track2Data_ServiceCode:             mdl.Track2Data_ServiceCode,
			CardAcceptorTerminalIdentification: mdl.CardAcceptorTerminalIdentification,
			Cvv2PresenceIndicator:              mdl.Cvv2PresenceIndicator,
			Cvv2Result:                         mdl.Cvv2Result,
			Token:                              mdl.Token,
			TokenAssuranceLevel:                mdl.TokenAssuranceLevel,
			DigitalWalletTokenRequestorTypeId:  mdl.DigitalWalletTokenRequestorTypeId,
			TokenExpirationDate:                mdl.TokenExpirationDate,
			PaymentAccountReferenceNumber:      mdl.PaymentAccountReferenceNumber,
			MerchantCategoryCode:               mdl.MerchantCategoryCode,
			CardAcceptorRegionCode:             mdl.CardAcceptorRegionCode,
			CardAcceptorCountryCode:            mdl.CardAcceptorCountryCode,
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
			PointOfServicePanEntryMode:                                            mdl.PointOfServicePanEntryMode,
			PointOfServicePinEntryMode:                                            mdl.PointOfServicePinEntryMode,
			PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator:       mdl.PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator,
			PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator: mdl.PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator,
			AvsResult: mdl.AvsResult,
		}
	}
}

func (mdl *CardTransactionFileModel) GetValues() []string {
	return []string{
		mdl.TransactionId,
		mdl.MasterId,
		mdl.CardId,
		mdl.PanLastFour,
		mdl.CardHolderFirstName,
		mdl.CardHolderMiddleName,
		mdl.CardHolderLastName,
		mdl.CardAcceptorLocation,
		mdl.CardAcceptorCity,
		mdl.CardAcceptorState,
		mdl.CardAcceptorZip,
		mdl.RetrievalReferenceNumber,
		mdl.SystemTraceAuditNumber,
		mdl.MerchantId,
		mdl.SubTypeCode,
		mdl.MerchantGroupCode,
		mdl.TerminalId,
		mdl.CashbackAmount,
		mdl.SurchargeAmount,
		mdl.UnverifiedDepositAmount,
		mdl.CashDepositAmount,
		mdl.AtmNetworkIndicator,
		mdl.NetworkProviderTypeId,
		mdl.TransactionFeeAmount,
		mdl.SettlementFeeAmount,
		mdl.AdditionalFees_CCA,
		mdl.AdditionalFees_ICA,
		mdl.MessageTypeIndicator,
		mdl.OutputMessageTypeIndicator,
		mdl.OutputResponseCode,
		mdl.AcquirerInstitutionCountryCode,
		mdl.AuthorizationIdentificationResponse,
		mdl.ResponseCode,
		mdl.PinValidationCode,
		mdl.AdditionalAmounts_Purchase,
		mdl.AdditionalAmounts_Gratuity,
		mdl.Advice_OriginatorCode,
		mdl.Advice_ReasonCode,
		mdl.PrivatelyDefintedData_IssuerNetworkIdCode,
		mdl.PrivatelyDefinedData_AdditionalTransactionElement_FallbackIndicator,
		mdl.PrivatelyDefinedData_ProcessingFlag_SpecialTransactionIndicator,
		mdl.PrivatelyDefinedData_ProcessingFlag_ISAIndicator,
		mdl.PrivatelyDefinedData_ProcessingFlag_PartialAuthIndicator,
		mdl.PrivatelyDefinedData_RiskData_ScoreSource,
		mdl.PrivatelyDefinedData_RiskData_ScoreValue,
		mdl.PrivatelyDefinedData_RiskData_ResponseCode,
		mdl.PrivatelyDefinedData_RiskData_FalconReason1,
		mdl.PrivatelyDefinedData_RiskData_FalconReason2,
		mdl.PrivatelyDefinedData_RiskData_FalconReason3,
		mdl.PrivatelyDefinedData_RiskData_VisaRiskScore,
		mdl.PrivatelyDefinedData_RiskData_VisaRiskReason,
		mdl.PrivatelyDefinedData_RiskData_VisaRiskConditionCode1,
		mdl.PrivatelyDefinedData_RiskData_VisaRiskConditionCode2,
		mdl.PrivatelyDefinedData_RiskData_VisaRiskConditionCode3,
		mdl.PrivatelyDefinedData_RiskData_VAAConditionCode1Rank,
		mdl.PrivatelyDefinedData_RiskData_RTDResultCode,
		mdl.PrivatelyDefinedData_RiskData_TravelStatusIndicator,
		mdl.TextInfo,
		mdl.Track2Data_ServiceCode,
		mdl.CardAcceptorTerminalIdentification,
		mdl.Cvv2PresenceIndicator,
		mdl.Cvv2Result,
		mdl.Token,
		mdl.TokenAssuranceLevel,
		mdl.DigitalWalletTokenRequestorTypeId,
		mdl.TokenExpirationDate,
		mdl.PaymentAccountReferenceNumber,
		mdl.MerchantCategoryCode,
		mdl.CardAcceptorRegionCode,
		mdl.CardAcceptorCountryCode,
		mdl.NationalPointOfServiceCondition_TerminalUnattended,
		mdl.NationalPointOfServiceCondition_TerminalOperator,
		mdl.NationalPointOfServiceCondition_TerminalPremises,
		mdl.NationalPointOfServiceCondition_CardPresentation,
		mdl.NationalPointOfServiceCondition_CardPresence,
		mdl.NationalPointOfServiceCondition_CardRetention,
		mdl.NationalPointOfServiceCondition_CardTransaction,
		mdl.NationalPointOfServiceCondition_SecurityCondition,
		mdl.NationalPointOfServiceCondition_TerminalType,
		mdl.NationalPointOfServiceCondition_TerminalEntryCapability,
		mdl.PointOfServicePanEntryMode,
		mdl.PointOfServicePinEntryMode,
		mdl.PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator,
		mdl.PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator,
		mdl.AvsResult,
	}
}

func (mdl *CardTransactionFileModel) ParseNative(line string, strict bool, preserve bool) error {
	pos := 0
	var err error

	pos, err = ConsumeIntField(&mdl.TransactionId, &mdl.Typed_TransactionId, line, pos, 19, preserve, err)
	pos, err = ConsumeIntField(&mdl.MasterId, &mdl.Typed_MasterId, line, pos, 19, preserve, err)
	pos, err = ConsumeIntField(&mdl.CardId, &mdl.Typed_CardId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.PanLastFour, line, pos, 4, preserve, err)
	pos, err = ConsumeField(&mdl.CardHolderFirstName, line, pos, 64, preserve, err)
	pos, err = ConsumeField(&mdl.CardHolderMiddleName, line, pos, 64, preserve, err)
	pos, err = ConsumeField(&mdl.CardHolderLastName, line, pos, 128, preserve, err)
	pos, err = ConsumeField(&mdl.CardAcceptorLocation, line, pos, 25, preserve, err)
	pos, err = ConsumeField(&mdl.CardAcceptorCity, line, pos, 13, preserve, err)
	pos, err = ConsumeField(&mdl.CardAcceptorState, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.CardAcceptorZip, line, pos, 9, preserve, err)
	pos, err = ConsumeField(&mdl.RetrievalReferenceNumber, line, pos, 12, preserve, err)
	pos, err = ConsumeField(&mdl.SystemTraceAuditNumber, line, pos, 6, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantId, line, pos, 15, preserve, err)
	pos, err = ConsumeField(&mdl.SubTypeCode, line, pos, 6, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantGroupCode, line, pos, 6, preserve, err)
	pos, err = ConsumeField(&mdl.TerminalId, line, pos, 15, preserve, err)
	pos, err = ConsumeFloatField(&mdl.CashbackAmount, &mdl.Typed_CashbackAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.SurchargeAmount, &mdl.Typed_SurchargeAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.UnverifiedDepositAmount, &mdl.Typed_UnverifiedDepositAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.CashDepositAmount, &mdl.Typed_CashDepositAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.AtmNetworkIndicator, line, pos, 3, preserve, err)
	pos, err = ConsumeIntField(&mdl.NetworkProviderTypeId, &mdl.Typed_NetworkProviderTypeId, line, pos, 1, preserve, err)
	pos, err = ConsumeFloatField(&mdl.TransactionFeeAmount, &mdl.Typed_TransactionFeeAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.SettlementFeeAmount, &mdl.Typed_SettlementFeeAmount, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.AdditionalFees_CCA, &mdl.Typed_AdditionalFees_CCA, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.AdditionalFees_ICA, &mdl.Typed_AdditionalFees_ICA, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.MessageTypeIndicator, line, pos, 4, preserve, err)
	pos, err = ConsumeField(&mdl.OutputMessageTypeIndicator, line, pos, 4, preserve, err)
	pos, err = ConsumeField(&mdl.OutputResponseCode, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.AcquirerInstitutionCountryCode, line, pos, 3, preserve, err)
	pos, err = ConsumeField(&mdl.AuthorizationIdentificationResponse, line, pos, 12, preserve, err)
	pos, err = ConsumeField(&mdl.ResponseCode, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PinValidationCode, line, pos, 1, preserve, err)
	pos, err = ConsumeFloatField(&mdl.AdditionalAmounts_Purchase, &mdl.Typed_AdditionalAmounts_Purchase, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeFloatField(&mdl.AdditionalAmounts_Gratuity, &mdl.Typed_AdditionalAmounts_Gratuity, line, pos, 10, 2, 2, preserve, err)
	pos, err = ConsumeField(&mdl.Advice_OriginatorCode, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.Advice_ReasonCode, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefintedData_IssuerNetworkIdCode, line, pos, 3, preserve, err)
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
	pos, err = ConsumeField(&mdl.CardAcceptorTerminalIdentification, line, pos, 8, preserve, err)
	pos, err = ConsumeField(&mdl.Cvv2PresenceIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.Cvv2Result, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.Token, line, pos, 19, preserve, err)
	pos, err = ConsumeField(&mdl.TokenAssuranceLevel, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.DigitalWalletTokenRequestorTypeId, line, pos, 10, preserve, err)
	pos, err = ConsumeField(&mdl.TokenExpirationDate, line, pos, 4, preserve, err)
	pos, err = ConsumeField(&mdl.PaymentAccountReferenceNumber, line, pos, 29, preserve, err)
	pos, err = ConsumeField(&mdl.MerchantCategoryCode, line, pos, 4, preserve, err)
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
	pos, err = ConsumeField(&mdl.PointOfServicePanEntryMode, line, pos, 2, preserve, err)
	pos, err = ConsumeField(&mdl.PointOfServicePinEntryMode, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_TransactionLevel_CredentialOnFileIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.PrivatelyDefinedData_TransactionLevel_CryptocurrencyPurchaseIndicator, line, pos, 1, preserve, err)
	pos, err = ConsumeField(&mdl.AvsResult, line, pos, 1, preserve, err)

	if err == nil && strict && pos < len(line) {
		err = fmt.Errorf("unexpected bytes after last field (expected width=%v, actual width=%v): '%v'", pos, len(line), line[pos:])
	}

	return err
}
