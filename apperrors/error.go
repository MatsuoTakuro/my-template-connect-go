package apperrors

type AppError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

var _ error = (*AppError)(nil)

func (ae *AppError) Error() string {
	return ae.Err.Error()
}

func (ae *AppError) Unwrap() error {
	return ae.Err
}
