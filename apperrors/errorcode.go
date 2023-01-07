package apperrors

import "errors"

type ErrCode string

const (
	Unknown ErrCode = "UKN000"

	InsertDataFailed ErrCode = "SVC001"
	GetDataFailed    ErrCode = "SVC002"
	NAData           ErrCode = "SVC003"
	NoTargetData     ErrCode = "SVC004"
	UpdateDataFailed ErrCode = "SVC005"

	ReqBodyDecodeFailed ErrCode = "CTL001"
	BadParam            ErrCode = "CTL002"
	ResBodyEncodeFailed ErrCode = "CTL003"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &AppError{
		ErrCode: code,
		Message: message,
		Err:     err,
	}
}

func (code ErrCode) NewAppError(message string) *AppError {
	return &AppError{
		ErrCode: code,
		Message: message,
		Err:     errors.New(message),
	}
}
