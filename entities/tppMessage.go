package entities

// TppMessage contains information for the TPP
type TppMessage struct {
	// Required: true
	Category Category `json:"category"`

	// Required: true
	Code MessageCode `json:"code"`

	Path string `json:"path,omitempty"`

	// Required: true
	// Max Length: 512
	Text *string `json:"text"`
}

// NewTppMessage creates a new TppMessage
func NewTppMessage(text string, category Category, code MessageCode, path string) TppMessage {
	m := &text
	return TppMessage{
		Text:     m,
		Category: category,
		Code:     code,
		Path:     path,
	}
}

// MessageCode indicates the message category
type MessageCode string

const (
	MessageCodeCERTIFICATEINVALID      MessageCode = "CERTIFICATE_INVALID"
	MessageCodeCERTIFICATEEXPIRED      MessageCode = "CERTIFICATE_EXPIRED"
	MessageCodeCERTIFICATEBLOCKED      MessageCode = "CERTIFICATE_BLOCKED"
	MessageCodeCERTIFICATEREVOKED      MessageCode = "CERTIFICATE_REVOKED"
	MessageCodeCERTIFICATEMISSING      MessageCode = "CERTIFICATE_MISSING"
	MessageCodeSIGNATUREINVALID        MessageCode = "SIGNATURE_INVALID"
	MessageCodeSIGNATUREMISSING        MessageCode = "SIGNATURE_MISSING"
	MessageCodeFORMATERROR             MessageCode = "FORMAT_ERROR"
	MessageCodePARAMETERNOTSUPPORTED   MessageCode = "PARAMETER_NOT_SUPPORTED"
	MessageCodePSUCREDENTIALSINVALID   MessageCode = "PSU_CREDENTIALS_INVALID"
	MessageCodeSERVICEINVALID          MessageCode = "SERVICE_INVALID"
	MessageCodeSERVICEBLOCKED          MessageCode = "SERVICE_BLOCKED"
	MessageCodeCORPORATEIDINVALID      MessageCode = "CORPORATE_ID_INVALID"
	MessageCodeCONSENTUNKNOWN          MessageCode = "CONSENT_UNKNOWN"
	MessageCodeCONSENTINVALID          MessageCode = "CONSENT_INVALID"
	MessageCodeCONSENTEXPIRED          MessageCode = "CONSENT_EXPIRED"
	MessageCodeTOKENUNKNOWN            MessageCode = "TOKEN_UNKNOWN"
	MessageCodeTOKENINVALID            MessageCode = "TOKEN_INVALID"
	MessageCodeTOKENEXPIRED            MessageCode = "TOKEN_EXPIRED"
	MessageCodeRESOURCEUNKNOWN         MessageCode = "RESOURCE_UNKNOWN"
	MessageCodeRESOURCEEXPIRED         MessageCode = "RESOURCE_EXPIRED"
	MessageCodeTIMESTAMPINVALID        MessageCode = "TIMESTAMP_INVALID"
	MessageCodePERIODINVALID           MessageCode = "PERIOD_INVALID"
	MessageCodeSCAMETHODUNKNOWN        MessageCode = "SCA_METHOD_UNKNOWN"
	MessageCodePRODUCTINVALID          MessageCode = "PRODUCT_INVALID"
	MessageCodePRODUCTUNKNOWN          MessageCode = "PRODUCT_UNKNOWN"
	MessageCodePAYMENTFAILED           MessageCode = "PAYMENT_FAILED"
	MessageCodeREQUIREDKIDMISSING      MessageCode = "REQUIRED_KID_MISSING"
	MessageCodeEXECUTIONDATEINVALID    MessageCode = "EXECUTION_DATE_INVALID"
	MessageCodeSESSIONSNOTSUPPORTED    MessageCode = "SESSIONS_NOT_SUPPORTED"
	MessageCodeACCESSEXCEEDED          MessageCode = "ACCESS_EXCEEDED"
	MessageCodeREQUESTEDFORMATSINVALID MessageCode = "REQUESTED_FORMATS_INVALID"
	MessageCodeCARDINVALID             MessageCode = "CARD_INVALID"
	MessageCodeNOPIISACTIVATION        MessageCode = "NO_PIIS_ACTIVATION"
)

// Category of the message
type Category string

const (
	CategoryWARNING Category = "WARNING"
	CategoryERROR   Category = "ERROR"
)
