package apperrors

import (
	"context"
	"errors"
	"log"

	"github.com/MatsuoTakuro/my-template-connect-go/api/contexts"
	"github.com/bufbuild/connect-go"
)

func ErrorHandlingGrpc(ctx context.Context, err error) error {
	var connectErr *connect.Error
	if errors.As(err, &connectErr) {
		return err
	}

	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = &AppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := contexts.GetTracdID(ctx)
	log.Printf("[%d]error: %s\n", traceID, appErr)

	switch appErr.ErrCode {
	case NAData:
		return connect.NewError(connect.CodeNotFound, appErr)
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		return connect.NewError(connect.CodeInvalidArgument, appErr)
	case ResBodyEncodeFailed:
		return connect.NewError(connect.CodeInternal, appErr)
	default:
		return connect.NewError(connect.CodeInternal, appErr)
	}
}
